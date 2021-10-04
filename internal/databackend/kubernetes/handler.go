package kubernetes

import (
	v1alpha1client "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/clientset/v1alpha1"
	v1alpha1types "github.com/sers-dev/kubetables/internal/databackend/kubernetes/api/types/v1alpha1"
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"github.com/sers-dev/kubetables/pkg/auth"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type Handler struct {
	client v1alpha1client.KtbanV1Alpha1Client
}

func (h *Handler) List() (types.Ktbans, error) {
	listOptions := v1.ListOptions{

	}
	ktbanListKube, err := h.client.Ktbans("ktban").List(listOptions)
	if err != nil {
		panic(err.Error())
	}

	return h.ConvertKubernetesList(*ktbanListKube), nil
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
		ktban := types.Ktban{
			Ip: kubeList.Items[i].Spec.Ip,
			PortFrom: kubeList.Items[i].Spec.PortFrom,
			PortTo: kubeList.Items[i].Spec.PortTo,
			//InterfaceGroup: kubeList.Items[i].Spec.InterfaceGroup,
			Protocol: kubeList.Items[i].Spec.Protocol,
			Direction: kubeList.Items[i].Spec.Direction,
		}
		ktbans.Items[i] = ktban
	}

	return ktbans
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
