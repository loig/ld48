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
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) updateTitleScreen() {

	if g.animationStep >= 1 {
		g.animationFrame++
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || g.animationFrame > totalAnimationFrames {
		g.animationStep++
		if g.animationStep >= 2 {
			g.animationStep = 0
			g.animationFrame = 0
			g.state = stateIntro
		}
	}

}

func (g *game) drawTitleScreen(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}
	screen.DrawImage(titleImage, &options)

	switch g.animationStep {
	case 1:
		fadeOut(screen, g.animationFrame, totalAnimationFrames)
	}
}
