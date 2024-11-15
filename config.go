package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 801
	screenHeight = 601
	playerSize   = 2
	moveSpeed    = 1

	gameTitle            = "Go Lines"
	startMessage         = "Press SPACE to start, R to reset"
	controlInstructions  = "Controls: Player 1 (red) - WSAD, Player 2 (green) - Arrows"
	winMessageTemplate   = "%s scored a point!"
	drawMessageTemplate  = "Head-on collision - draw, no points!"
	pressSpaceToContinue = "Press SPACE to continue or R to reset..."
	player1              = "Player 1"
	player2              = "Player 2"
)

var (
	backgroundColor = color.Black
	player1Color    = color.RGBA{255, 0, 0, 255}
	player2Color    = color.RGBA{0, 255, 0, 255}

	player1Keys = struct {
		Up, Down, Left, Right ebiten.Key
	}{Up: ebiten.KeyW, Down: ebiten.KeyS, Left: ebiten.KeyA, Right: ebiten.KeyD}

	player2Keys = struct {
		Up, Down, Left, Right ebiten.Key
	}{Up: ebiten.KeyArrowUp, Down: ebiten.KeyArrowDown, Left: ebiten.KeyArrowLeft, Right: ebiten.KeyArrowRight}
)
