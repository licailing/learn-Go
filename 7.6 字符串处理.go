package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	/*字符串操作*/
	/*
		func Contains(s,substr string) bool
		字符串中是否包含substr,返回bool值
	*/
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true

	/*
		func Join(a []string,sep string) string
		字符串连接
	*/
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", ")) //foo, bar, baz

	/*
		func Index(s,sep string) int
		在字符串中查找sep所在的位置，返回位置值，找不到返回-1
	*/
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1

	/*
		func Repeat(s string,count int) string
		重复s字符串count次，最后返回重复的字符串
	*/
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana

	/*
		func Replace(s,old,new string,n int) string
		替换字符串,n小于0表示全部替换
	*/
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo

	/*
		func Split(s,sep string) []string
		字符串分割,返回切片
	*/
	fmt.Printf("%q\n", strings.Split("a,b,c", ",")) //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("xyz", ""))    //["x" "y" "z"]

	/*
		func Trim(s string,cutset string) string
		去除字符串首尾部指定字符串
	*/
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! ")) //["Achtung"]

	/*
		func Fields(s string) []string
		去除字符串的空格符，并且按照空格分割返回slice
	*/
	fmt.Printf("Fields are: %q\n", strings.Fields(" foo bar baz   ")) //Fields are: ["foo" "bar" "baz"]

	/*字符串转换*/
	/* Append系列函数将整数等转为字符串后，添加到现有的字节数组中 */
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str)) //4567false"abcdefg"'单'

	/* Format系列函数把其他类型转换为字符串 */
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e) //false 123.23 1234 12345 1023

	/* Parse系列函数把字符串转为其他类型 */
	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1) //false 123.23 1234 12345 1023
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
