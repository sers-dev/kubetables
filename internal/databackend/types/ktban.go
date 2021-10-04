package types

type Ktban struct {
	Ip string
	PortFrom int
	PortTo int
	InterfaceGroup string
	Protocol string
	Direction string
}

type Ktbans struct {
	Items []Ktban
}