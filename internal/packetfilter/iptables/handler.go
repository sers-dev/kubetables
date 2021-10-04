package iptables

import (
	"errors"
	"fmt"
	"github.com/coreos/go-iptables/iptables"
	"github.com/sers-dev/kubetables/internal/databackend/types"
	"net"
	"strings"
)

type Handler struct {
	ip4Tables *iptables.IPTables
	ip6Tables *iptables.IPTables
}

type IptCommand struct {
	Chain    string
	Table    string
	RuleSpec []string
}

var chainName = "kubetables"
var tableName = "filter"


func Initialize() (*Handler, error) {
	h := Handler{}
	var err error
	h.ip4Tables, err = iptables.NewWithProtocol(iptables.ProtocolIPv4)
	if err != nil {
		panic(err.Error())
	}
	h.ip6Tables, err = iptables.NewWithProtocol(iptables.ProtocolIPv6)
	if err != nil {
		panic(err.Error())
	}

	createChainIfNotExists(h.ip4Tables)
	createChainIfNotExists(h.ip6Tables)

	return &h,nil
}

func (h *Handler) RuleExists(ktban types.Ktban) (bool, error) {
	iptc, err := h.buildCommand(ktban)
	if err != nil {
		panic(err.Error())
	}
	var ruleExists bool
	if isIpv4(&ktban.Ip) {
		ruleExists, err = h.ip4Tables.Exists(tableName, chainName, iptc.RuleSpec...)
	} else {
		ruleExists, err = h.ip6Tables.Exists(tableName, chainName, iptc.RuleSpec...)
	}
	if err != nil {
		panic(err.Error())
	}
	return ruleExists, nil
}

func (h *Handler) AppendRule (ktban types.Ktban) (bool, error) {
	command, err := h.buildCommand(ktban)
	if err != nil {
		panic(err.Error())
	}
	if isIpv4(&ktban.Ip) {
		err := h.ip4Tables.Append(command.Table, command.Chain, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := h.ip6Tables.Append(command.Table, command.Chain, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	}
	return true, nil
}

//InsertRule TO DO: .INSERT PARAM POS 1 ?
func (h *Handler) InsertRule (ktban types.Ktban) (bool, error) {
	command, err := h.buildCommand(ktban)
	if err != nil {
		panic(err.Error())
	}
	if isIpv4(&ktban.Ip) {
		err := h.ip4Tables.Insert(command.Table, command.Chain, 1, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := h.ip6Tables.Insert(command.Table, command.Chain, 1, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	}
	return true, nil
}

func (h *Handler) DeleteRule (ktban types.Ktban) (bool, error) {
	command, err := h.buildCommand(ktban)
	if err != nil {
		panic(err.Error())
	}
	if isIpv4(&ktban.Ip) {
		err := h.ip4Tables.Delete(command.Table, command.Chain, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	} else {
		err := h.ip6Tables.Delete(command.Table, command.Chain, command.RuleSpec...)
		if err != nil {
			panic(err.Error())
		}
	}
	return true, nil
}

func createChainIfNotExists (iptablesVx *iptables.IPTables) {
	chainExists, err := iptablesVx.ChainExists(tableName, chainName)
	if err != nil {
		panic(err.Error())
	}
	if !chainExists {
		err := iptablesVx.NewChain(tableName, chainName)
		if err != nil {
			panic(err.Error())
		}
	}
}

func isIpv4(ip *string) bool {
	parsedIp := net.ParseIP(*ip)
	if parsedIp.To4() != nil {
		return true
	}
	return false
}

func (h *Handler) getCorrectVersion(ip *string) (*iptables.IPTables, error) {
	ipCidrSplit := strings.Split(*ip, "/")
	if net.ParseIP(ipCidrSplit[0]) == nil {
		return nil, errors.New(fmt.Sprintf("IP could not be parsed, apparently invalid: %s", *ip))
	}
	for i := 0; i < len(ipCidrSplit[0]); i++ {
		switch ipCidrSplit[0][i] {
		case '.':
			return h.ip4Tables, nil
		case ':':
			return h.ip6Tables, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Could not classify IP %s", ipCidrSplit[0]))
}

func (h *Handler) buildCommand(ktban types.Ktban) (*IptCommand, error) {
	var ruleSpec []string

	// CHECK DIRECTION
	var input bool
	if strings.ToLower(ktban.Direction) == "in" {
		input = true
	} else if strings.ToLower(ktban.Direction) == "out"  {
		input = false
	} else {
		return nil, errors.New(fmt.Sprintf("Invalid direction %s", ktban.Direction))
	}

	// SET PROTOCOL
	ruleSpec = append(ruleSpec, "--protocol")
	lowerCaseProto := strings.ToLower(ktban.Protocol)
	strippedLowerProto := strings.TrimPrefix(lowerCaseProto, "!")

	if !isValidProto(&strippedLowerProto) {
		return nil, errors.New(fmt.Sprintf("Invalid protocol name %s", ktban.Protocol))
	}
	ruleSpec = append(ruleSpec, lowerCaseProto)

	// SET PORT
	if input {
		ruleSpec = append(ruleSpec, "--dport")
	} else {
		ruleSpec = append(ruleSpec, "--sport")
	}
	ruleSpec = append(ruleSpec, fmt.Sprintf("%d:%d", ktban.PortFrom, ktban.PortTo))

	// SET IP
	if input {
		ruleSpec = append(ruleSpec, "-s")
	} else {
		ruleSpec = append(ruleSpec, "-d")
	}
	ruleSpec = append(ruleSpec, strings.Split(ktban.Ip, "/")[0])

	// TO-DO: SET INTERFACE GROUPS

	iptc := IptCommand{
		Table : tableName,
		Chain : chainName,
		RuleSpec : ruleSpec,
	}
	return &iptc, nil
}