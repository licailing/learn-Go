package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Verson      string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIp   string   `xml:"serverIp"`
}

func main() {
	file, err := os.Open("E:/go_workspace/src/learn-Go/config/servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}

/*
{{ servers} 1 [{{ server} Shangehai_VPN } {{ server} Beijing_VPN }]
        <server>
                <serverName>Shangehai_VPN</serverName>
                <serverIP>127.0.0.1</serverIP>
        </server>
        <server>
                <serverName>Beijing_VPN</serverName>
                <serverIP>127.0.0.2</serverIP>
        </server>
}
*/
