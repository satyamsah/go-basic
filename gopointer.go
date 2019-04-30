package main

import (
	"fmt"
)
/*
func main() {
	fmt.Println("Hello, playground")
	arr := []int{2,3,4,5}
	var parr *[]int
	parr= &arr
	parr1:=appendValue(parr,7)
	fmt.Println(*parr1)
	
}

func appendValue(parr *[]int, val int)(* []int){
	*parr=append(*parr,val)
	return parr
	
}
*/


func main() {
	fmt.Println("Hello, playground")
	arr := []int{2,3,4,5}
	var parr *[]int
	parr= &arr
	appendValue(parr,7)
	fmt.Println(*parr)
	
}

func appendValue(parr *[]int, val int){
	*parr=append(*parr,val)
	
}
