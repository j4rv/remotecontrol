package main

import (
	"github.com/go-vgo/robotgo"
)

func moveMouse(x, y int) error {
	robotgo.MoveRelative(x, y)
	return nil
}

func mouseClick(button string) error {
	robotgo.MouseClick(button)
	return nil
}
