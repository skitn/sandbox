package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Guess the number!")

	fmt.Println("please input your guess.")

	rand.Seed(time.Now().Unix())
	answer := random(1, 101)

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()

		text2, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Guess the number!")
			continue
		}

		fmt.Println(fmt.Sprintf("You guessed: %d", text2))

		if text2 < answer {
			fmt.Println("Too small!")
		} else if text2 > answer {
			fmt.Println("Too big!")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
