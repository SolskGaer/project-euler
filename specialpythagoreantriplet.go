package main

import (
	"fmt"
)

const num = 1000

func main(){
	for i := num; i > 0; i--{
		for j := 1; j < num - i; j++{
			k := num - i - j
			if isTriplet(i, j, k){
				fmt.Println(i*j*k)
				return
			}
		}
	}
}

func isTriplet(i, j, k int) bool{
	if i > j && i > k{
		if i * i == j * j + k * k{
			return true
		}
	}
	if j > i && j > k{
		if j * j == i * i + k * k{
			return true
		}
	}
	if k > i && k > j{
		if k * k == i * i + j * j{
			return true
		}
	}
	return false
}
