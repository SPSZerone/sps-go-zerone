package net

import (
	"github.com/BurntSushi/toml"
)

func NewAddrFromContent(content string) (addr Addr, err error) {
	_, err = toml.Decode(content, &addr)
	return
}

func NewAddrFromBytes(bytes []byte) (addr Addr, err error) {
	err = toml.Unmarshal(bytes, &addr)
	return
}

func NewAddrFromFile(file string) (addr Addr, err error) {
	_, err = toml.DecodeFile(file, &addr)
	return
}

type Addr struct {
	Type string `toml:"type"` // e.g. "tcp" "udp"
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Zone string `toml:"zone"`
}

func (a *Addr) GetType() string {
	return a.Type
}

func (a *Addr) GetHost() string {
	return a.Host
}

func (a *Addr) GetPort() int {
	return a.Port
}

func (a *Addr) GetZone() string {
	return a.Zone
}
