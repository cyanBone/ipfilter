package ipfilter

import (
	"net"
	"strings"
)

type IPFilters struct {
	Config []IPFilter
}

func NewIPFilters() *IPFilters {
	return &IPFilters{}
}

func (I *IPFilters) Load(ips string) (IPFilter, error) {
	IpSplit := strings.Split(ips, "\n")

	for _, s2 := range IpSplit {
		chars := delChar(s2)
		if chars != "" {
			var filter IPFilter
			switch {
			case strings.Contains(chars, "-"):
				filter = &SizeFilter{}
				break
			case strings.Contains(chars, "/"):
				filter = &MaskFilter{}
				break
			default:
				filter = &SignFilter{}
				break
			}

			load, err := filter.Load(chars)
			if err != nil {
				return nil, err
			}
			I.Config = append(I.Config, load)
		}
	}
	return I, nil
}

func (I *IPFilters) Check(ip net.IP) bool {
	for _, filter := range I.Config {
		check := filter.Check(ip)
		if check {
			return check
		}
	}

	return false
}

func delChar(str string) string {
	//删除空格
	str = strings.Replace(str, " ", "", -1)

	//删除linux换行多的字符
	str = strings.Replace(str, "\r", "", -1)
	return str
}
