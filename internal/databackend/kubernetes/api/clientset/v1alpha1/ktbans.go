package v1alpha1

import (
	"context"
	v1alpha1 "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
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

//func (c *ktbanClient) StartWatch(addFunc kubernetes.AddResource, updateFunc kubernetes.UpdateResource, deleteFunc kubernetes.DeleteResource) {
//	access, err := auth.GetKubernetesAccess()
//	if err != nil {
//		panic(err.Error())
//	}
//	optionsModifierFunc := func(options *metav1.ListOptions) {}
//
//	watchlist := cache.NewFilteredListWatchFromClient(
//		access.ClientSet.CoreV1().RESTClient(),
//		"ktbans",
//		c.ns,
//		optionsModifierFunc)
//
//	_, controller := cache.NewInformer(
//		watchlist,
//		&v1alpha1.Ktban{},
//		1*time.Minute,
//		cache.ResourceEventHandlerFuncs{
//			AddFunc: func(obj interface{}) {
//				ktban := obj.(*v1alpha1.Ktban)
//				kubernetes.AddKtban(ktban)
//			},
//			UpdateFunc: func(oldObj, newObj interface{}) {
//
//			},
//			DeleteFunc: func(oldObj interface{}){
//
//			},
//		},
//		)
//	go controller.Run(make(chan struct{}))
//	for {
//		time.Sleep(time.Second)
//	}
//}

