package util

import (
	"errors"
)

func Gcd(num1, num2 int) (int, error){
	if num1 < num2{
		num1, num2 = num2, num1
	}
	if num2 == 0{
		return 0,errors.New("there shouldn't be zero")
	}
	for num1 % num2 != 0{
		num1, num2 = num2, num1 % num2
	}
	return num2,nil
}

func Lcm(num1, num2 int) (int, error){
	if num1 == 0 || num2 == 0{
		return 0,nil
	}
	gcdt, ok := Gcd(num1, num2)
	if ok != nil{
		return -1,ok
	}
	return num1 / gcdt * num2, nil
}
