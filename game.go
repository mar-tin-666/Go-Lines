package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player1Pos     [2]int
	player2Pos     [2]int
	player1Trail   [][2]int
	player2Trail   [][2]int
	player1Dir     [2]int
	player2Dir     [2]int
	score1, score2 int
	noScore        bool
	running        bool
	winner         string
	frameCount     int // Licznik klatek do obsługi szybkości ruchu
	startScreen    bool
}

func (g *Game) Reset(startScreen bool) {
	g.player1Pos = [2]int{screenWidth / 4, screenHeight / 2}
	g.player2Pos = [2]int{3 * screenWidth / 4, screenHeight / 2}
	g.player1Trail = nil
	g.player2Trail = nil
	g.player2Dir = [2]int{-1, 0}
	g.player1Dir = [2]int{1, 0}
	g.running = false
	g.noScore = false
	g.startScreen = startScreen
	g.winner = ""
	g.frameCount = 0

}
func (g *Game) Update() error {
	if g.startScreen {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.startScreen = false
			g.running = true
		}
		return nil
	}

	if !g.running {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Reset(false)
			g.running = true
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Reset(true)
		}
		return nil
	}

	// Game reset
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Reset(true)
	}

	// Player 1 controls
	if inpututil.IsKeyJustPressed(player1Keys.Up) && g.player1Dir != [2]int{0, 1} {
		g.player1Dir = [2]int{0, -1}
	}
	if inpututil.IsKeyJustPressed(player1Keys.Down) && g.player1Dir != [2]int{0, -1} {
		g.player1Dir = [2]int{0, 1}
	}
	if inpututil.IsKeyJustPressed(player1Keys.Left) && g.player1Dir != [2]int{1, 0} {
		g.player1Dir = [2]int{-1, 0}
	}
	if inpututil.IsKeyJustPressed(player1Keys.Right) && g.player1Dir != [2]int{-1, 0} {
		g.player1Dir = [2]int{1, 0}
	}

	// Player 2 controls
	if inpututil.IsKeyJustPressed(player2Keys.Up) && g.player2Dir != [2]int{0, 1} {
		g.player2Dir = [2]int{0, -1}
	}
	if inpututil.IsKeyJustPressed(player2Keys.Down) && g.player2Dir != [2]int{0, -1} {
		g.player2Dir = [2]int{0, 1}
	}
	if inpututil.IsKeyJustPressed(player2Keys.Left) && g.player2Dir != [2]int{1, 0} {
		g.player2Dir = [2]int{-1, 0}
	}
	if inpututil.IsKeyJustPressed(player2Keys.Right) && g.player2Dir != [2]int{-1, 0} {
		g.player2Dir = [2]int{1, 0}
	}

	// Position update
	if g.frameCount%moveSpeed == 0 {
		g.player1Pos[0] += g.player1Dir[0] * playerSize
		g.player1Pos[1] += g.player1Dir[1] * playerSize
		g.player2Pos[0] += g.player2Dir[0] * playerSize
		g.player2Pos[1] += g.player2Dir[1] * playerSize

		// Add trail
		g.player1Trail = append(g.player1Trail, g.player1Pos)
		g.player2Trail = append(g.player2Trail, g.player2Pos)

		// Check collision
		if g.player1Pos == g.player2Pos {
			g.noScore = true
			g.winner = fmt.Sprint(drawMessageTemplate)
			g.running = false
		} else if g.checkCollision(g.player1Pos, g.player1Trail[:len(g.player1Trail)-1]) || g.checkCollision(g.player1Pos, g.player2Trail) || g.isOutOfBounds(g.player1Pos) {
			g.score2++
			g.winner = fmt.Sprintf(winMessageTemplate, player2)
			g.running = false
		} else if g.checkCollision(g.player2Pos, g.player2Trail[:len(g.player2Trail)-1]) || g.checkCollision(g.player2Pos, g.player1Trail) || g.isOutOfBounds(g.player2Pos) {
			g.score1++
			g.winner = fmt.Sprintf(winMessageTemplate, player1)
			g.running = false
		}

	}

	g.frameCount++
	return nil
}

func (g *Game) checkCollision(pos [2]int, trail [][2]int) bool {
	for _, t := range trail {
		if pos == t {
			return true
		}
	}
	return false
}

func (g *Game) isOutOfBounds(pos [2]int) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= screenWidth || pos[1] >= screenHeight
}
