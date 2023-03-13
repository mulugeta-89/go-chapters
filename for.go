package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var print = fmt.Println
func main(){
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(51)
	for true {
		print("The random number is ", randomNum)
		print("Enter your guess between 1 and 50:- ")
		reader := bufio.NewReaderSize(os.Stdin, 1024)
		// var guess string
		// fmt.Scanln(&guess)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		intGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if intGuess > randomNum {
			print("guess lower")
		} else if intGuess < randomNum {
			print("guess higher ")
		} else {
			print("You have guessed it ")
			break
		}

	}
}