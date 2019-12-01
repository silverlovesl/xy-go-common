package utils

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// ResolveHostIP ホストマシンのアドレスを解析する
func ResolveHostIP() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
			ip := networkIP.IP.String()
			fmt.Println("Resolved Host IP: " + ip)
			return ip
		}
	}
	return ""
}

// GenerateRandomString ランダム文字列生成
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+^*/()[]{}/!@#$?|" // Don't use % char
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}
