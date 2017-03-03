package main

import (
	"fmt"
	"./util"
)

const num = 10001

func main(){
	var count = 1
	for i := 1; ; i++{
		if util.IsPrime(i){
			count++
			if count == num{
				fmt.Println(i)
				break
			}
		}
	}
}
