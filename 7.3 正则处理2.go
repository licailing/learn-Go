package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))
	/* Find: am */

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，不然就是返回固定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll:", all)
	/* FindAll: [[97 109] [108 101 97 114] [110 105 110 103] [108 97 110 103] [117 97 103 101]] */

	//查找符合条件的index位置，开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex:", index)
	/* FindIndex: [2 4] */

	//查找符合条件的所有的index位置,n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allindex)
	/* FindAllIndex: [[2 4] [5 9] [9 13] [17 21] [21 25]] */

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()里面的
	//第三个是第二个()里面的
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)
	/* FindSubmatch: [[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110
	 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 10
	3 101]] */
	for _, v := range submatch {
		fmt.Println(string(v))
	}
	/*
		am learning Go language
		 learning Go
		uage
	*/

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)
	/* [2 25 4 17 21 25] */

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)
	/* [[[97 109 32 108 101 97 114 110 105 110 103 32 71 111 32 108 97 110 103 117 97 103 101] [32 108 101 97 114 110 105 110 103 32 71 111 32] [117 97 103 101]]] */

	//FindAllSubmatchIndex,查找所有字匹配的index
	submathallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submathallindex)
	/* [[2 25 4 17 21 25]] */
}
