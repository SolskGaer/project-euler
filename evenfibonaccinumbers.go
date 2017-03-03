package main

import (
	"fmt"
)

func main(){
	var sum int
	x := 0
	y := 1
	for y < 4000000{
		if y % 2 == 0{
			sum += y
		}
		x, y = y, x+y
	}
	fmt.Println(sum)
}
