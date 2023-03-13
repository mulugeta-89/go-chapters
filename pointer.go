package main

import (
	"fmt"
)

var print = fmt.Println

func changeVal(num *int) {
	*num = 12
}
func main(){
	num1 := 20
	print(num1)
	changeVal(&num1)
	print(num1)

	num := 20
	var num2 *int = &num
	print("num2 : ", *num2)
	*num2 = 21
	print("num2: ", *num2)
	print(num)

}