package main

import (
	"time"

	"github.com/Alexey492/ticTacToe/logic"
)

func main() {
	var m logic.MapInterface = &logic.Map{}
	m.InitMap(logic.Size)
	m.PrintMap()

	for {
		m.HumanTurn()
		if m.IsGameFinished() {
			time.Sleep(5 * time.Second)
			break
		}
		m.ComputerAmatteurModeTurn()
		if m.IsGameFinished() {
			time.Sleep(10 * time.Second)
			break
		}
	}
}
