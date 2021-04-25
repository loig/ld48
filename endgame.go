// A game for Ludum Dare, 48th edition
//    Copyright (C) 2021 Loïg Jezequel
/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) updateEnd() {

	if g.state == stateEndGame {
		if g.p.yposition >= float64(screenHeight-2*cellSize) {
			g.state = stateEndBis
			g.animationStep = 1
		}
	} else {
		g.animationFrame++

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || (g.animationStep < 3 && (g.animationFrame > totalAnimationFrames ||
			g.animationStep == 1 && g.animationFrame > 10)) {
			g.animationStep++
			g.animationFrame = 0
			if g.animationStep > 3 {
				g.animationStep = 0
				g.resetGame()
				g.state = stateTitle
			}
			return
		}
	}

}

func (g *game) drawEnd(screen *ebiten.Image) {

	g.drawEndImage(screen)

	switch g.animationStep {
	case 1:
		g.drawPlouf(screen)
	case 3:
		g.drawScore(screen)
	}

}

func (g *game) drawEndImage(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, g.endYPosition)
	screen.DrawImage(endImage, &options)
}

func (g *game) drawPlouf(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.p.xposition*cellSize+2*cellSize), g.p.yposition)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(11*cellSize, 3*cellSize, 12*cellSize, 4*cellSize)).(*ebiten.Image), &options)
}

func (g *game) drawScore(screen *ebiten.Image) {
	firstDigit := g.score / 10
	secondDigit := g.score % 10
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(3*cellSize), float64(5*cellSize))
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(10*cellSize, cellSize, 11*cellSize, 2*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect((6+firstDigit)*cellSize, 2*cellSize, (7+firstDigit)*cellSize, 3*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect((6+secondDigit)*cellSize, 2*cellSize, (7+secondDigit)*cellSize, 3*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(12*cellSize, 4*cellSize, 13*cellSize, 5*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 2*cellSize, 8*cellSize, 3*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(6*cellSize, 2*cellSize, 7*cellSize, 3*cellSize)).(*ebiten.Image), &options)
}
