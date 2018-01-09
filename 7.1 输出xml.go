package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIp   string `xml:"serverIP"`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shangehai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header)) //输出头部
	os.Stdout.Write(output)             //输出内容
}

/*
<?xml version="1.0" encoding="UTF-8"?>
  <servers version="1">
      <server>
          <serverName>Shangehai_VPN</serverName>
          <serverIP>127.0.0.1</serverIP>
      </server>
      <server>
          <serverName>Beijing_VPN</serverName>
          <serverIP>127.0.0.2</serverIP>
      </server>
  </servers>
*/
