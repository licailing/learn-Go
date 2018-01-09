package main

import (
	"encoding/json"
	"fmt"
)

/* struct字段(首字母大写) */
type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)                       //{[{Shanghai_VPN 127.0.0.1} {Beijing_VPN 127.0.0.2}]}
	fmt.Println(s.Servers[0].ServerName) //Shanghai_VPN
	for k, v := range s.Servers {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}
	/*
		key:0,value:{Shanghai_VPN 127.0.0.1}
		key:1,value:{Beijing_VPN 127.0.0.2}
	*/
}
