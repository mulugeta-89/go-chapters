package main

import (
	"fmt"
	"math/rand"
	"time"
)
var print = fmt.Println
func main(){
	rand.Seed(time.Now().Unix())
	num := rand.Intn(101)
	print(num)

}