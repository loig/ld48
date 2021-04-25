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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type player struct {
	xposition int
	yposition float64
	yspeed    float64
	isFalling bool
	pose      int
}

const (
	pose1 int = iota
	pose2
	pose3
	pose4
	pose5
	endPose
)

func (p *player) update() {
	p.updateXPosition()
	if p.isFalling {
		p.updateYPosition()
	}
}

func (p *player) updateXPosition() {
	currentPosition := p.xposition
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
	if currentPosition != p.xposition && !p.isFalling {
		currentPose := p.pose
		p.pose = rand.Intn(endPose)
		if currentPose == p.pose {
			p.pose = (p.pose + 1) % endPose
		}
	}
}

func (p *player) updateYPosition() {
	p.yposition += p.yspeed
}

func (p *player) startFall() {
	p.isFalling = true
	p.yposition = 0
	p.yspeed = playerFallSpeed
	p.pose = 12
}

func (p *player) fallingDone() bool {
	return p.yposition >= float64((gridHeight-1)*cellSize)
}

func (p *player) transitionDone() bool {
	return p.yposition <= 0
}

func (p *player) draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64((p.xposition+leftMargin)*cellSize), p.yposition)
	screen.DrawImage(spriteSheetImage.SubImage(image.Rect(p.pose*cellSize, 3*cellSize, p.pose*cellSize+cellSize, 4*cellSize)).(*ebiten.Image), &options)
}
