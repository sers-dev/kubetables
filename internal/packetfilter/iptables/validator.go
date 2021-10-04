package iptables

var Protocols = [...]string{
"#",
"ah",
"all",
"ax.25",
"dccp",
"ddp",
"egp",
"eigrp",
"encap",
"esp",
"etherip",
"fc",
"ggp",
"gre",
"hip",
"hmp",
"hopopt",
"icmp",
"icmpv6",
"idpr-cmtp",
"idrp",
"igmp",
"igp",
"ip",
"ipcomp",
"ipencap",
"ipip",
"ipv6",
"ipv6-frag",
"ipv6-icmp",
"ipv6-nonxt",
"ipv6-opts",
"ipv6-route",
"isis",
"iso-tp4",
"l2tp",
"manet",
"mh",
"mobility-header",
"mpls-in-ip",
"ospf",
"pim",
"pup",
"rdp",
"rohc",
"rspf",
"rsvp",
"sctp",
"shim6",
"skip",
"st",
"tcp",
"udp",
"udplite",
"vmtp",
"vrrp",
"wesp",
"xns-idp",
"xtp",
}

func isValidProto(protoname *string) bool {
	protoFound := false
	for _, pn := range Protocols {
		if pn == *protoname {
			protoFound = true
		}
	}
	return protoFound
}