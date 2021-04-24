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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type fallingObject struct {
	xposition, yposition int
	alive                bool
}

func (fO *fallingObject) update() {
	if fO.alive {
		fO.yposition++
		fO.alive = fO.yposition < elevatorLevel
	}
}

func (fO *fallingObject) draw(screen *ebiten.Image) {
	if fO.alive {
		ebitenutil.DrawRect(screen, float64(fO.xposition*cellSize), float64(fO.yposition*cellSize), float64(cellSize), float64(cellSize), color.RGBA{255, 0, 0, 255})
	}
}

func newFallingObject(xposition int) fallingObject {
	return fallingObject{
		xposition: xposition,
		yposition: 0,
		alive:     true,
	}
}

func (fO *fallingObject) reset(xposition int) {
	fO.alive = true
	fO.xposition = xposition
	fO.yposition = 0
}
