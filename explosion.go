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

type explosion struct {
	visible              bool
	frame                int
	explosionType        int
	xposition, yposition float64
}

const (
	explosionDiamond int = iota
	explosionRock1
	explosionRock2
)

func (e *explosion) update() {
	if e.visible {
		e.frame++
		if e.frame > 10 {
			e.visible = false
		}
	}
}

func (e *explosion) draw(screen *ebiten.Image) {
	if e.visible {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(e.xposition+float64(2*cellSize), e.yposition)
		screen.DrawImage(spriteSheetImage.SubImage(image.Rect((8+e.explosionType)*cellSize, 4*cellSize, (9+e.explosionType)*cellSize, 5*cellSize)).(*ebiten.Image), &options)
	}
}

func (e *explosion) reset(objectType int, x, y float64) {
	e.visible = true
	e.frame = 0
	e.xposition = x
	e.yposition = y
	if objectType == diamond {
		e.explosionType = explosionDiamond
	} else {
		if rand.Intn(2) == 0 {
			e.explosionType = explosionRock1
		} else {
			e.explosionType = explosionRock2
		}
	}
}
