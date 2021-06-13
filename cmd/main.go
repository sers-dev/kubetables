package main

import (
    "context"
    "fmt"
    "github.com/sers-dev/kubetables/pkg/auth"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
    kubernetesaccess, nil := auth.GetKubernetesAccess()

    listOptions := metav1.ListOptions{
        LabelSelector: "app.kubernetes.io/component=mail",
    }
    pods, err := kubernetesaccess.ClientSet.CoreV1().Pods("mail").List(context.TODO(), listOptions)
    if err != nil {
        panic(err.Error())
    }

    for _, pod := range pods.Items {
        fmt.Println("PODNAME:", pod.ObjectMeta.Name)
    }
}