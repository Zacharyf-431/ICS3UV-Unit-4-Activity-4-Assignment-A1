/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-12-10
	* This file plays a guessing game with the user 
	*/

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
		// assign constants 
		const Min = 1
		const Max = 10
		var randomNum = rand.Intn(Max-Min+1) + Min
		var guess = -1
		reader := bufio.NewReader(os.Stdin)

		// greets user 
		fmt.Println("Welcome to the Guessing Game!")
		fmt.Println("Guess a number between 1 and 10. Enter 0 to quit.")


		// INPUT
		for guess != 0 && guess != randomNum {
			fmt.Print("Enter your guess: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			guess, _ = strconv.Atoi(input)

			// OUTPUT
			if guess == 0 {
				fmt.Println("Game terMinated.")
			} else if guess == randomNum {
				fmt.Printf("Congratulations! You guessed the correct number: %d\n", randomNum)
			} else {
				fmt.Println("Try again.")
			}
		}

		fmt.Println("\nDone.")
}