package packetfilter

import (
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"github.com/sers-dev/kubetables/internal/packetfilter/iptables"
)

type PacketFilter interface {
	RuleExists (ktban types.Ktban) (bool, error)
	AppendRule (ktban types.Ktban) (bool, error)
	InsertRule (ktban types.Ktban) (bool, error)
	DeleteRule (ktban types.Ktban) (bool, error)
}

func CreatePacketFilter () (pf PacketFilter)  {
	pf, err := iptables.Initialize()
	if err != nil {
		panic(err.Error())
	}

	return pf
}