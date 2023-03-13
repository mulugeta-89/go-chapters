package main

import (
	stuff "example/project/mypackage"
	"fmt"
)
var print = fmt.Println

func main(){
	print("Hello", stuff.Name)
}