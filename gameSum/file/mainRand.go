package main

import (
	"fmt"
	"math/rand"
)

type equation struct {
	num1     int
	num2     int
	res      int
	inputVal int
}

func main() {

	quizLimit := 5
	exercise := make([]equation, 5, 5)
	score := 0

	//guess 5 random number from 1 to 10 and put in an array

	for i := 0; i < quizLimit; i++ {
		exercise[i].num1 = rand.Intn(10)
		exercise[i].num2 = rand.Intn(10)
		exercise[i].res = exercise[i].num1 + exercise[i].num2
		fmt.Printf("%d, %d=", exercise[i].num1, exercise[i].num2)
		fmt.Scanf("%d", &exercise[i].inputVal)
		if exercise[i].inputVal == exercise[i].res {
			score += 1
		}
	}
	fmt.Println("Your Score is:", score, "Out of ", quizLimit)

	//Take input from user

	//check answer

}
