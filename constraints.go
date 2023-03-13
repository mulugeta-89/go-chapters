package main

import (
	"fmt"
)
var print = fmt.Println

type MyContstraint interface {
	int | float64
}
func adder[M MyContstraint](a M, b M) M {
	return a +b
}
func main(){
	print(adder(1,3))
	print(adder(1.2,2.0))
}