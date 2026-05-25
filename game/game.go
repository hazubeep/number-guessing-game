package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
)

type LevelDifficulty int

const (
	Easy LevelDifficulty = iota + 1
	Medium
	Hard
)

func (l LevelDifficulty) String() string {
	switch l {
	case Easy:
		return "Easy"
	case Medium:
		return "Medium"
	case Hard:
		return "Hard"
	default:
		return "Unknown"
	}
}

type Game struct {
	TargetNumber int
	Attempts     int
	MaxAttempts  int
	Difficulty   LevelDifficulty
}

type LevelConfig struct {
	MaxAttempts int
}

func NewGame(difficulty LevelDifficulty) Game {
	config := GetLevelConfig(difficulty)

	game := Game{
		TargetNumber: rand.Intn(100),
		Attempts:     0,
		MaxAttempts:  config.MaxAttempts,
		Difficulty:   difficulty,
	}

	return game
}

func GetLevelConfig(level LevelDifficulty) LevelConfig {
	switch level {
	case Easy:
		return LevelConfig{MaxAttempts: 10}
	case Medium:
		return LevelConfig{MaxAttempts: 5}
	case Hard:
		return LevelConfig{MaxAttempts: 3}
	default:
		return LevelConfig{}
	}
}

func StartGame(scanner *bufio.Scanner) {
	var game Game

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("Your number of chances depends on the selected difficulty.")

	fmt.Printf("\n")

	fmt.Println("Please select a difficulty level:")
	fmt.Println("1. Easy   (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard   (3 chances)")

	for {

		fmt.Printf("\nEnter your choice: ")

		input, exit := readInput(scanner)

		if exit {
			return
		}

		num, err := strconv.Atoi(input)

		if err != nil || num < 1 || num > 3 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		game = NewGame(LevelDifficulty(num))

		fmt.Printf("\nGreat! You selected the %s difficulty level.\n", game.Difficulty)
		fmt.Printf("Let's start the game!\n\n")

		break
	}

	for {

		if game.Attempts >= game.MaxAttempts {
			fmt.Printf("Game over! The correct number was %d.\n", game.TargetNumber)
			break
		}

		fmt.Printf("Attempt %d/%d - Enter your guess: ", game.Attempts+1, game.MaxAttempts)

		input, exit := readInput(scanner)

		if exit {
			return
		}

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		game.Attempts++

		if guess < game.TargetNumber {
			fmt.Printf("Incorrect! The number is greater than %d.\n\n", guess)
		} else if guess > game.TargetNumber {
			fmt.Printf("Incorrect! The number is less than %d.\n\n", guess)
		} else {
			fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts.\n", game.Attempts)
			break
		}
	}
}
