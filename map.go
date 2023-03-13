package main

import (
	"fmt"
)
var print = fmt.Println
func main(){
	names := make(map[string]string)
	names["mulugeta"] = "hailegnaw"
	names["shimels"] = "kebede"
	names["alemu"] = "shibeshi"

	heroes := map[int]string{1: "Mulugeta", 2: "shimels", 3: "kebede"}
	
	for k,v := range names {
		fmt.Printf("%s - %s \n" , k, v)
	}

	print("the first one is ", heroes[1])
	print("The other ones are ")
	for k,v := range heroes {
		fmt.Printf("%d - %s \n", k, v)
	}
	delete(names, "alemu")
	for k,v := range names {
		fmt.Printf("%s - %s \n" , k, v)
	}
}