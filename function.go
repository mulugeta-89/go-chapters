package main

import (
	"fmt"
)

var print = fmt.Println
func sayHello() string {
	return "Hello world"
}

func getSum(x int, y int) int {
	return x+y
}
func getTwo(x int) (int, int) {
	return x+1, x+2
}
func getQuotient(x float64, y float64) (ans float64, err error){
	if y== 0 {
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		return x/y, nil
	}
}
func variadicFunc(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
func passArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main(){
	print(sayHello())
	print(getSum(1,2))
	print(getTwo(5))
	print(getQuotient(1,2))
	print(variadicFunc(1,2,3,4,5))
	nums := []int{1,2,3,4}
	print(passArray(nums))

}