package main

import (
    "fmt"
    "github.com/sers-dev/kubetables/internal/databackend/kubernetes"
    "github.com/sers-dev/kubetables/internal/packetfilter"
)

//https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html
func main() {
    fmt.Println("INITIALIZE KUBERNETES")
    kubeHandler, err := kubernetes.Initialize()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("KUBERNETES LIST")
    ktbans, err := kubeHandler.List()
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("IPT INITIALIZE")
    packetFilter := packetfilter.CreatePacketFilter()
    err = packetFilter.CreateInitialRules(ktbans)
    if err != nil {
        //LOG ERROR
    }

    if err != nil {
        panic(err.Error())
    }
    if ktbans.Items != nil {
        for i := range ktbans.Items {
            fmt.Println("IP:", ktbans.Items[i].Ip)
            ruleExists, _ := packetFilter.RuleExists(ktbans.Items[i])
            println("RULE EXISTS?", ruleExists)
            if !ruleExists {
                err := packetFilter.AppendRule(ktbans.Items[i])
                if err != nil {
                    panic(err.Error())
                }
            }

            fmt.Println("DIRECTION:", ktbans.Items[i].Direction)
        }
    }
}