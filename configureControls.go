package main

// Import de la bibliothèque raylib
import (
	"github.com/gen2brain/raylib-go/raylib"
)

// Fonction ConfigureControls permettant de configurer les touches du jeu
func ConfigureControls(controls *Controls) {

	// Boucle infinie pour configurer les touches
	for {

		// Efface l'arrière-plan avec la couleur blanche
		rl.ClearBackground(rl.RayWhite)

		// Affiche le titre de la configuration des touches en noir
		rl.DrawText("Configuration des touches", 200, 50, 30, rl.Black)

		// Affiche un message pour indiquer d'appuyer sur une touche
		rl.DrawText("Appuyez sur une touche pour changer", 200, 100, 20, rl.DarkGray)

		// Affiche des messages pour chaque touche à configurer
		rl.DrawText("Touche pour sauter: ", 200, 200, 20, rl.DarkGray)
		rl.DrawText("Touche pour aller à gauche: ", 200, 240, 20, rl.DarkGray)
		rl.DrawText("Touche pour aller à droite: ", 200, 280, 20, rl.DarkGray)
		rl.DrawText("Touche pour tirer: ", 200, 320, 20, rl.DarkGray)
		rl.DrawText("Touche pour revenir au menu: ", 200, 360, 20, rl.DarkGray)

		// Affiche un message pour indiquer comment revenir au menu principal
		rl.DrawText("Appuyez sur ESCAPE pour revenir au menu principal", 200, 400, 20, rl.DarkGray)

		// Termine le dessin
		rl.EndDrawing()

		// Vérifie si la touche ESCAPE est enfoncée pour sortir de la configuration
		if rl.IsKeyPressed(rl.KeyEscape) {
			return
		}

		// Vérifie quelles touches ont été pressées et les attribue aux contrôles
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
