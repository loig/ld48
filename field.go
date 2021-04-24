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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type field struct {
	background [gridHeight][gridWidth + leftMargin + rightMargin]int
	elevator   [gridHeight][gridWidth + leftMargin + rightMargin]int
}

const (
	sideTile int = iota
	leftBorderTile1
	leftBorderTile2
	rightBorderTile1
	rightBorderTile2
	mainBackTile
	otherBackTile
)

const (
	emptyTile int = iota
	elevatorLeftTile
	elevatorTile
	elevatorRightTile
	elevatorChainTile
	chainTile
)

func initField() field {
	f := field{}
	for line := 0; line < gridHeight; line++ {
		f.genLine(line)
		f.genElevator(line)
	}
	return f
}

func (f *field) genLine(line int) {
	for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
		if row < leftMargin-1 {
			f.background[line][row] = sideTile
		} else if row < leftMargin {
			if rand.Intn(2) == 0 {
				f.background[line][row] = leftBorderTile1
			} else {
				f.background[line][row] = leftBorderTile2
			}
		} else if row < gridWidth+leftMargin {
			if rand.Intn(15) == 0 {
				f.background[line][row] = otherBackTile
			} else {
				f.background[line][row] = mainBackTile
			}
		} else if row == gridWidth+leftMargin {
			if rand.Intn(2) == 0 {
				f.background[line][row] = rightBorderTile1
			} else {
				f.background[line][row] = rightBorderTile2
			}
		} else {
			f.background[line][row] = sideTile
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
	for line := 0; line < gridHeight-1; line++ {
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			f.background[line][row] = f.background[line+1][row]
		}
	}
	f.genLine(gridHeight - 1)
}

func (f *field) drawBackground(screen *ebiten.Image) {
	for line := 0; line < gridHeight; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize))
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.background[line][row]*cellSize, 0, f.background[line][row]*cellSize+cellSize, cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}

func (f *field) drawForeground(screen *ebiten.Image) {
	for line := 0; line < gridHeight; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize))
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.elevator[line][row]*cellSize, 2*cellSize, f.elevator[line][row]*cellSize+cellSize, 3*cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}
