package main

import (
	"fmt"
	"math/rand"
	"time"
)

var userGuess int
var secretNumber int

func main() {
	prompts()
	randomNumber()
	checkGuess(userGuess, secretNumber)
}

func prompts() {
	fmt.Println("Hello, please enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Welcome", name)
	fmt.Println("Please guess a number between 1 and 10: ")
	var userGuess int
	fmt.Scanln(&userGuess)
}

func randomNumber() {

	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(10) + 1

}

func checkGuess(userGuess int, secretNumber int) {
	if userGuess == secretNumber {
		fmt.Println("Congratulations! You guessed the correct number.")
	} else if userGuess < secretNumber {
		fmt.Println("Too low! The secret was", secretNumber)
	} else {
		fmt.Println("Too high! The secret was", secretNumber)
	}
}
