package kubernetesapi

import (
	"fmt"
	clientV1alpha1 "github.com/sers-dev/kubetables/internal/kubernetesapi/clientset/v1alpha1"
	typesV1alpha1 "github.com/sers-dev/kubetables/internal/kubernetesapi/types/v1alpha1"
	"github.com/sers-dev/kubetables/pkg/auth"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func Prepare() {
	_ = typesV1alpha1.AddToScheme(scheme.Scheme)
	config, err := auth.GetKubernetesConfig()
	if err != nil {
		panic(err.Error())
	}
	listOptions := metav1.ListOptions{}

	clientSet, err := clientV1alpha1.NewForConfig(&config)
	if err != nil {
		panic(err.Error())
	}
	ktbanList, err := clientSet.Ktbans("ktban").List(listOptions)

	if err != nil {
		panic(err.Error())
	}
	for i := range ktbanList.Items {
		fmt.Println("ITEM: ", ktbanList.Items[i].Name)
	}
}
