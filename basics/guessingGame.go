package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//generate a random number
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 - 100")
	fmt.Println("Can you guess what that number is?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations, you got guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Sorry your guess was too low, Please try a larger number!")
		} else {
			fmt.Println("Sorry your guess was too high, Please try a smaller number!")
		}
	}

}
