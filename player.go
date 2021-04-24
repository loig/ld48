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

type player struct {
	xposition, yposition float64
	xspeed               float64
	width, height        float64
	falling              bool
}

func (p *player) update() {
	p.updateSpeed()
	p.updatePosition()
}

func (p *player) updateSpeed() {

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		if p.falling {
			if p.xspeed < playerMaxSpeed {
				p.xspeed += playerSpeedIncrement
			}
			return
		}
		p.xspeed = playerMaxSpeed
		return
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		if p.falling {
			if p.xspeed > -playerMaxSpeed {
				p.xspeed -= playerSpeedIncrement
			}
			return
		}
		p.xspeed = -playerMaxSpeed
		return
	}

	if p.falling {
		if p.xspeed > 0 {
			p.xspeed -= playerSpeedIncrement
			if p.xspeed < 0 {
				p.xspeed = 0
			}
			return
		}
		if p.xspeed < 0 {
			p.xspeed += playerSpeedIncrement
			if p.xspeed > 0 {
				p.xspeed = 0
			}
			return
		}
	}

	p.xspeed = 0

}

func (p *player) updatePosition() {
	p.xposition += p.xspeed
}

func (p *player) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.xposition-p.width/2, p.yposition-p.height/2, p.width, p.height, color.White)
}
