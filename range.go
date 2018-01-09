package main 

import "fmt"

func main(){
	/* range用于for循环中迭代数组(array)、切片(slice)、链表(channel)或集合(map)的元素 */
	/* 迭代数组 */
	nums := []int{2,3,4}
	sum := 0

	for _,num := range nums{
		sum += num
	}
	fmt.Println("sum:",sum) //sum: 9

	for i,num := range nums {
		if(num == 3){
			fmt.Println("index:",i)
		}
	}
	//index: 1

	/* 迭代map */
	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n",k,v)
	}
	/**
		a -> apple
		b -> banana
	*/
	/* 迭代字符串(unicode) */
	for i,c := range "go" {
		fmt.Println(i,c)
	}
	/**
		0 103
		1 111
	*/
}