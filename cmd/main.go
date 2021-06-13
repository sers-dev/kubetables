package main

import (
    "github.com/sers-dev/kubetables/internal/kubernetesapi"
)

//https://www.martin-helmich.de/en/blog/kubernetes-crd-client.html
func main() {
    //kubernetesAccess, _ := auth.GetKubernetesAccess()

    //listOptions := metav1.ListOptions{
    //    LabelSelector: "app.kubernetes.io/component=mail",
    //}
    //pods, err := kubernetesaccess.ClientSet.CoreV1().Pods("mail").List(context.TODO(), listOptions)
    //if err != nil {
    //    panic(err.Error())
    //}
//
    //for _, pod := range pods.Items {
    //    fmt.Println("PODNAME:", pod.ObjectMeta.Name)
    //}
    kubernetesapi.Prepare()
}