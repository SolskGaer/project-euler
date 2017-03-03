package main

import (
	"fmt"
	"log"
	"./util"
)

func main(){
	var sm = 1
	var ok error
	for i := 1; i <= 20; i++{
		sm,ok = util.Lcm(i, sm)
		if ok != nil{
			log.Fatal("error")
		}
	}
	fmt.Println(sm)
}
