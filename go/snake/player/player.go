package player

import (
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/term"
)

type Snake [][]int

func (snake Snake) Move() [2]int {
	var coord [2]int
	// last two position
	snk := snake[len(snake)-2:]
	y1, x1 := snk[1][0], snk[1][1]
	y2, x2 := snk[0][0], snk[0][1]

	var input string
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	char := make([]byte, 1)
	_, err = os.Stdin.Read(char)
	if err != nil {
		fmt.Println(err)
	}
	input = string(char[0])
	// fmt.Printf("the char %q was hit", string(char[0]))

	if input == "A" || input == "a" {
		coord[0] = y1
		coord[1] = x1 - 1
	} else if input == "D" || input == "d" {
		coord[0] = y1
		coord[1] = x1 + 1
	} else if input == "S" || input == "s" {
		coord[0] = y1 + 1
		coord[1] = x1
	} else if input == "W" || input == "w" {
		coord[0] = y1 - 1
		coord[1] = x1
	} else {
		coord = nextMove(y1, x1, y2, x2, coord)
	}

	return coord
}

func nextMove(y1, x1, y2, x2 int, coord [2]int) [2]int {
	nm := coord[:]
	if x1 > x2 && y1 == y2 { // x to right
		nm[1] = x1 + 1
	} else if x1 < x2 && y1 == y2 { // x to left
		nm[1] = x1 + 1
	} else if y1 < y2 && x1 == x2 { // y to up
		nm[0] = y1 - 1
	} else if y1 > y2 && x1 == x2 { // y to down
		nm[0] = y2 + 1
	}
	return [2]int{nm[0], nm[1]}
}

func (snake Snake) RandomNextMove() [2]int {
	var possibility [][]int
	snakeLength := len(snake)
	head := [2]int{snake[snakeLength-1][0], snake[snakeLength-1][1]}
	body := [2]int{snake[snakeLength-2][0], snake[snakeLength-2][1]}

	// check possible move not include negative coordinate
	y1 := [2]int{head[0] - 1, head[1]}
	if y1[0] >= 0 && y1[1] >= 0 && y1 != body {
		possibility = append(possibility, y1[:])
	}

	x1 := [2]int{head[0], head[1] - 1}
	if x1[0] >= 0 && x1[1] >= 0 && x1 != body {
		possibility = append(possibility, x1[:])
	}

	y2 := [2]int{head[0] + 1, head[1]}
	if y2[0] >= 0 && y2[1] >= 0 && y2 != body {
		possibility = append(possibility, y2[:])
	}

	x2 := [2]int{head[0], head[1] + 1}
	if x2[0] >= 0 && x2[1] >= 0 && x2 != body {
		possibility = append(possibility, x2[:])
	}

	next := rand.Intn(len(possibility))
	res := [2]int{possibility[next][0], possibility[next][1]}
	return res
}
