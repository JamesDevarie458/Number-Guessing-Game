package main

import (
	"fmt"
	"math/rand"
	"time"
)

// These two variables are outside any function because they were not origianlly working inside a function.
var userGuess int
var secretNumber int
var name string

// Main funtction to just call all the other functions.
func main() {
	intro()
	prompts()
	randomNumber()
	checkGuess()
	tryAgain()
}

func intro() {
	fmt.Println("Welcome to the number guessing game!")
	fmt.Println("What is your name?: ")
	fmt.Scanln(&name)
	fmt.Println("Welcome", name)
}

/*
	 The function prompts gets the user's name and their guess for a number between 1 and 10.
		The number is then stored in the userGuess variable.
*/
func prompts() {
	fmt.Println("Please guess a number between 1 and 10: ")
	fmt.Scanln(&userGuess)
	if userGuess < 1 || userGuess > 10 {
		fmt.Println("Invalid Guess.")
		prompts()
	} else {
		fmt.Println("You guessed", userGuess)
	}
}

// The function randomNumber generates a random number between 1 and 10 and sores it in the secretNumber variable.
func randomNumber() {

	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(10) + 1

}

// The function checkGuess compares the user's guess to the secret number provides the user feedback.
func checkGuess() {
	if userGuess == secretNumber {
		fmt.Println("Congratulations! You guessed the correct number.")
	} else if userGuess < secretNumber {
		fmt.Println("Too low! The secret was", secretNumber)
	} else if userGuess > secretNumber {
		fmt.Println("Too high! The secret was", secretNumber)
	} else {
		fmt.Println("LOL")
	}
}

// The function tryAgain asks the user if they want to try again and if so, calls the other functions again.
func tryAgain() {
	var response string

	fmt.Println("Would you like to try again? (yes/no)")
	fmt.Scanln(&response)

	if response == "yes" {
		prompts()
		randomNumber()
		checkGuess()
		tryAgain()
	} else if response == "no" {
		fmt.Println("Thanks for playing", name)
	} else {
		fmt.Println("Invalid response. Please enter 'yes' or 'no'")
		tryAgain()
	}
}
