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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type field struct {
	background       [gridHeight + 2][gridWidth + leftMargin + rightMargin]int
	backgroundYShift float64
	backgroundYSpeed float64
	elevator         [gridHeight][gridWidth + leftMargin + rightMargin]int
	walls            [gridHeight][gridWidth + leftMargin + rightMargin]int
	fallingLevelNum  int
}

const (
	sideTile1 int = iota
	sideTile2
	leftBorderTile1
	leftBorderTile2
	leftBorderTile3
	rightBorderTile1
	rightBorderTile2
	rightBorderTile3
	mainBackTile
	otherBackTile1
	otherBackTile2
	otherBackTile3
	otherBackTile4
)

const (
	emptyTile int = iota
	elevatorLeftTile
	elevatorTile
	elevatorRightTile
	elevatorChainTile
	chainTile
)

const (
	noWallTile int = iota
	wallTile
)

func initField() field {
	f := field{}
	for line := 0; line < gridHeight; line++ {
		f.genLine(line)
		f.genElevator(line)
	}
	f.genLine(gridHeight)
	f.genLine(gridHeight + 1)
	f.backgroundYSpeed = initialBackgroundYSpeed
	return f
}

func (f *field) genLine(line int) {
	for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
		if row < leftMargin-1 {
			if rand.Intn(10) != 0 {
				f.background[line][row] = sideTile1
			} else {
				f.background[line][row] = sideTile2
			}
		} else if row < leftMargin {
			if rand.Intn(10) != 0 {
				f.background[line][row] = leftBorderTile1
			} else {
				if rand.Intn(2) == 0 {
					f.background[line][row] = leftBorderTile2
				} else {
					f.background[line][row] = leftBorderTile3
				}
			}
		} else if row < gridWidth+leftMargin {
			if rand.Intn(25) == 0 {
				switch rand.Intn(4) {
				case 0:
					f.background[line][row] = otherBackTile1
				case 1:
					f.background[line][row] = otherBackTile2
				case 2:
					f.background[line][row] = otherBackTile3
				case 3:
					f.background[line][row] = otherBackTile4
				}
			} else {
				f.background[line][row] = mainBackTile
			}
		} else if row == gridWidth+leftMargin {
			if rand.Intn(10) != 0 {
				f.background[line][row] = rightBorderTile1
			} else {
				if rand.Intn(2) == 0 {
					f.background[line][row] = rightBorderTile2
				} else {
					f.background[line][row] = rightBorderTile3
				}
			}
		} else {
			if rand.Intn(10) != 0 {
				f.background[line][row] = sideTile1
			} else {
				f.background[line][row] = sideTile2
			}
		}
	}
}

func (f *field) genElevator(line int) {
	if line == elevatorLevel {
		f.elevator[line][leftMargin] = elevatorLeftTile
		f.elevator[line][leftMargin+1] = elevatorChainTile
		for row := leftMargin + 2; row < leftMargin+gridWidth-2; row++ {
			f.elevator[line][row] = elevatorTile
		}
		f.elevator[line][leftMargin+gridWidth-2] = elevatorChainTile
		f.elevator[line][leftMargin+gridWidth-1] = elevatorRightTile
	} else if line < elevatorLevel {
		f.elevator[line][leftMargin+1] = chainTile
		f.elevator[line][leftMargin+gridWidth-2] = chainTile
	}
}

func (f *field) update() {
	f.backgroundYShift += f.backgroundYSpeed
	if f.backgroundYShift >= float64(cellSize) {
		for line := 0; line < gridHeight+1; line++ {
			for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
				f.background[line][row] = f.background[line+1][row]
			}
		}
		f.genLine(gridHeight + 1)
		f.backgroundYShift = 0
	}
}

func (f *field) drawBackground(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 60, 57, 255})
	for line := 0; line < gridHeight+2; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize-cellSize)-f.backgroundYShift)
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.background[line][row]*cellSize, 0, f.background[line][row]*cellSize+cellSize, cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}

func (f *field) drawElevator(screen *ebiten.Image) {
	for line := 0; line < gridHeight; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize))
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.elevator[line][row]*cellSize, 2*cellSize, f.elevator[line][row]*cellSize+cellSize, 3*cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}

func (f *field) setFallingLevel() {
	f.walls = [12][12]int{
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
		[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	}
	f.fallingLevelNum++
}

func (f *field) drawWalls(screen *ebiten.Image) {
	for line := 0; line < gridHeight; line++ {
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			if f.walls[line][row] != noWallTile {
				ebitenutil.DrawRect(screen, float64(row*cellSize), float64(line*cellSize), float64(cellSize), float64(cellSize), color.RGBA{0, 0, 255, 255})
			}
		}
	}
}
