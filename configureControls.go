package main

import (
	//"fmt"
	//"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

func ConfigureControls(controls *Controls) {

	for {
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Configuration des touches", 200, 50, 30, rl.Black)
		rl.DrawText("Appuyez sur une touche pour changer", 200, 100, 20, rl.DarkGray)
		rl.DrawText("Touche pour sauter: ", 200, 200, 20, rl.DarkGray)
		rl.DrawText("Touche pour aller à gauche: ", 200, 240, 20, rl.DarkGray)
		rl.DrawText("Touche pour aller à droite: ", 200, 280, 20, rl.DarkGray)
		rl.DrawText("Touche pour tirer: ", 200, 320, 20, rl.DarkGray)
		rl.DrawText("Touche pour revenir au menu: ", 200, 360, 20, rl.DarkGray)

		rl.DrawText("Appuyez sur ESCAPE pour revenir au menu principal", 200, 400, 20, rl.DarkGray)

		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyEscape) {
			return
		}

		if rl.IsKeyPressed(rl.KeyW) {
			controls.KeyJump = rl.GetKeyPressed()
		}

		if rl.IsKeyPressed(rl.KeyA) {
			controls.KeyLeft = rl.GetKeyPressed()
		}

		if rl.IsKeyPressed(rl.KeyD) {
			controls.KeyRight = rl.GetKeyPressed()
		}

		if rl.IsKeyPressed(rl.KeyF) {
			controls.KeyShoot = rl.GetKeyPressed()
		}

		if rl.IsKeyPressed(rl.KeyM) {
			controls.KeyMenu = rl.GetKeyPressed()
		}
	}
}
