package simulate

import "fmt"

type Addr struct {
	ip   string
	port int
}

func (a *Addr) Network() string {
	return "simulate"
}

func (a *Addr) String() string {
	return fmt.Sprintf("%s:%d", a.ip, a.port)
}

func NewAddr(ip string, port int) *Addr {
	return &Addr{
		ip:   ip,
		port: port,
	}
}
