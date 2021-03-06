package constant

import (
	"net"
)

// Adapter Type
const (
	Direct AdapterType = iota
	Reject
	Selector
	Shadowsocks
	Socks5
	URLTest
)

type ProxyAdapter interface {
	Conn() net.Conn
	Close()
}

type ServerAdapter interface {
	Addr() *Addr
	Close()
}

type Proxy interface {
	Name() string
	Type() AdapterType
	Generator(addr *Addr) (ProxyAdapter, error)
}

// AdapterType is enum of adapter type
type AdapterType int

func (at AdapterType) String() string {
	switch at {
	case Direct:
		return "Direct"
	case Reject:
		return "Reject"
	case Selector:
		return "Selector"
	case Shadowsocks:
		return "Shadowsocks"
	case Socks5:
		return "Socks5"
	case URLTest:
		return "URLTest"
	default:
		return "Unknow"
	}
}
