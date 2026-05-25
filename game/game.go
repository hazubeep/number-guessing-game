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
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Printf("\n")

	fmt.Println("Please select the difficulty level: ")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	for {

		fmt.Print("Enter Your choice: ")

		if !scanner.Scan() {
			return
		}

		if scanner.Text() == "exit" {
			fmt.Println("Goodbye")
			break
		}

		input, error := strconv.Atoi(scanner.Text())

		if error != nil || input < 1 || input > 3 {
			fmt.Println("Invalid choice, try again")
			continue
		}

		game := NewGame(LevelDifficulty(input))

		fmt.Printf("Great! You have selected the %s difficulty level.\n\n", game.Difficulty.String())
		break
	}
}
