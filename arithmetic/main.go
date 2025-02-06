package main

import (
	"fmt"
	"math"
)

func main(){
	var a int = 1 + 1
	var b int = 2 * 2
	var c int = int(math.Pow(float64(a), float64(b)))
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}