package main
import (
	"fmt"
	"math/rand"
	"time"
	)

func main(){

	ran := numGenerate()

	var score int
	var attempt int
	var input int
	var rangeData int

	fmt.Printf("Guess a number between 1-5: \n")
	fmt.Scanf("%d", &input)

	rangeData = 1

	for i := 0; i < rangeData; i++ {
		if input == 100 {
			return 
		}else {
			if ran == input {
				score += 4
				fmt.Printf("Your correct %d is the right number. Your score is %d ", input, score)

			}else {
				attempt ++
				fmt.Printf("Your wrong %d is not the generated number, %d is %d", input, ran, attempt)
			}
		}
	}

}

func numGenerate() int {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(6)
	return num
}

// func scoreCalculator(s int, ) int {

// }