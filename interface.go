package ipfilter

import "net"

type IPFilter interface {
	//载入配置
	Load(ips string) (IPFilter,error)

	//校验
	Check(ip net.IP) bool
}
