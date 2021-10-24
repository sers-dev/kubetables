package kubernetes

import (
	v1alpha1client "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/clientset/v1alpha1"
	v1alpha1types "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"github.com/sers-dev/kubetables/pkg/auth"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"os"
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

	return h.convertKubernetesList(*ktbanListKube), nil
}

func (h *Handler) Watch(ch chan types.Event, sigs chan os.Signal) {
	listOptions := v1.ListOptions{}
	watcher, err := h.client.Ktbans("ktban").Watch(listOptions)
	if err != nil {
		panic(err.Error())
	}

	eventsToActOn := []watch.EventType{ watch.Added, watch.Modified, watch.Deleted }
	resultChan :=  watcher.ResultChan()
	for {
		select {
			case kubernetesEvent := <-resultChan:
				for _, eventToActOn := range eventsToActOn {
					if eventToActOn == kubernetesEvent.Type {
						abstractEvent := h.convertEventTypes(kubernetesEvent.Type)
						kubernetesKtbanObj := kubernetesEvent.Object.(*v1alpha1types.Ktban)
						abstractObj := h.convertKtbanType(*kubernetesKtbanObj)
						ch <- types.Event{
							Type:   abstractEvent,
							Object: abstractObj,
							Abort:  false,
						}
						continue
					}
				}
				if kubernetesEvent.Type == watch.Error {
					watcher.Stop()
					ch <- types.Event{ Abort: true }
					return
				}
			case <-sigs:
				watcher.Stop()
				ch <- types.Event{ Abort: true }
				return
		}
	}
}

func (h *Handler) convertKubernetesList(kubeList v1alpha1types.KtbanList) types.Ktbans {
	var ktbans types.Ktbans
	itemCount := len(kubeList.Items)
	if itemCount < 1 {
		ktbans.Items = nil

		return ktbans
	}
	ktbans.Items = make([]types.Ktban, itemCount)

	for i := range kubeList.Items {
		ktbans.Items[i] = h.convertKtbanType(kubeList.Items[i])
	}

	return ktbans
}

func (h *Handler) convertKtbanType(object v1alpha1types.Ktban) types.Ktban {
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

func (h *Handler) convertEventTypes(kubernetesEvent watch.EventType) types.WatchEvent {
	var abstractEvent types.WatchEvent
	switch kubernetesEvent {
	case watch.Added:
		abstractEvent = types.Added
	case watch.Modified:
		abstractEvent = types.Modified
	case watch.Deleted:
		abstractEvent = types.Deleted
	}
	return abstractEvent
}