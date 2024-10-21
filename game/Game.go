package game

import (
	"fmt"
	"math/rand"
)

func checkAnswer(answer, result int) string {
	if result == answer {
		return "Benar!"
	} else {
		return "Salah!"
	}
}

func MathQuiz(soal, answer int) {

	var is_again string

	// cek jika sudah 10 soal
	if soal > 10 {
		fmt.Print("Mau main lagi?[Y/N] : ")
		fmt.Scan(&is_again)
	}

	if is_again == "n" || is_again == "N" {
		fmt.Println("Terima kasih sudah bermain ya!")
		panic("Program selesai")
	}

	if is_again == "y" || is_again == "Y" {
		soal = 1 // reset lagi soal
		MathQuiz(soal, answer)
	}

	num1 := rand.Intn(10)
	num2 := rand.Intn(10)
	result := num1 + num2

	fmt.Println(soal,". ",num1," + ",num2," = ")
	fmt.Print("Jawaban : ")
	fmt.Scan(&answer)

	fmt.Println(checkAnswer(answer, result))

	fmt.Println("Jawaban benar = ", result)
	fmt.Println()
	soal += 1

	MathQuiz(soal, answer)
}
