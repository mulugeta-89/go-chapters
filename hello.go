package main

import (
	"fmt"
	"strings"
)

var print = fmt.Println
func main(){
	name := "The Mulugeta"
	replacer := strings.NewReplacer("The", "MotherFucker")
	name2 := replacer.Replace(name)
	print(name2)
	print("the length of name ", len(name))
	print("spliting name ", strings.Split("M-u-l-u-g-e-t-a", "-"))
	print("Index of u", strings.Index(name, "u"))
	print("replacing e in name ", strings.Replace(name, "u", "9", -1))
	print("removing space from name2", strings.TrimSpace(name2))
	print("has prefix name ", strings.HasPrefix(name, "Th"))
	print("has suffix name", strings.HasSuffix(name, "ta"))
	print("name contains The", strings.Contains(name, "The"))
	print("changing name to upper ", strings.ToUpper(name))
	print("chaning name to lower", strings.ToLower(name))
}