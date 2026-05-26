# Number Guessing Game

A simple CLI-based number guessing game built with Go, where the computer randomly selects a number and the player has to guess it within a limited number of chances based on the chosen difficulty level.

Project reference: https://roadmap.sh/projects/number-guessing-game

## Features

- Three difficulty levels: Easy, Medium, and Hard
- Tracks the number of attempts per round
- Hints on each incorrect guess indicating whether the target number is higher or lower
- Graceful exit support (Ctrl+D / EOF)

## Difficulty Levels

| Level  | Chances |
|--------|---------|
| Easy   | 10      |
| Medium | 5       |
| Hard   | 3       |

## Requirements

- Go 1.21 or later

## Installation

Clone the repository:

```bash
git clone https://github.com/hazubeep/number-guessing-game.git
cd number-guessing-game
```

## Usage

Run the game directly using `go run`:

```bash
go run main.go
```

Or build the binary first, then run it:

```bash
go build -o bin/number-guessing-game .
./bin/number-guessing-game
```

## Sample Output

```
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
Your number of chances depends on the selected difficulty.

Please select a difficulty level:
1. Easy   (10 chances)
2. Medium (5 chances)
3. Hard   (3 chances)

Enter your choice: 2

Great! You selected the Medium difficulty level.
Let's start the game!

Attempt 1/5 - Enter your guess: 50
Incorrect! The number is less than 50.

Attempt 2/5 - Enter your guess: 25
Incorrect! The number is greater than 25.

Attempt 3/5 - Enter your guess: 35
Incorrect! The number is less than 35.

Attempt 4/5 - Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
```

## Project Structure

```
number-guessing-game/
├── game/
│   ├── game.go       # Core game logic, difficulty config, and game loop
│   └── helper.go     # Input reading helper
├── bin/              # Compiled binaries (generated on build)
├── main.go           # Entry point
├── go.mod            # Go module definition
└── README.md
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
