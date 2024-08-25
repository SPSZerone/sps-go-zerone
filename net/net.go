package net

import "net"

type Addresses interface {
	Addresses() []net.Addr
	Append(addr ...net.Addr)
}

func NewTCPAddrIPv4(host string, port int) *net.TCPAddr {
	return &net.TCPAddr{IP: ParseIP(host), Port: port}
}

func NewTCPAddrIPv6(host string, port int, zone string) *net.TCPAddr {
	return &net.TCPAddr{IP: ParseIP(host), Port: port, Zone: zone}
}

func ParseIP(s string) net.IP {
	if s == "localhost" {
		return net.ParseIP("127.0.0.1")
	}
	return net.ParseIP(s)
}

func NewAddresses(addr ...net.Addr) Addresses {
	return NewAddressList(addr...)
}

func NewAddressList(addr ...net.Addr) *AddressList {
	addresses := &AddressList{}
	addresses.Init(addr...)
	return addresses
}

type AddressList struct {
	addresses []net.Addr
}

func (a *AddressList) Init(addr ...net.Addr) {
	a.addresses = make([]net.Addr, 0)
	if len(addr) > 0 {
		a.Append(addr...)
	}
}

func (a *AddressList) Append(addr ...net.Addr) {
	a.addresses = append(a.addresses, addr...)
}

func (a *AddressList) AppendTCPAddrIPv4(host string, port int) {
	a.Append(NewTCPAddrIPv4(host, port))
}

func (a *AddressList) AppendTCPAddrIPv6(host string, port int, zone string) {
	a.Append(NewTCPAddrIPv6(host, port, zone))
}

func (a *AddressList) Addresses() []net.Addr {
	return a.addresses
}
