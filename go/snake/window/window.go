package window

import (
	"fmt"
	"snake/consts"
	"snake/player"
)

// Y,X
type Vec2D [][]string

func (window *Vec2D) InitVec2D(Xsize, Ysize int, snake player.Snake) {
	*window = make([][]string, Ysize)
	for y := range *window {
		(*window)[y] = make([]string, Xsize)
		for x := range (*window)[y] {
			(*window)[y][x] = consts.BLANK
		}
	}

	spawnSnake(*window, snake)
}

func (window Vec2D) Render() {
	for y := range window {
		for x := range window[y] {
			fmt.Print(string(window[y][x]))
			if len(window[y])-1 == x {
				fmt.Println()
			}
		}
	}
}

func (window Vec2D) Update(snake player.Snake, alive *bool) {

	// nm := snake.Move()
	nm := snake.RandomNextMove()
	collide := isCollide(window, nm)
	if collide {
		*alive = false
		return
	}

	head := len(snake) - 1

	// remove last tail from window
	y1, x1 := snake[0][0], snake[0][1]
	window[y1][x1] = consts.BLANK

	if len(snake) > 2 {
		// switch prev head with body
		yh, xh := snake[head][0], snake[head][1]
		window[yh][xh] = consts.BODY
	}

	// update snake
	_ = append(snake[:0], snake[1:]...)

	// update new head to snake
	y, x := nm[0], nm[1]
	snake[head] = []int{y, x}

	// add new tail to window
	y2, x2 := snake[0][0], snake[0][1]
	window[y2][x2] = consts.TAIL

	// add new head to window
	y3, x3 := snake[head][0], snake[head][1]
	window[y3][x3] = consts.HEAD
}

func WindowSize(window Vec2D) (int, int) {
	y, x := len(window)-1, len(window[0])-1
	return y, x
}

func spawnSnake(window Vec2D, snake player.Snake) {
	for i, snk := range snake {
		y, x := snk[0], snk[1]
		if i == len(snake)-1 {
			window[y][x] = consts.HEAD
		} else if i == 0 {
			window[y][x] = consts.TAIL
		} else {
			window[y][x] = consts.BODY
		}
	}
}

func isCollide(window Vec2D, coord [2]int) bool {
	// snake coord
	y1, x1 := coord[0], coord[1]

	Ymax, Xmax := WindowSize(window)
	collideWithWall := y1 > Ymax || x1 > Xmax
	if collideWithWall {
		return true
	}

	collideWithSelf := window[y1][x1] == consts.TAIL || window[y1][x1] == consts.BODY
	return collideWithSelf
}
