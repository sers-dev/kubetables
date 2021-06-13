package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

func (in *Ktban) DeepCopyInto(out *Ktban) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = KtbanSpec{
		Ip: in.Spec.Ip,
		PortRange: in.Spec.PortRange,
		InterfaceGroup: in.Spec.InterfaceGroup,
		Protocol: in.Spec.Protocol,
		Direction: in.Spec.Direction,
	}
}

func (in *Ktban) DeepCopyObject() runtime.Object {
	out := Ktban{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *KtbanList) DeepCopyObject() runtime.Object {
	out := KtbanList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Ktban, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}