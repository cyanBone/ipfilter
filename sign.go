package ipfilter

import (
	"errors"
	"net"
)

type SignFilter struct {
	IP net.IP
}

func (f *SignFilter) Load(ips string) (IPFilter, error) {
	to4 := net.ParseIP(ips).To4()
	if to4 == nil {
		return f, errors.New("sign filter parse ip error")
	}
	f.IP = to4
	return f, nil
}

func (f *SignFilter) Check(ip net.IP) bool {
	return f.IP.Equal(ip)
}
