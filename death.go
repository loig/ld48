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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) deathUpdate() {

	if g.animationStep < 3 {
		g.animationFrame++
	}

	if g.animationStep == 0 && g.animationFrame > 20 {
		g.animationStep++
		g.animationFrame = 0
	} else if g.animationFrame > totalAnimationFrames {
		g.animationStep++
		g.animationFrame = 0
	}

	if g.animationStep > 0 && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.animationStep = 0
		g.animationFrame = 0
		g.resetGame()
		g.state = stateElevatorDanger
	}

}

func (g *game) drawDeath(screen *ebiten.Image) {
	switch g.animationStep {
	case 0:
		g.drawBlood(screen)
	case 1:
	case 2:
		fadeOut(screen, g.animationFrame, totalAnimationFrames)
		g.drawEnter(screen)
	default:
		screen.Fill(color.Black)
		g.drawEnter(screen)
	}
}

func (g *game) drawBlood(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.p.xposition*cellSize+2*cellSize), g.p.yposition)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(11*cellSize, 4*cellSize, 12*cellSize, 5*cellSize)).(*ebiten.Image), &options)
}

func (g *game) drawEnter(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(5*cellSize+cellSize/2), float64(6*cellSize))
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(15*cellSize, 4*cellSize, 16*cellSize, 5*cellSize)).(*ebiten.Image), &options)
}
