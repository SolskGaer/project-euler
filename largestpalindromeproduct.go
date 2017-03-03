package main

import (
	"fmt"
)

func main(){
	var product int
	for i := 999; i >= 100; i--{
		for j := 999; j >= 100; j--{
			producttemp := i * j
			if isPalindrome(producttemp) && producttemp > product{
				product = producttemp
			}
		}
	}
	fmt.Println(product)
}

func isPalindrome(num int) bool{
	str := fmt.Sprintf("%d", num)
	for i := 0; i < len(str); i++{
		if str[i] != str[len(str)-1-i]{
			return false
		}
	}
	return true
}
