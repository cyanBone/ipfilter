package ipfilter

import (
	"net"
	"testing"
)

func TestIPFilters_Load(t *testing.T) {
	_, err := NewIPFilters().Load("192.168.0.1\n192.168.0.2\n10.0.0.1/24")
	if err != nil {
		t.Error("数据载入错误")
	}
}

func TestIPFilters_Check(t *testing.T) {
	load, err := NewIPFilters().Load("192.168.0.1\n192.168.0.2\n10.0.0.1/24")
	if err != nil {
		t.Error("数据载入错误")
	}

	ip := net.ParseIP("10.0.0.1")

	check := load.Check(ip)
	if !check {
		t.Error("数据校验错误")
	}
}

func BenchmarkIPFilters_Check(b *testing.B) {
	for i := 0; i < b.N; i++ {
		load, err := NewIPFilters().Load("192.168.0.1\n192.168.0.2\n10.0.0.1/24")
		if err != nil {
			b.Error("数据载入错误")
		}
		ip := net.ParseIP("192.168.0.1")

		check := load.Check(ip)
		if !check {
			b.Error("数据校验错误")
		}
	}
}