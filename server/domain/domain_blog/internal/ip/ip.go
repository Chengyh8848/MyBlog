package ip

import (
	"encoding/hex"
	"errors"
	"net"
	"strconv"
	"strings"
)

// ToInt ip地址转换int类型
func ToInt(ip string) (uint64, error) {
	var t uint64
	if ip == "" {
		ips := GetHostIps()
		if len(ips) == 0 {
			return t, errors.New("获取本机ip失败")
		}
		ip = ips[0]
	}

	ips := strings.Split(ip, ".")
	hx := make([]byte, 0, 4)
	for i, _ := range ips {
		m, _ := strconv.Atoi(ips[i])
		hx = append(hx, byte(m))
	}
	s := hex.EncodeToString(hx)
	t, _ = strconv.ParseUint(s, 16, 64)
	return t, nil
}

// GetHostNetwork 获取主机ip地址与mac地址
func GetHostNetwork() map[string]string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	network := make(map[string]string, len(interfaces))
	for _, face := range interfaces {
		if adders, err := face.Addrs(); err == nil {
			for _, addr := range adders {
				if ipNet, ok := addr.(*net.IPNet); ok {
					if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
						network[ipNet.IP.String()] = face.HardwareAddr.String()
					}
				}
			}
		}
	}
	return network
}

// GetHostIps 获取本地ip地址
func GetHostIps() []string {
	ips := make([]string, 0, 2)
	if interfaces, err := net.Interfaces(); err == nil {
		for _, face := range interfaces {
			if adders, err := face.Addrs(); err == nil {
				for _, addr := range adders {
					if ipNet, ok := addr.(*net.IPNet); ok {
						if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
							ips = append(ips, ipNet.IP.String())
						}
					}
				}
			}
		}
	}
	return ips
}
