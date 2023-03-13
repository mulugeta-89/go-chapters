package stuff

import "strconv"

var Name string = "Mulugeta"

func IntTostr(nums []int) []string {
	var newString []string
	for _, val := range nums {
		newString = append(newString, strconv.Itoa(val))
	}
	return newString
}