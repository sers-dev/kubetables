package v1alpha1

import (
	"context"
	"github.com/sers-dev/kubetables/internal/kubernetesapi/types/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KtbanInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.KtbanList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Ktban, error)
	//Create(*v1alpha1.Ktban) (*v1alpha1.Ktban, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type ktbanClient struct {
	restClient 	rest.Interface
	ns 			string
}

func (c *ktbanClient) List(opts metav1.ListOptions) (*v1alpha1.KtbanList, error) {
	result := v1alpha1.KtbanList{}
	err := c.restClient.Get().
		Namespace(c.ns).
		Resource("ktbans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *ktbanClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.Ktban, error){
	result := v1alpha1.Ktban{}
	err := c.restClient.Get().
		Namespace(c.ns).
		Resource("ktbans").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *ktbanClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.Get().
		Namespace(c.ns).
		Resource("ktbans").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}