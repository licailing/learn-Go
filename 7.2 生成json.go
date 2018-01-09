package main

import (
	"encoding/json"
	"os"
)

type Server struct {
	/* ID不会导出到json中 */
	ID int `json:"-"`

	/* ServerName2 的值会进行二次JSON编码 */
	/* tag中有自定义名称（struct 字段为首字母大写，如果想首字母小写） */
	ServerName string `json:"serverName"`

	/* 如果字段类型是bool,string,int,int64 */
	/* 而tag中带有",string"的选项，那么字段输出到json中时会把json转为json字符串 */
	ServerName2 string `json:"serverName2,string"`

	/* 如果ServerIP 为空，则不会输出到json串中 */
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {
	s := Server{
		ID:          3,
		ServerName:  `Go "1.0"`,
		ServerName2: `Go "1.0"`,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b) //{"serverName":"Go \"1.0\"","serverName2":"\"Go \\\"1.0\\\"\""}
}
