package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	//加密
	hello := "您好，世界！hello world"
	debyte := base64Encode([]byte(hello))
	fmt.Println(string(debyte))
	//5oKo5aW977yM5LiW55WM77yBaGVsbG8gd29ybGQ=

	//解密
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}
	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}
	fmt.Println(string(enbyte))
	//您好，世界！hello world
}
