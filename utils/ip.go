package utils

import(
	"net"
)
//CommonPrefixLen return the longest common prefix 
func CommonPrefixLen (ip1, ip2 net.IP) int {
	ip1Byte := []byte(ip1)
	ip2Byte := []byte(ip2)
	var cpl int
	var b1, b2 byte
	skip := 8
	for i := 0; i < 4; i++ {
		
		if ip1Byte[i] == ip2Byte[i]{
			cpl += skip
			continue
		}
		b1, b2 =ip1Byte[i],ip2Byte[i]
		for j := 0; j < skip; j++ {
			b1 >>= 1
			b2 >>=1
			if b1 == b2 {
				cpl += skip - j -1
				return cpl
			}
		}
	}
	return cpl
}