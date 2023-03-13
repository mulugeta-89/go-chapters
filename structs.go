package main

import (
	"fmt"
)
var print = fmt.Println
type customer struct {
	name string
	address string
	dept int
} 
type rectangle struct {
	length, width int
}
func (r rectangle) area() int {
	return r.length * r.width
}
func displayAddress(s customer) string {
	return s.address
}
func changeAddress(s *customer) {
	s.address = "Wuchale"
}
func main(){
	var s1 customer
	s1.name = "Mulugeta"
	s1.address = "Addis Ababa"
	s1.dept = 20
	print(displayAddress(s1))
	changeAddress(&s1)
	print("after change")
	print(displayAddress(s1))
	print(s1)

	r1 := rectangle{2,3}
	fmt.Printf("the area for r1 is %d", r1.area())
 
}