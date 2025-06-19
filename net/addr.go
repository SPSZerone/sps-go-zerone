package net

type AddrConfig interface {
	GetType() string
	GetHost() string
	GetPort() int
	GetZone() string
}
