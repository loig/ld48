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

func (g *game) updateElevatorBreak() bool {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.animationFrame = 0
		g.animationStep = 0
		return true
	}

	g.animationFrame++

	if g.animationStep < 3 && g.animationFrame > 20 {
		g.animationFrame = 0
		g.animationStep++
		if g.animationStep == 1 {
			g.f.elevator[elevatorLevel][5] = emptyTile
			g.f.elevator[elevatorLevel][7] = emptyTile
			g.f.elevator[elevatorLevel][8] = emptyTile
			g.playSound(earthquakeSound, false)
		}
		if g.animationStep == 2 {
			g.f.elevator[elevatorLevel][2] = emptyTile
			g.f.elevator[elevatorLevel][3] = emptyTile
			g.f.elevator[elevatorLevel][4] = emptyTile
			g.f.elevator[elevatorLevel][6] = emptyTile
			g.f.elevator[elevatorLevel][9] = emptyTile
			g.playSound(earthquakeSound, false)
		}
	} else if g.animationFrame > totalAnimationFrames {
		g.animationStep++
		g.animationFrame = 0
		if g.animationStep >= 4 {
			g.animationStep = 0
			return true
		}
	}

	if g.animationStep >= 3 {
		g.p.pose = 12
		g.p.yposition += playerFallSpeed
	}

	return false
}

func (g *game) updateAfterElevator() bool {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.animationFrame = 0
		g.animationStep = 0
		return true
	}

	g.animationFrame++
	if g.animationFrame > totalAnimationFrames {
		g.animationFrame = 0
		return true
	}

	return false
}

func (g *game) drawElevatorBreak(screen *ebiten.Image) {

	g.f.drawElevator(screen)

	if g.animationStep == 1 {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(5*cellSize), float64(elevatorLevel*cellSize))
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(2*cellSize), 0)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(cellSize), 0)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
	}

	if g.animationStep == 2 {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(2*cellSize), float64(elevatorLevel*cellSize))
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(cellSize), 0)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(cellSize), 0)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(2*cellSize), float64(elevatorLevel*cellSize))
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
		options.GeoM.Translate(float64(3*cellSize), float64(elevatorLevel*cellSize))
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(7*cellSize, 4*cellSize, 8*cellSize, 5*cellSize)).(*ebiten.Image), &options)
	}

	if g.animationStep == 3 {
		fadeOut(screen, g.animationFrame, totalAnimationFrames)
	}
}
