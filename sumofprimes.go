package main

import (
	"fmt"
	"./util"
)

const num = 2000000

func main(){
	var sum int
	for i := 1; i <= num; i++{
		if util.IsPrime(i){
			sum += i
		}
	}
	fmt.Println(sum)
}
