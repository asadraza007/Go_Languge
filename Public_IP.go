package main

import (
	"fmt"
	"net"
	"strings"
)

//PublicIP Get System PublicIP
func PublicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

func main() {
	fmt.Println("Public ip:", PublicIP())
}
