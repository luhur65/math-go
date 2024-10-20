package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var name string
	var num1, num2, answer int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Math Quiz By Golang")
	fmt.Println("Welcome, ", name)
	fmt.Println()
	
	for i := 1; i <= 10; i++ {

		num1 = rand.Intn(10)
		num2 = rand.Intn(10)

		result := num1 + num2
		fmt.Println(i,". ",num1," + ",num2," = ")
		fmt.Print("Jawaban : ")
		fmt.Scan(&answer)

		if result == answer {
			fmt.Println("Benar!")
		} else {
			fmt.Println("Salah!")
			fmt.Println("Jawaban benar = ", result)
		}

		fmt.Println()

	}

	fmt.Println("Hebat!, kamu berhasil")
	fmt.Scan()

}