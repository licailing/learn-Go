package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	fileinfo, err := os.Stat(`E:\go_workspace\learn-Go\Buffer.go`)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileinfo.Name())                       //获取文件名
	fmt.Println(fileinfo.IsDir())                      //判断是否是目录，返回bool类型
	fmt.Println(fileinfo.ModTime().Format("20060102")) //获取文件修改时间
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.Size()) //获取文件大小
	fmt.Println(fileinfo.Sys())
	fmt.Println(reflect.TypeOf(time.Now().Format("20060102")))
}
