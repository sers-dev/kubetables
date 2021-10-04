package v1alpha1

import (
	v1alpha12 "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KtbanV1Alpha1Interface interface {
	Ktbans(namespace string) KtbanInterface
}

type KtbanV1Alpha1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*KtbanV1Alpha1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha12.GroupName, Version: v1alpha12.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KtbanV1Alpha1Client{ restClient: client }, nil
}

func (c *KtbanV1Alpha1Client) Ktbans(namespace string) KtbanInterface {
	return &ktbanClient{
		restClient: c.restClient,
		ns: namespace,
	}
}