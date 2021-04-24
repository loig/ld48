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
	xposition, yposition float64
	yspeed               float64
	width, height        float64
	out                  bool
}

func (fO *fallingObject) update() {
	if !fO.out {
		fO.yposition += fO.yspeed
		fO.out = fO.yposition-fO.width/2 > float64(screenHeight)
	}
}

func (fO *fallingObject) draw(screen *ebiten.Image) {
	if !fO.out {
		ebitenutil.DrawRect(screen, fO.xposition-fO.width/2, fO.yposition-fO.height/2, fO.width, fO.height, color.RGBA{255, 0, 0, 255})
	}
}

func newFallingObject(xposition, width, height float64) fallingObject {
	return fallingObject{
		xposition: xposition,
		yposition: -width / 2,
		yspeed:    fallingObjectSpeed,
		width:     width,
		height:    height,
	}
}
