package main

import (
	"fmt"
	"snake/consts"
	"snake/player"
	"snake/utils"
	"snake/window"
	"time"

	cls "github.com/inancgumus/screen"
)

func main() {
	var screen window.Vec2D
	// min 2 for head and tail
	snake := player.Snake{
		{0, 0},
		{0, 1},
		{0, 2},
		// {0, 3},
		// {0, 4},
		// {0, 5},
		// {0, 6},
		// {0, 7},
		// {0, 8},
	}
	alive := true

	w, h := utils.TerminalSize()
	screen.InitVec2D(int(float32(w)*consts.SCALE/0.5), int(float32(h)*consts.SCALE), snake)
	// screen.InitVec2D(10, 10, snake)

	// clear terminal
	cls.Clear()

	// main loop
	for alive {
		// clear previous render
		cls.MoveTopLeft()

		screen.Update(snake, &alive)
		screen.Render()

		// speed
		time.Sleep(consts.SPEED * time.Millisecond)
	}

	fmt.Println("Game Over!")
}
