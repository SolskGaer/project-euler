package main

import (
	"fmt"
	"math"
	"./util"
)

func main(){
	num := 600851475143
	largestfactor := 0
	squareroot := int(math.Sqrt(float64(num)))
	for i := 1; i <= squareroot; i++{
		if num % i == 0{
			if (util.IsPrime(i)){
				largestfactor = i
			}
			if (util.IsPrime(num / i)){
				largestfactor = num / i
				break
			}
		}
	}
	fmt.Println(largestfactor)
}

