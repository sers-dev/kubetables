package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type KtbanSpec struct {
	Ip string `json:"ip"`
	PortFrom int `json:"portFrom"`
	PortTo int `json:"portTo"`
	//InterfaceGroup string `json:"interfaceGroup"`
	Protocol string `json:"protocol"`
	Direction string `json:"direction"`
}

type Ktban struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KtbanSpec `json:"spec"`
}

type KtbanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Ktban `json:"items"`
}