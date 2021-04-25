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

import "github.com/hajimehoshi/ebiten/v2"

func (g *game) Draw(screen *ebiten.Image) {

	switch g.state {
	case stateTitle:
		g.drawTitleScreen(screen)
	case stateIntro:
		g.drawIntro(screen)
	case stateElevatorDanger:
		g.f.drawBackground(screen, true, g.earthShakingXShift, g.earthShakingYShift)
		g.p.draw(screen)
		g.fOL.draw(screen)
		g.f.drawElevator(screen)
	case stateElevatorDead:
		g.f.drawBackground(screen, true, g.earthShakingXShift, g.earthShakingYShift)
		g.fOL.draw(screen)
		g.f.drawElevator(screen)
		g.drawDeath(screen)
	case stateFallDanger, stateFallTransition:
		g.f.drawBackground(screen, false, g.earthShakingXShift, g.earthShakingYShift)
		g.p.draw(screen)
		g.f.drawWalls(screen)
	case stateFallDead:
		g.f.drawBackground(screen, false, g.earthShakingXShift, g.earthShakingYShift)
		g.f.drawWalls(screen)
		g.drawDeath(screen)
	case stateFallDone:
		g.f.drawBackground(screen, false, g.earthShakingXShift, g.earthShakingYShift)
		g.drawEnd(screen)
		g.p.draw(screen)
		g.f.drawWalls(screen)
	case stateEndGame:
		g.f.drawBackground(screen, false, g.earthShakingXShift, g.earthShakingYShift)
		g.drawEnd(screen)
		g.p.draw(screen)
	case stateEndBis:
		g.f.drawBackground(screen, false, g.earthShakingXShift, g.earthShakingYShift)
		g.drawEnd(screen)
	}

}
