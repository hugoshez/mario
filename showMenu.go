package main

import (
	//"fmt"
	//"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

// La fonction ShowMenu affiche le menu principal du jeu et permet à l'utilisateur
// de configurer les touches ou de quitter le jeu.
func ShowMenu(controls *Controls) bool {
	// Affiche le titre du menu principal
	rl.DrawText("Menu principal", 200, 50, 30, rl.Black)

	// Affiche un message pour indiquer comment accéder à la configuration des touches
	rl.DrawText("Appuyez sur C pour configurer les touches", 200, 100, 20, rl.DarkGray)

	// Affiche un message pour indiquer comment quitter le jeu
	rl.DrawText("Appuyez sur ECHAP pour quitter", 200, 120, 20, rl.DarkGray)

	// Vérifie si la touche "C" est enfoncée pour configurer les touches
	if rl.IsKeyPressed(rl.KeyC) {
		ConfigureControls(controls) // Appelle la fonction de configuration des touches
	}

	// Vérifie si la touche "ECHAP" est enfoncée pour quitter le menu
	if rl.IsKeyPressed(rl.KeyEscape) {
		return false // Retourne false pour indiquer que le menu doit être fermé
	}

	// Retourne true pour indiquer que le menu doit rester ouvert
	return true
}
