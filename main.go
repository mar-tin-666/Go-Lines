package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &Game{}
	game.Reset(true)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(gameTitle)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
