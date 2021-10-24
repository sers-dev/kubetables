package main

import (
    "fmt"
    "github.com/sers-dev/kubetables/internal/databackend"
    "github.com/sers-dev/kubetables/internal/databackend/types"
    "github.com/sers-dev/kubetables/internal/packetfilter"
    "os"
    "os/signal"
    "syscall"
)

//https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html
func main() {
    dataBackendHandler, err := databackend.CreateDataBackend()
    if err != nil {
        panic(err.Error())
    }

    packetFilter := packetfilter.CreatePacketFilter()
    ktbans, err := dataBackendHandler.List()
    if err != nil {
        panic(err.Error())
    }

    err = packetFilter.CreateInitialRules(ktbans)
    if err != nil {
        //LOG ERROR
    }
    ch := make(chan types.Event)
    sigs := make(chan os.Signal, 1)
    signalsToTerminate := []os.Signal{ syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL }

    signal.Notify(sigs, signalsToTerminate...)
    go dataBackendHandler.Watch(ch, sigs)

    for event := range ch {
        if event.Abort {
            fmt.Println("Processed terminating signal / watcher encountered error, aborting")
            break
        }
        switch event.Type {
        case types.Added:
            err := packetFilter.AppendRule(event.Object)
            if err != nil {
                panic(err.Error())
            }
        case types.Deleted:
            err := packetFilter.DeleteRule(event.Object)
            if err != nil {
                panic(err.Error())
            }
        case types.Modified:
            // TODO
        }
    }
}