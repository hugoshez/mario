package main

import (
	"fmt"
	"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

func main() {

	// ------------------ MODIF ANTOINE -----------------------
	var controls Controls
	controls.Init()
	// ---------------------------------------------------------

	// Définition de la largeur et de la hauteur de la fenêtre du jeu
	screenWidth := int32(800)
	screenHeight := int32(450)

	// Initialisation de la fenêtre avec les dimensions et le titre
	rl.InitWindow(screenWidth, screenHeight, "Mario Terroriste")

	// Initialisation des variables de santé et de score
	var Health int = 100
	var Score int = 0

	// Création des tableaux pour stocker les éléments du jeu
	grounds := []Ground{}
	bullets := []Bullet{}

	// Création d'objets de sol et de table
	ground := Ground{0, 420, screenWidth, 30, rl.Brown}
	table := Ground{screenWidth / 2, screenHeight / 2, screenWidth / 4, 30, rl.Brown}

	// Ajout des objets au tableau des éléments du jeu
	grounds = append(grounds, ground)
	grounds = append(grounds, table)

	// Chargement des images des personnages et des ennemis
	Character := rl.LoadImage("assets/mario.png")
	LeftCharacter := rl.LoadImage("assets/leftmario.png")
	Enemies := []Enemy{}
	RightGoomba := rl.LoadImage("assets/leftgoomba.png")
	LeftGoomba := rl.LoadImage("assets/rightgoomba.png")

	// Création d'un ennemi initial et ajout au tableau des ennemis
	first_enemy := Enemy{0, 370, 5, 1, true, true, rl.White}
	Enemies = append(Enemies, first_enemy)

	// Chargement de la texture du personnage
	Enemy_texture := rl.LoadTextureFromImage(RightGoomba)
	texture := rl.LoadTextureFromImage(Character)

	// Coordonnées initiales du personnage
	var x_coords int32 = screenWidth/2 - texture.Width/2
	var y_coords int32 = screenHeight/2 - texture.Height/2 - 40

	// Variable pour gérer le tir
	var should_shoot bool = true

	// Définition de la fréquence de rafraîchissement de l'écran
	rl.SetTargetFPS(60)

	// Variable pour suivre la direction du personnage
	var isRight bool = true

	// Variable pour afficher ou masquer le menu
	showMenu := false

	// Boucle principale du jeu
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Si showMenu est vrai, affiche le menu de configuration et attend une réponse
		if showMenu {
			showMenu = ShowMenu(&controls)
		}

		// Affiche la santé du joueur en fonction de sa valeur
		fmt.Println(should_shoot)
		if Health > 50 {
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Green)
		} else if Health < 50 && Health > 10 {
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Orange)
		} else {
			rl.DrawText("Health: "+strconv.Itoa(Health), 0, 0, 20, rl.Red)
		}

		// Affiche le score du joueur
		rl.DrawText("Score: "+strconv.Itoa(Score), 2, 300, 20, rl.Blue)

		// Affiche la texture du personnage à ses coordonnées actuelles
		rl.DrawTexture(texture, x_coords, y_coords, rl.White)

		// Vérifie les collisions avec le sol et ajuste les coordonnées du personnage
		for _, current_ground := range grounds {
			rl.DrawRectangle(current_ground.posX, current_ground.posY, current_ground.width, current_ground.height, current_ground.Color)
			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords), float32(50), float32(50)), rl.NewRectangle(float32(current_ground.posX), float32(current_ground.posY), float32(current_ground.width), float32(current_ground.height))) {
				y_coords = current_ground.posY - 50
				if rl.IsKeyDown(controls.KeyJump) {
					y_coords -= 300
				}
			}
		}

		// Gestion des ennemis
		for index, current_enemy := range Enemies {
			if Enemies[index].Draw {
				// Gestion du mouvement de l'ennemi
				if current_enemy.posX > screenWidth {
					Enemies[index].velocity = -5
				}
				if current_enemy.posX < 0 {
					Enemies[index].velocity = 5
				}
				if current_enemy.Direction {
					Enemy_texture = rl.LoadTextureFromImage(RightGoomba)
					Enemies[index].posX = int32(Enemies[index].posX + Enemies[index].velocity)
					Enemies[index].Direction = false
				} else {
					Enemy_texture = rl.LoadTextureFromImage(LeftGoomba)
					Enemies[index].Direction = true
				}

				// Vérifie les collisions avec l'ennemi et ajuste la santé du joueur
				if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_coords), float32(y_coords), float32(50), float32(50)), rl.NewRectangle(float32(current_enemy.posX), float32(current_enemy.posY), float32(50), float32(50))) {
					Health = Health - int(current_enemy.Damage)
					fmt.Println("detect")
				}

				// Affiche l'ennemi avec sa texture
				rl.DrawTexture(Enemy_texture, current_enemy.posX, current_enemy.posY, current_enemy.Color)

				// Gestion des balles
				for index1, current_bullet := range bullets {
					if rl.CheckCollisionRecs(rl.NewRectangle(float32(current_bullet.posX), float32(current_bullet.posY), float32(current_bullet.radius), float32(current_bullet.radius)), rl.NewRectangle(float32(current_enemy.posX), float32(current_enemy.posY), float32(50), float32(50))) {
						// Si une balle touche l'ennemi, le désactive et crée de nouveaux ennemis
						Enemies[index].Draw = false
						new_enemy1 := Enemy{0, 370, current_enemy.velocity + 1, current_enemy.Damage * 2, true, true, rl.White}
						new_enemy2 := Enemy{800, 370, current_enemy.velocity + 1, current_enemy.Damage * 2, true, true, rl.White}
						Enemies = append(Enemies, new_enemy1)
						Enemies = append(Enemies, new_enemy2)
						Score++
						bullets[index1] = Bullet{}
						should_shoot = true
					}
					if current_bullet.posX < 0 || current_bullet.posX > screenWidth {
						// Si une balle sort de l'écran, la désactive
						if current_bullet.Draw {
							fmt.Println(current_bullet.posX)
							fmt.Println(current_bullet.posY)
							bullets[index1].Draw = false
							if should_shoot {
								should_shoot = false
							} else {
								should_shoot = true
							}
						} else {
							current_bullet.Draw = false
						}
					} else {
						// Déplace la balle en fonction de sa direction
						if current_bullet.Right_direction {
							bullets[index1].posX = int32(bullets[index1].posX + 5)
							rl.DrawCircle(current_bullet.posX, current_bullet.posY, current_bullet.radius, current_bullet.Color)
						} else {
							bullets[index1].posX = int32(bullets[index1].posX - 5)
							rl.DrawCircle(current_bullet.posX, current_bullet.posY, current_bullet.radius, current_bullet.Color)
						}
					}
				}
			}
		}

		// Si le personnage est au-dessus du sol, ajuste sa position vers le bas
		if y_coords+45 < ground.posY {
			if y_coords > ground.posY {
			} else {
				y_coords += 5
			}
		}

		// Affiche ou masque le menu en fonction de la touche de menu
		if rl.IsKeyReleased(controls.KeyMenu) {
			showMenu = !showMenu
		}

		// Gestion des mouvements du personnage
		if rl.IsKeyDown(controls.KeyRight) {
			texture = rl.LoadTextureFromImage(Character)
			x_coords += 5
			isRight = true
		}
		if rl.IsKeyDown(controls.KeyLeft) {
			texture = rl.LoadTextureFromImage(LeftCharacter)
			x_coords -= 5
			isRight = false
		}

		// Gestion du tir du personnage
		if rl.IsKeyDown(controls.KeyShoot) {
			if isRight {
				if should_shoot {
					current_bullet := Bullet{int32(x_coords + 50), int32(y_coords + 25), float32(10), true, true, rl.Red}
					bullets = append(bullets, current_bullet)
					should_shoot = false
				}
			} else {
				if should_shoot {
					current_bullet := Bullet{int32(x_coords - 1), int32(y_coords + 25), float32(10), false, true, rl.Red}
					bullets = append(bullets, current_bullet)
					should_shoot = false
				}
			}
		}

		// Gestion des limites de l'écran et de la fin du jeu
		if y_coords <= 0 {
			y_coords += 15
		}
		if x_coords <= 0 {
			x_coords += 15
		}
		if x_coords >= screenWidth {
			x_coords -= 15
		}
		if Health <= 0 {
			Enemies = nil
			grounds = nil
			bullets = nil
			rl.UnloadTexture(texture)
			rl.DrawText("Your final score is: "+strconv.Itoa(Score), 30, 40, 30, rl.Red)
		}

		// Met fin au rendu de la frame
		rl.EndDrawing()
	}

	// Ferme la fenêtre de jeu
	rl.CloseWindow()
}
