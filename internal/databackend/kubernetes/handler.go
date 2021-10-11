package kubernetes

import (
	v1alpha1client "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/clientset/v1alpha1"
	v1alpha1types "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"github.com/sers-dev/kubetables/pkg/auth"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
)

type Handler struct {
	client v1alpha1client.KtbanV1Alpha1Client
}

func Initialize() (*Handler, error) {
	_ = v1alpha1types.AddToScheme(scheme.Scheme)
	config, err := auth.GetKubernetesConfig()
	if err != nil {
		return nil, err
	}

	clientSet, err := v1alpha1client.NewForConfig(&config)
	if err != nil {
		return nil, err
	}
	h := Handler{ client: *clientSet }

	return &h, nil
}

func (h *Handler) List() (types.Ktbans, error) {
	listOptions := v1.ListOptions{}
	ktbanListKube, err := h.client.Ktbans("ktban").List(listOptions)
	if err != nil {
		panic(err.Error())
	}

	return h.ConvertKubernetesList(*ktbanListKube), nil
}

func (h *Handler) Watch() watch.Interface {
	listOptions := v1.ListOptions{}
	watcher, err := h.client.Ktbans("ktban").Watch(listOptions)
	if err != nil {
		panic(err.Error())
	}
	return watcher
}

func (h *Handler) ConvertKubernetesList(kubeList v1alpha1types.KtbanList) types.Ktbans {
	var ktbans types.Ktbans
	itemCount := len(kubeList.Items)
	if itemCount < 1 {
		ktbans.Items = nil

		return ktbans
	}
	ktbans.Items = make([]types.Ktban, itemCount)

	for i := range kubeList.Items {
		ktbans.Items[i] = h.ConvertKtbanType(kubeList.Items[i])
	}

	return ktbans
}

func (h *Handler) ConvertKtbanType(object v1alpha1types.Ktban) types.Ktban {
	ktban := types.Ktban{
		Ip:       object.Spec.Ip,
		PortFrom: object.Spec.PortFrom,
		PortTo:   object.Spec.PortTo,
		//InterfaceGroup: kubeList.Items[i].Spec.InterfaceGroup,
		Protocol:  object.Spec.Protocol,
		Direction: object.Spec.Direction,
	}
	return ktban
}
