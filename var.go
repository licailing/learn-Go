package main

var x, y int
var ( //这种只能出现在全局变量中，函数体内不支持
	a int
	b bool
)

//并行或同时赋值
var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
// g,h := 123,"hello"
/*会报错
# command-line-arguments
.\var.go:12:1: syntax error: non-declaration statement outside function body
*/

func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
	//0 0 0 false 1 2 123 hello 123 hello
}
