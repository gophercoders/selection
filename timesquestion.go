package main

import (
	"fmt"

	"github.com/gophercoders/random"
	"github.com/gophercoders/simpleio"
)

func main() {

	var a int
	var b int
	var answer int

	fmt.Println("The timesquestion shows you how to use if and else.")
	fmt.Println("Can you remember your times tables?")
	fmt.Println("")

	a = random.GetRandomNumberInRange(1, 12)
	b = random.GetRandomNumberInRange(1, 12)

	fmt.Print("What is ")
	fmt.Print(a)
	fmt.Print(" * ")
	fmt.Print(b)
	fmt.Println("?")

	answer = simpleio.ReadNumberFromKeyboard()

	if answer == a*b {
		fmt.Println("Congratulations! You are correct! ")
	} else if answer > a*b {
		fmt.Println("Sorry, your guess was to big.")
		fmt.Print("The correct answer is ")
		fmt.Print(a)
		fmt.Print(" * ")
		fmt.Print(b)
		fmt.Print(" = ")
		fmt.Println(a * b)
	} else {
		fmt.Println("Sorry your guess was to small.")
		fmt.Print("The correct answer is ")
		fmt.Print(a)
		fmt.Print(" * ")
		fmt.Print(b)
		fmt.Print(" = ")
		fmt.Println(a * b)
	}
	fmt.Println("Run the program again to try another question.")
}
