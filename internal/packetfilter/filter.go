package packetfilter

import (
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"github.com/sers-dev/kubetables/internal/packetfilter/iptables"
	"os/exec"
)

type PacketFilter interface {
	RuleExists (ktban types.Ktban) (bool, error)
	AppendRule (ktban types.Ktban) error
	InsertRule (ktban types.Ktban) error
	DeleteRule (ktban types.Ktban) error
	CreateInitialRules (ktbans types.Ktbans) error
}

func CreatePacketFilter () (pf PacketFilter)  {
	if commandExists("iptables") {
		pf, err := iptables.Initialize()
		if err != nil {
			panic(err.Error())
		}
		return pf
	}

	return nil
}


func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}