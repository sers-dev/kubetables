package main

import (
    "github.com/sers-dev/kubetables/internal/databackend/kubernetes"
    "github.com/sers-dev/kubetables/internal/databackend/types"
    "github.com/sers-dev/kubetables/internal/packetfilter"
)

//https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html
func main() {
    kubeHandler, err := kubernetes.Initialize()
    if err != nil {
        panic(err.Error())
    }

    packetFilter := packetfilter.CreatePacketFilter()
    ktbans, err := kubeHandler.List()
    if err != nil {
        panic(err.Error())
    }

    err = packetFilter.CreateInitialRules(ktbans)
    if err != nil {
        //LOG ERROR
    }
    ch := make(chan types.Event)
    go kubeHandler.Watch(ch)

    for event := range ch {
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
    ch <- types.Event {
        Abort: true,
    }
}