package main

import (
	"bufio"
	"os"

	"github.com/hazubeep/number-guessing-game/game"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	game.StartGame(scanner)

}
