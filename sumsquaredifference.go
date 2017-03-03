package main

import (
	"fmt"
)

func main(){
	const num = 100
	var sumofsquares, sum int
	for i := 1; i <= num; i++{
		sumofsquares += i * i
		sum += i
	}
	fmt.Println(sum * sum - sumofsquares)
}
