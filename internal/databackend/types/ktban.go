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

type WatchEvent string

const (
	Added    WatchEvent = "ADDED"
	Modified WatchEvent = "MODIFIED"
	Deleted  WatchEvent = "DELETED"
	Bookmark WatchEvent = "BOOKMARK"
	Error    WatchEvent = "ERROR"
)

type Event struct {
	Type WatchEvent
	Object Ktban
	Abort bool
}