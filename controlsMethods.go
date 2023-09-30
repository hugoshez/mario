package main

import (
	//"fmt"
	//"strconv"
	"github.com/gen2brain/raylib-go/raylib"
)

func (c *Controls) Init() {
	c.KeyJump = rl.KeySpace
	c.KeyLeft = rl.KeyLeft
	c.KeyRight = rl.KeyRight
	c.KeyShoot = rl.KeyF
	c.KeyMenu = rl.KeyG
}

func (c *Controls) setKeyJump(touch int32) {
	c.KeyJump = touch
}

func (c *Controls) setKeyLeft(touch int32) {
	c.KeyLeft = touch
}
func (c *Controls) setKeyRight(touch int32) {
	c.KeyRight = touch
}
func (c *Controls) setKeyShoot(touch int32) {
	c.KeyShoot = touch
}
func (c *Controls) setKeyMenu(touch int32) {
	c.KeyMenu = touch
}
