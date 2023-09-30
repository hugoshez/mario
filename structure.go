package main

import (
	//"fmt"
	//"strconv"

	"github.com/gen2brain/raylib-go/raylib"
)

type Ground struct {
	posX   int32
	posY   int32
	width  int32
	height int32
	Color  rl.Color
}

type Bullet struct {
	posX            int32
	posY            int32
	radius          float32
	Right_direction bool
	Draw            bool
	Color           rl.Color
}

type Enemy struct {
	posX      int32
	posY      int32
	velocity  int32
	Damage    float32
	Direction bool
	Draw      bool
	Color     rl.Color
}

type Controls struct {
	KeyJump  int32
	KeyLeft  int32
	KeyRight int32
	KeyShoot int32
	KeyMenu  int32
}
