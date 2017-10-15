package main

import "fmt"

//https://studygolang.com/articles/5877
func main() {
	//初始化切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := make([]int, 10, 10) //内置函数：make([]type,length,capacity)
	b := numbers[2:]         //切分现有数组或切片生成切片arr[startIndex:endIndex] --endIndex(不包括)

	printSlice(numbers)
	// len=10 cap=10 slice=[0 1 2 3 4 5 6 7 8 9]
	printSlice(a)
	// len=10 cap=10 slice=[0 0 0 0 0 0 0 0 0 0]
	printSlice(b)
	// len=8 cap=8 slice=[2 3 4 5 6 7 8 9]

	/* 切取现有数组生成新切片 */
	fmt.Println("numbers[:] == ", numbers[:])
	// numbers[:] ==  [0 1 2 3 4 5 6 7 8 9]
	fmt.Println("numbers[2:] == ", numbers[2:])
	// numbers[2:] ==  [2 3 4 5 6 7 8 9]
	fmt.Println("numbers[:4] == ", numbers[:4])
	// numbers[:4] ==  [0 1 2 3]
	fmt.Println("numbers[1:8] == ", numbers[1:8])
	// numbers[1:8] ==  [1 2 3 4 5 6 7]

	/* 切片切取原数组（生成引用）,改变新切片就会改变原数组 */
	b[1] = 9
	printSlice(numbers)
	// len=10 cap=10 slice=[0 1 2 9 4 5 6 7 8 9]
	printSlice(b)
	// len=8 cap=8 slice=[2 9 4 5 6 7 8 9]

	/* 切片生长(复制和追加) */
	var s []int   //空切片
	printSlice(s) //len=0 cap=0 slice=[]

	s = append(s, 1) /* append(dst,src) 追加 */
	printSlice(s)    //len=1 cap=2 slice=[1]

	x := []int{1, 2, 3}
	s = append(s, x...) //追加数组
	printSlice(x)       //len=3 cap=3 slice=[1 2 3]

	/*创建s两倍容量的新切片*/
	s1 := make([]int, len(s), (cap(x)+1)*2) //+1 in case cap(s) == 0
	copy(s1, s)                             //复制
	printSlice(s1)                          //len=4 cap=8 slice=[1 1 2 3]
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
