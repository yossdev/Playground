package player

import (
	"math/rand"
)

type Snake [][]int

func (snake Snake) Move() [2]int {
	// reader := bufio.NewReader(os.Stdin)
	// char, _, err := reader.ReadRune()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	var coord [2]int
	// last two position
	snk := snake[len(snake)-2:]
	y1, x1 := snk[1][0], snk[1][1]
	y2, x2 := snk[0][0], snk[0][1]

	// keys map[A:65 D:68 S:83 W:87 a:97 d:100 s:115 w:119]
	// left
	// if char == 65 || char == 97 {

	// 	return coord
	// }

	// // right
	// if char == 65 || char == 97 {

	// 	return coord
	// }
	// // down
	// if char == 65 || char == 97 {

	// 	return coord
	// }
	// // up
	// if char == 65 || char == 97 {

	// 	return coord
	// }

	autoMove := nextMove(y1, x1, y2, x2, coord)
	return autoMove
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
	return coord
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
