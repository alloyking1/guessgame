package main
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	)

func main(){
	fmt.Printf("Guess a number between 1-5: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	ran := numGenerate()

	if ran == input {
		fmt.Printf("Your correct %q is the right number", input)
	}else {
		fmt.Printf("Your wrong %q is not the generated number", input)
	}
}

func numGenerate() int64 {
	var num int64
	num = 22
	return num
}