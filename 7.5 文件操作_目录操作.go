/**
目录操作
//创建名为name的目录，权限设置是perm，例如0777
func Mkdir(name string,perm FileMode) error

//根据path创建多级子目录，例如astaxie/test1/test2
func MkdirAll(path string,perm FileMode) error

//删除名称为name的目录，当目录下有文件或则其他目录时会出错
//删除名称为name的文件
func Remove(name string) error

//根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除
func RemoveAll(path string) error

*/
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)
	err := os.Remove("astaxie")
	if err != nil {
		fmt.Println(err)
		//remove astaxie: The directory is not empty.
	}
	os.RemoveAll("astaxie")
}
