package ipfilter

import (
	"bytes"
	"errors"
	"net"
	"strings"
)

type SizeFilter struct {
	StartIP net.IP
	EndIP   net.IP
}

func (f *SizeFilter) Load(ips string) (IPFilter, error) {
	ipSplit := strings.Split(ips, "-")
	if len(ipSplit) != 2 {
		return nil, errors.New("ip is no - error")
	}

	ip := net.ParseIP(ipSplit[0])
	startIP := ip.To4()
	if startIP == nil {
		return f, errors.New("ip parse - error")
	}

	endIP, err := fixIP(ipSplit[0], ipSplit[1])
	if err != nil {
		return f, err
	}

	f.StartIP = startIP
	f.EndIP = endIP
	return f, err
}

func (f *SizeFilter) Check(ip net.IP) bool {
	return bytes.Compare(ip, f.StartIP) >= 0 && bytes.Compare(ip, f.EndIP) <= 0
}

func fixIP(ip, str string) (net.IP, error) {
	if strings.Contains(str, ".") {
		to4 := net.ParseIP(str).To4()
		if to4 == nil {
			return nil, errors.New("fix ip to ip error")
		}
		return to4, nil
	}

	index := strings.LastIndex(ip, ".")
	newIP := ip[0:index+1] + str
	to4 := net.ParseIP(newIP).To4()
	if to4 == nil {
		return nil, errors.New("fix string to ip error")
	}
	return to4, nil
}
