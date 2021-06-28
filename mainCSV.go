package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	success := 0
	agree := ""
	data, err := ioutil.ReadFile("./file/data.csv")
	check(err)
	fmt.Println("Are you ready!, Press Y if Yes and N if No")
	fmt.Scanf("%s", &agree)

	if agree == "Y" {
		success = startQuiz(data, success)
		fmt.Printf("Your Score is %d\n", success)
	}

}

func startQuiz(data []byte, success int) int {

	quizAns := strings.Split(string(data), "\n")
	var quizParam []string
	var ans string
	ch := make(chan string)

	// for i, val := range quizAns {
	// 	fmt.Println(strconv.Itoa(i), val, strings.Split(val, ","))
	// }
	lenQuiz := len(quizAns)
	for i := 0; i < lenQuiz; i++ {
		go func() {
			//Making them go sequential is left
			quizParam = strings.Split(quizAns[i], ",")
			fmt.Printf("%s", quizParam[0])
			fmt.Scanf("%d", &ans)
			//input, err := strconv.Atoi(ans)
			//check(err)
			ch <- ans

		}()
		select {
		case <-ch:
			if ans == quizParam[1] {
				success += 1
			}
		case <-time.After(3 * time.Second):
			fmt.Printf("\n")

		}
	}
	return success
}
