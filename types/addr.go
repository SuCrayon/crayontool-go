package types

import (
	"fmt"
	"sync"
)

type Addr struct {
	IP   string
	Port uint16
	str  string
	once sync.Once
}

func NewAddr(ip string, port uint16) *Addr {
	return &Addr{
		IP:   ip,
		Port: port,
	}
}

func (a *Addr) GetAddrStr() string {
	a.once.Do(func() {
		a.str = fmt.Sprintf("%s:%d", a.IP, a.Port)
	})
	return a.str
}

func (a *Addr) SetIP(ip string) *Addr {
	a.IP = ip
	return a
}

func (a *Addr) SetPort(port uint16) *Addr {
	a.Port = port
	return a
}
