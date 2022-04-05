package main

import (
	"fmt"

	localip "github.com/shiyzu/local-ip-address"
)

func main() {
	ipAddr := localip.GetLolcalIp()
	fmt.Println(ipAddr)
}
