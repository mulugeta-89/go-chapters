package main

import (
	"fmt"
)
var print = fmt.Println
type myInterfaces interface {
	run()
	eat()
}
type goat struct {
	name string
	age int
}
type human struct {
	name string
	gender string
}
func (g goat) run() string {
	return "The goat is running" + g.name
}
func (h human) run() string {
	return "The human is running" + h.name
}
func (g goat) eat() string {
	return "The goat is eating" + g.name
}
func (h human) eat() string {
	return "The human is eating" + h.name
}
func main() {
	g1 := goat{"kasa", 12}
	h1 := human{"Mulugeta", "male"}
	print(g1.run())
	print(g1.eat())
	print(h1.run())
	print(h1.eat())
}
