package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {

	// Draw welcome screen
	if g.startScreen {
		screen.Fill(backgroundColor)
		ebitenutil.DebugPrintAt(screen, gameTitle, screenWidth/2-30, screenHeight/2-40)
		ebitenutil.DebugPrintAt(screen, controlInstructions, screenWidth/2-170, screenHeight/2)
		ebitenutil.DebugPrintAt(screen, startMessage, screenWidth/2-100, screenHeight/2+40)
		return
	}

	// Draw scores
	score := fmt.Sprintf("%d : %d", g.score1, g.score2)
	ebitenutil.DebugPrintAt(screen, score, screenWidth/2-20, 10)

	// Draw lines
	for _, t := range g.player1Trail {
		screen.Set(t[0], t[1], player1Color)
	}
	for _, t := range g.player2Trail {
		screen.Set(t[0], t[1], player2Color)
	}

	// Draw winning message
	if !g.running {
		messageX := screenWidth/2 - 80
		if g.noScore {
			messageX = screenWidth/2 - 120
		}
		ebitenutil.DebugPrintAt(screen, g.winner, messageX, screenHeight/2-20)
		ebitenutil.DebugPrintAt(screen, pressSpaceToContinue, screenWidth/2-125, screenHeight/2+20)
		return
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
