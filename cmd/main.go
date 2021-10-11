package main

import (
    "fmt"
    "github.com/sers-dev/kubetables/internal/databackend/kubernetes"
    "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
    "github.com/sers-dev/kubetables/internal/packetfilter"
    "k8s.io/apimachinery/pkg/watch"
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

    watcher := kubeHandler.Watch()
    for event := range watcher.ResultChan() {
        ktban := event.Object.(*v1alpha1.Ktban)

        ktbanType := kubeHandler.ConvertKtbanType(*ktban)
        switch event.Type {
        case watch.Added:
            fmt.Println("ADDED KTBAN")
            err := packetFilter.AppendRule(ktbanType)
            if err != nil {
                panic(err.Error())
            }
        case watch.Modified:
            fmt.Println("MODIFIED KTBAN")
        case watch.Deleted:
            fmt.Println("DELETED KTBAN")

            err := packetFilter.DeleteRule(ktbanType)
            if err != nil {
                panic(err.Error())
            }
        }
    }
}