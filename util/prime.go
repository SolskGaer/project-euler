package util

import (
	"math"
)

func IsPrime(num int) bool{
	if num == 2{
		return true
	}
	if num == 1 || num == 0 || num % 2 == 0{
		return false
	}
	var squareroot = int(math.Sqrt(float64(num)))
	for i := 3; i <= squareroot; i += 2{
		if num % i == 0{
			return false
		}
	}
	return true
}
