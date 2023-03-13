package main

import (
	"fmt"
)
var print = fmt.Println
func main(){
	names := make([]string, 5)
	names[0] = "mulugeta"
	names[1] = "hailegnaw"
	names[2] = "belay"
	names[3] = "gebrie"
	names[4] = "feleke"

	for i:=0; i< len(names); i++{
		print(names[i])
	}
	for i, v := range names {
		fmt.Printf("%d : %s\n", i , v)
	}
	nums := [5]int{1,2,3,4,5}
	nums1 := nums[1:3]
	print(nums1)
	nums1[1] = 200
	print(nums)
	print(nums1)
	nums1 = append(nums1, 120)
	print(nums1)
	print(nums)
}