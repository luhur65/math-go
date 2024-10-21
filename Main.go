package main

import (
	"fmt"
	"math-go/game"
)

func main() {

	var name string
	// var num1, num2, answer int
	var answer int
	var soal int = 1

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Math Quiz By Golang")
	fmt.Println("Welcome, ", name)
	fmt.Println()

	game.MathQuiz(soal, answer)

	fmt.Println("Hebat!, kamu berhasil")
	fmt.Scan()

}