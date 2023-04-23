package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	cls "github.com/inancgumus/screen"
	"golang.org/x/term"
)

// Y,X
type vec2D [][]string
type snake [][]int

const (
	BLANK = "."
	SPEED = 100
	HEAD  = "$"
	BODY  = "="
	TAIL  = "+"
)

func main() {
	var screen vec2D
	alive := true
	snake := snake{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{0, 5},
		{0, 6},
		{0, 7},
		{0, 8},
	}

	w, h := terminalSize()
	screen.initVec2D(w/2, h/2, snake)

	// clear terminal
	cls.Clear()

	// main loop
	for alive {
		// clear previous render
		cls.MoveTopLeft()

		screen.update(snake, &alive)
		screen.render()

		// speed
		time.Sleep(SPEED * time.Millisecond)
	}

	fmt.Println("Game Over!")
}

func (screen *vec2D) initVec2D(Xsize, Ysize int, snake snake) {
	*screen = make([][]string, Ysize)
	for y := range *screen {
		(*screen)[y] = make([]string, Xsize)
		for x := range (*screen)[y] {
			(*screen)[y][x] = BLANK
		}
	}

	addSnakeToScreen(*screen, snake)
}

func (screen vec2D) render() {
	for y := range screen {
		for x := range screen[y] {
			fmt.Print(string(screen[y][x]))
			if len(screen[y])-1 == x {
				fmt.Println()
			}
		}
	}
}

func (screen vec2D) update(snake snake, alive *bool) {
	// get next possible move
	nm := nextMove(snake, screen)
	if len(nm) == 0 {
		*alive = false
		return
	}

	// remove tail
	y, x := snake[0][0], snake[0][1]
	screen[y][x] = BLANK
	// switch prev head with body
	yh, xh := snake[len(snake)-1][0], snake[len(snake)-1][1]
	screen[yh][xh] = BODY

	// add new head to snake
	y1, x1 := nm[0], nm[1]
	_ = append(snake[:0], snake[1:]...)
	head := len(snake) - 1
	snake[head] = []int{y1, x1}

	// update tail
	y2, x2 := snake[0][0], snake[0][1]
	screen[y2][x2] = TAIL

	// add new head to screen
	y3, x3 := snake[head][0], snake[head][1]
	screen[y3][x3] = HEAD
}

func addSnakeToScreen(screen vec2D, snake snake) {
	for i, snk := range snake {
		y, x := snk[0], snk[1]
		if i == len(snake)-1 {
			screen[y][x] = HEAD
		} else if i == 0 {
			screen[y][x] = TAIL
		} else {
			screen[y][x] = BODY
		}
	}
}

func nextMove(snake snake, screen vec2D) []int {
	var possibility [][]int
	snakeLength := len(snake)
	head := snake[snakeLength-1]

	// check possible move not include negative coordinate
	y1 := [2]int{head[0] - 1, head[1]}
	if y1[0] >= 0 && y1[1] >= 0 {
		possibility = append(possibility, y1[:])
	}

	x1 := [2]int{head[0], head[1] - 1}
	if x1[0] >= 0 && x1[1] >= 0 {
		possibility = append(possibility, x1[:])
	}

	y2 := [2]int{head[0] + 1, head[1]}
	if y2[0] >= 0 && y2[1] >= 0 {
		possibility = append(possibility, y2[:])
	}

	x2 := [2]int{head[0], head[1] + 1}
	if x2[0] >= 0 && x2[1] >= 0 {
		possibility = append(possibility, x2[:])
	}

	y, x := screenSize(screen)
	validNextMove := [][]int{}
	// check for collision with wall and self
	for _, coord := range possibility {
		y1, x1 := coord[0], coord[1]

		collideWithWall := y1 > y || x1 > x
		if collideWithWall {
			continue
		}

		collideWithSelf := screen[y1][x1] == TAIL || screen[y1][x1] == BODY
		if collideWithSelf {
			collideWithSelf = false
			continue
		}

		validNextMove = append(validNextMove, []int{y1, x1})
	}

	nValidNextMove := len(validNextMove)
	if nValidNextMove == 0 {
		return []int{}
	}

	if nValidNextMove < 2 {
		return validNextMove[0]
	}

	next := rand.Intn(nValidNextMove)
	return validNextMove[next]
}

// utils
func terminalSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return w, h
}

func screenSize(screen vec2D) (int, int) {
	y, x := len(screen)-1, len(screen[0])-1
	return y, x
}
