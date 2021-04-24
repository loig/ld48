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

type fallingObject struct {
	xposition, yposition int
	alive                bool
	objectType           int
}

const (
	stone1 int = iota
	stone2
)

func (fO *fallingObject) update() {
	if fO.alive {
		fO.yposition++
		fO.alive = fO.yposition < elevatorLevel
	}
}

func (fO *fallingObject) draw(screen *ebiten.Image) {
	if fO.alive {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64((fO.xposition+leftMargin)*cellSize), float64(fO.yposition*cellSize))
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(fO.objectType*cellSize, cellSize, fO.objectType*cellSize+cellSize, 2*cellSize)).(*ebiten.Image), &options)
	}
}

func newFallingObject(xposition int) (fO fallingObject) {
	fO.reset(xposition)
	return fO
}

func (fO *fallingObject) reset(xposition int) {
	fO.alive = true
	fO.xposition = xposition
	fO.yposition = 0
	if rand.Intn(2) == 0 {
		fO.objectType = stone1
	} else {
		fO.objectType = stone2
	}
}
