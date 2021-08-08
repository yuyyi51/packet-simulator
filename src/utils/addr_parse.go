package utils

import (
	"strconv"
	"strings"
)

func ParsePort(addr string) int {
	strs := strings.Split(addr, ":")
	portStr := strs[len(strs)-1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0
	}
	return port
}

func ParseAddr(addr string) string {
	strs := strings.Split(addr, ":")
	addrStr := strs[0]
	return addrStr
}
