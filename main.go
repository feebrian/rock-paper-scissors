package main

import (
	"fmt"
	"math/rand"
)

func randomizer() string {
	choice := [3]string{"Rock", "Paper", "Scissors"}
	i := rand.Intn(3)

	return choice[i]
}

func checkAnswer(input string) {
	bot := randomizer()

	switch {
	case bot == "Rock" && input == "Rock":
		fmt.Println(bot, "vs", input)
		fmt.Println("Draw")
	case bot == "Rock" && input == "Paper":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Win!")
	case bot == "Rock" && input == "Scissors":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Lose")
	case bot == "Paper" && input == "Paper":
		fmt.Println(bot, "vs", input)
		fmt.Println("Draw")
	case bot == "Paper" && input == "Rock":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Lose!")
	case bot == "Paper" && input == "Scissors":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Win")
	case bot == "Scissors" && input == "Scissors":
		fmt.Println(bot, "vs", input)
		fmt.Println("Draw")
	case bot == "Scissors" && input == "Paper":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Lose!")
	case bot == "Scissors" && input == "Rock":
		fmt.Println(bot, "vs", input)
		fmt.Println("You Win")
	default:
		fmt.Println("Please input valid value!")
	}
}

func main() {
	fmt.Println("Choose Rock, Paper, or Scissors!")
	var input string
	fmt.Scan(&input)
	checkAnswer(input)
}
