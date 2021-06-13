package kubernetesapi

import (
	"context"
	"github.com/sers-dev/kubetables/internal/kubernetesapi/types/v1alpha1"
	"github.com/sers-dev/kubetables/pkg/auth"
	"k8s.io/client-go/kubernetes/scheme"
)

func prepare(kubeAccess auth.KubernetesAccess) {
	_ = v1alpha1.AddToScheme(scheme.Scheme)
	result := v1alpha1.KtbanList{}
	err := kubeAccess.ClientSet.RESTClient().
		Get().
		Resource("ktbans").
		Do(context.TODO()).
		Into(&result)

	if err != nil {
		panic(err.Error())
	}
}