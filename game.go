package main
import (
	"fmt"
	"math/rand"
	"time"
	)

func main(){

	ran := numGenerate()

	var input int
	fmt.Printf("Guess a number between 1-5: \n")
	fmt.Scanf("%d", &input)

	if ran == input {
		fmt.Printf("Your correct %d is the right number", input)
	}else {
		fmt.Printf("Your wrong %d is not the generated number, %d is", input, ran)
	}
}

func numGenerate() int {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(6)
	return num
}