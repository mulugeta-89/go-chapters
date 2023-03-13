package main

import (
	"fmt"
)

var print = fmt.Println
func doubleArr(arr *[5]int){
	for i:=0; i < 5; i++ {
		arr[i] *= 2
	}
}
func getAverage(sli ...float64) float64 {
	var sum float64 = 0
	var length float64 = float64(len(sli))
	for i:=0; i < len(sli); i++ {
		sum += sli[i]
	}
	return sum/length
}
func main(){
	arr := [5]int{1,2,3,4,5}
	print(arr)
	doubleArr(&arr)
	print(arr)

	sli := []float64{1.3,2.5,3.3,4.4,5.5}
	print(getAverage(sli ...))

}