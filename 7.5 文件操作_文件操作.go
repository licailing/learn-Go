/**
**文件操作

*新建文件
//根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的
func Create(name string) (file *File,err Error)

//根据文件描述符创建相应的文件，返回一个文件对象
func NewFile(fd uintptr,name string) *File

*打开文件
//该方法打开一个名称为name的文件，但是是只读方式，内部实现其实调用了OpenFile
func Open(name string) (file *File,err Error)

//打开名称为name的文件，flag是打开的方式，只读、读写等,perm是权限
func OpenFile(name string,flag int,perm uint32) (file *File,err Error)

**写文件
//写入byte类型的信息到文件
func (file *File) Write(b []byte) (n int,err Error)

//在指定位置开始写入byte类型的信息
func (file *File) WriteAt(b []byte,off int64) (n int,err Error)

//写入string信息到文件
func (file *File) WriteString(s string) (ret int,err Error)

**读文件
//读取数据到b中
func (file *File) Read(b []byte) (n int,err Error)

//从off开始读取数据到b中
func (file *File) ReadAt(b []byte,off int64) (n int,err Error)
*/

package main

import (
	"fmt"
	"os"
)

// func main() {
// 	userFile := "astaxie.txt"
// 	fout, err := os.Create(userFile)
// 	if err != nil {
// 		fmt.Println(userFile, err)
// 		return
// 	}
// 	defer fout.Close()
// 	for i := 0; i < 10; i++ {
// 		fout.WriteString("Just a test!\r\n")
// 		fout.Write([]byte("Just a test!\r\n"))
// 	}
// }

func main() {
	userFile := "astaxie.txt"
	//只读方法
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
