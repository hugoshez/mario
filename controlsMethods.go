package main

import (
	//"fmt"
	//"strconv"
	"github.com/gen2brain/raylib-go/raylib"
)

// Fonction Init initialise les touches par défaut pour la structure Controls
func (c *Controls) Init() {
	c.KeyJump = rl.KeySpace
	c.KeyLeft = rl.KeyLeft
	c.KeyRight = rl.KeyRight
	c.KeyShoot = rl.KeyF
	c.KeyMenu = rl.KeyG
}

// Méthode setKeyJump permet de configurer la touche pour sauter
func (c *Controls) setKeyJump(touch int32) {
	c.KeyJump = touch
}

// Méthode setKeyLeft permet de configurer la touche pour aller à gauche
func (c *Controls) setKeyLeft(touch int32) {
	c.KeyLeft = touch
}

// Méthode setKeyRight permet de configurer la touche pour aller à droite
func (c *Controls) setKeyRight(touch int32) {
	c.KeyRight = touch
}

// Méthode setKeyShoot permet de configurer la touche pour tirer
func (c *Controls) setKeyShoot(touch int32) {
	c.KeyShoot = touch
}

// Méthode setKeyMenu permet de configurer la touche pour revenir au menu
func (c *Controls) setKeyMenu(touch int32) {
	c.KeyMenu = touch
}
