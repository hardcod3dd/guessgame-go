package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	getNumber("Enter a number and press ENTER: ")
}

func getNumber(s string) {
	number := rand.Intn(10)
	fmt.Print(s)

	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Enter a number")
	} else if num == number {
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("Number was %d and you guessed %d. Congratz\n", number, num)
	} else {
		fmt.Println("--------------------------------------------------------")
		fmt.Printf("You guessed %d but the number was %d. Better luck next time.\n", num, number)
	}
}
