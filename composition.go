package main

import (
	"fmt"
)
var print = fmt.Println

type student struct {
	fname string
	lname string
	address
}
type address struct {
	addre string
	phone string
}
func (s student) show() {
	fmt.Printf("student %s %s lives in %s and has phone number %s\n", s.fname, s.lname, s.address.addre, s.address.phone)
}
func main(){
	ad1 := address {
		"Wuchale",
		"0928351857",
	}
	s1 := student {
		"Mulugeta",
		"hailengaw",
		ad1,
	}
	
	s1.show()
}