package fetch

import (
	"encoding/binary"
	"net"
)

func GetIpsFromMask(ipMask string) (ips []string, err error) {
	_, ipv4Net, err := net.ParseCIDR(ipMask)
	if err != nil {
		return nil, err
	}

	mask := binary.BigEndian.Uint32(ipv4Net.Mask)
	start := binary.BigEndian.Uint32(ipv4Net.IP)
	finish := (start & mask) | (mask ^ ((1 << 32) - 1))

	for i := start; i <= finish; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		ips = append(ips, ip.String())
	}

	return ips, nil
}
