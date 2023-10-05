package main

import (
	//"fmt"
	//"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

// Définition de la structure Ground, qui représente une plateforme ou un sol dans le jeu.
type Ground struct {
	posX   int32    // Position en X de la plateforme
	posY   int32    // Position en Y de la plateforme
	width  int32    // Largeur de la plateforme
	height int32    // Hauteur de la plateforme
	Color  rl.Color // Couleur de la plateforme (utilise la bibliothèque raylib)
}

// Définition de la structure Bullet, qui représente une balle tirée dans le jeu.
type Bullet struct {
	posX            int32    // Position en X de la balle
	posY            int32    // Position en Y de la balle
	radius          float32  // Rayon de la balle (utilisé pour l'affichage)
	Right_direction bool     // Indique la direction de la balle (droite ou gauche)
	Draw            bool     // Indique si la balle doit être dessinée
	Color           rl.Color // Couleur de la balle (utilise la bibliothèque raylib)
}

// Définition de la structure Enemy, qui représente un ennemi dans le jeu.
type Enemy struct {
	posX      int32    // Position en X de l'ennemi
	posY      int32    // Position en Y de l'ennemi
	velocity  int32    // Vitesse de déplacement de l'ennemi
	Damage    float32  // Dommages infligés par l'ennemi
	Direction bool     // Direction de déplacement de l'ennemi (gauche ou droite)
	Draw      bool     // Indique si l'ennemi doit être dessiné
	Color     rl.Color // Couleur de l'ennemi (utilise la bibliothèque raylib)
}
type DifferentEnemy struct {
	posX      int32
	posY      int32
	velocity  int32
	Damage    float32
	Direction bool
	Draw      bool
	Color     rl.Color
	Texture   rl.Texture2D // Nouvelle variable pour la texture de l'apparence différente
}

// Définition de la structure Controls, qui stocke les touches configurées pour le jeu.
type Controls struct {
	KeyJump  int32 // Touche associée au saut
	KeyLeft  int32 // Touche associée au déplacement vers la gauche
	KeyRight int32 // Touche associée au déplacement vers la droite
	KeyShoot int32 // Touche associée au tir
	KeyMenu  int32 // Touche associée au retour au menu
}
