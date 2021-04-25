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

const (
	grassTile1 int = iota
	grassTile2
	skyChainTile
	chainHolderTile
	pillarDownLeftTile
	pillarUpLeftTile
	pillarDownRightTile
	pillarUpRightTile
	pillarMiddleTile
	skyTile
	cloudTile1
	cloudTile2
	cloudTile3
	cloudTile4
	cloudTile5
	dummyTile
)

func (g *game) updateIntro() {

	g.animationFrame++

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || g.animationFrame > totalAnimationFrames {
		g.animationStep++
		g.animationFrame = 0
		return
	}

}

func (g *game) drawIntro(screen *ebiten.Image) {

	switch g.animationStep {
	case 1:
		drawGeneralIntroScene(screen)
		fadeIn(screen, g.animationFrame, totalAnimationFrames)
	case 2:
		drawGeneralIntroScene(screen)
	case 3:
		drawGeneralIntroScene(screen)
		drawIntroStep1(screen)
	case 4:
		drawGeneralIntroScene(screen)
		drawIntroStep2(screen)
	case 5:
		drawGeneralIntroScene(screen)
	case 6:
		drawGeneralIntroScene(screen)
		fadeOut(screen, g.animationFrame, totalAnimationFrames)
	}

}

func drawIntroStep1(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(4*cellSize), float64(6*cellSize))
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(0, 6*cellSize, cellSize, 7*cellSize)).(*ebiten.Image), &options)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(cellSize, 6*cellSize, 2*cellSize, 7*cellSize)).(*ebiten.Image), &options)
}

func drawIntroStep2(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(6*cellSize), float64(6*cellSize))
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(0, 6*cellSize, cellSize, 7*cellSize)).(*ebiten.Image), &options)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(2*cellSize, 6*cellSize, 3*cellSize, 7*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(0, 6*cellSize, cellSize, 7*cellSize)).(*ebiten.Image), &options)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(2*cellSize, 6*cellSize, 3*cellSize, 7*cellSize)).(*ebiten.Image), &options)
}

var layer1 [12][12]int = [12][12]int{
	[12]int{skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile},
	[12]int{skyTile, cloudTile4, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile},
	[12]int{skyTile, skyTile, skyTile, skyTile, cloudTile1, skyTile, skyTile, cloudTile2, skyTile, skyTile, skyTile, skyTile},
	[12]int{skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile},
	[12]int{skyTile, skyTile, cloudTile3, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile},
	[12]int{skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, cloudTile5, skyTile, skyTile, skyTile},
	[12]int{skyTile, pillarUpLeftTile, pillarMiddleTile, chainHolderTile, pillarMiddleTile, pillarMiddleTile, pillarMiddleTile, pillarMiddleTile, chainHolderTile, pillarMiddleTile, pillarUpRightTile, skyTile},
	[12]int{grassTile1, pillarDownLeftTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, skyTile, pillarDownRightTile, grassTile2},
	[12]int{dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile},
	[12]int{dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile},
	[12]int{dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile},
	[12]int{dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile, dummyTile},
}

var layer2 [12][12]int = [12][12]int{
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile, mainBackTile},
	[12]int{sideTile1, leftBorderTile1, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, rightBorderTile1, sideTile1},
	[12]int{sideTile1, leftBorderTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, rightBorderTile1, sideTile1},
	[12]int{sideTile1, leftBorderTile1, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, rightBorderTile1, sideTile1},
	[12]int{sideTile1, leftBorderTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, backBackTile2, backBackTile1, rightBorderTile1, sideTile1},
}

var layer3 [12][12]int = [12][12]int{
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, chainTile, emptyTile, emptyTile, emptyTile, emptyTile, chainTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, elevatorLeftTile, elevatorChainTile, elevatorTile, elevatorTile, elevatorTile, elevatorTile, elevatorChainTile, elevatorRightTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
	[12]int{emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile, emptyTile},
}

func drawGeneralIntroScene(screen *ebiten.Image) {

	for line := 0; line < 12; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize))
		for row := 0; row < 12; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(layer1[line][row]*cellSize, 5*cellSize, layer1[line][row]*cellSize+cellSize, 6*cellSize)).(*ebiten.Image), &options)
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(layer2[line][row]*cellSize, 0, layer2[line][row]*cellSize+cellSize, cellSize)).(*ebiten.Image), &options)
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(layer3[line][row]*cellSize, 2*cellSize, layer3[line][row]*cellSize+cellSize, 3*cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(4*cellSize), float64(7*cellSize))
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(14*cellSize, 3*cellSize, 15*cellSize, 4*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(2*cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(13*cellSize, 3*cellSize, 14*cellSize, 4*cellSize)).(*ebiten.Image), &options)
	options.GeoM.Translate(float64(cellSize), 0)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(15*cellSize, 3*cellSize, 16*cellSize, 4*cellSize)).(*ebiten.Image), &options)
}
