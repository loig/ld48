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
	xposition  int
	yposition  float64
	yspeed     float64
	alive      bool
	objectType int
}

const (
	stone1 int = iota
	stone2
	stone3
	stone4
	stone5
	stone6
	stone7
	stone8
	stone9
	stone10
	diamond
)

func (fO *fallingObject) update() {
	if fO.alive {
		fO.yposition += fO.yspeed
		fO.alive = fO.yposition < float64(elevatorLevel*cellSize-cellSize)
	}
}

func (fO *fallingObject) draw(screen *ebiten.Image) {
	if fO.alive {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64((fO.xposition+leftMargin)*cellSize), fO.yposition)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect(fO.objectType*cellSize, cellSize, fO.objectType*cellSize+cellSize, 2*cellSize)).(*ebiten.Image), &options)
	}
}

func newFallingObject(xposition int, yspeed float64, isDiamond bool) (fO fallingObject) {
	fO.reset(xposition, yspeed, isDiamond)
	return fO
}

func (fO *fallingObject) reset(xposition int, yspeed float64, isDiamond bool) {
	fO.alive = true
	fO.xposition = xposition
	fO.yposition = 0
	fO.yspeed = yspeed
	if isDiamond {
		fO.objectType = diamond
	} else {
		switch rand.Intn(10) {
		case 0:
			fO.objectType = stone1
		case 1:
			fO.objectType = stone2
		case 2:
			fO.objectType = stone3
		case 3:
			fO.objectType = stone4
		case 4:
			fO.objectType = stone5
		case 5:
			fO.objectType = stone6
		case 6:
			fO.objectType = stone7
		case 7:
			fO.objectType = stone8
		case 8:
			fO.objectType = stone9
		case 9:
			fO.objectType = stone10
		}
	}
}
