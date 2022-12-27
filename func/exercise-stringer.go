package main

import (
	"fmt"
	"strconv"
)

type IpAddr [4]int64

func (ip IpAddr) String() string {
	ipString := ""
	countAddr := len(ip)
	for key, val := range ip {
		if key+1 == countAddr {
			ipString += strconv.FormatInt(val, 10)
			break
		}
		ipString += strconv.FormatInt(val, 10) + "."
	}
	return ipString
}

func main() {

	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDns": {8, 3, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v:%v \n", name, ip)
	}
	fmt.Println("-------------------------")
	for key, val := range hosts {
		fmt.Println(key, val.String())
	}

	//var test = IpAddr{127, 0, 0, 1}
	//fmt.Println(test.String())
}
