package utils

import (
	"os"

	"golang.org/x/term"
)

func TerminalSize() (int, int) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0, 0
	}
	return w, h
}
