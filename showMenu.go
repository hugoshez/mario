package main

import (
	//"fmt"
	//"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

func ShowMenu(controls *Controls) bool {
	rl.DrawText("Menu principal", 200, 50, 30, rl.Black)
	rl.DrawText("Appuyez sur C pour configurer les touches", 200, 100, 20, rl.DarkGray)
	rl.DrawText("Appuyez sur ECHAP pour quitter", 200, 120, 20, rl.DarkGray)

	if rl.IsKeyPressed(rl.KeyC) {
		ConfigureControls(controls)
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		return false
	}

	return true
}
