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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type player struct {
	xposition, yposition int
}

func (p *player) update() {
	p.updateXPosition()
}

func (p *player) updateXPosition() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.xposition++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.xposition--
	}
	if p.xposition < 0 {
		p.xposition = 0
	}
	if p.xposition >= gridWidth {
		p.xposition = gridWidth - 1
	}
}

func (p *player) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, float64((p.xposition+leftMargin)*cellSize), float64(p.yposition*cellSize), float64(cellSize), float64(cellSize), color.White)
}
