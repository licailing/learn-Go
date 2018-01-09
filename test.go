package main

import "fmt"
import (
	"os/exec"
	//"time"
)

func main() {
	out, err := exec.Command("phantomjs", "E:/phantomjs_workspace/page.js", "http://localhost/stock-star/dist/payment.html", "a.png").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	//fmt.Println(time.Now().Unix()) //1515487793
}
