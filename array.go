package main

import (
	"fmt"
)
var print = fmt.Println

func main(){
	arr1 := [5]int{1,2,3,4,5}
	for i := 0; i < len(arr1); i++{
		print(arr1[i])
	}
	for i, num := range arr1 {
		fmt.Printf("%d : %d \n", i, num)
	}
	arr2 := [2][2]int{
		{1,2},
		{3,4},
	}
	for i:= 0; i< len(arr2); i++ {
		for j:= 0; j < len(arr2[i]); j++{
			print(arr2[i][j])
		}
	}
}