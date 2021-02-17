package ipfilter

import (
	"net"
)

type MaskFilter struct {
	IPNet net.IPNet
}

func (f *MaskFilter) Load(ips string) (IPFilter, error) {
	_, ipNet, err := net.ParseCIDR(ips)
	if err != nil {
		return f, err
	}
	f.IPNet = *ipNet
	return f, err
}

func (f *MaskFilter) Check(ip net.IP) bool {
	return f.IPNet.Contains(ip)
}
