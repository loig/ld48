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

func (g *game) Update() error {

	switch g.state {
	case stateIntro:
		g.updateIntro()
	case stateElevatorDanger:
		g.f.update()
		g.p.update()
		g.fOL.update(g.sH.isNextFallingObjectStep())
		if g.fOL.doneFalling() {
			g.state = stateElevatorDone
		}
		if g.fallingObjectsCollision() {
			g.state = stateElevatorDead
		}
		if g.fallingDiamondCollision() {
			g.score++
		}
	case stateElevatorDone:
		g.state = stateFallDanger
		g.p.startFall()
		g.f.setFallingLevel()
	case stateElevatorDead:
	case stateFallDanger:
		g.p.update()
		if g.fallingPlayerCollision() {
			g.state = stateFallDead
		}
		if g.fallingPlayerDiamondCollision() {
			g.score++
		}
		if g.p.fallingDone() {
			g.f.isTransition = true
			g.p.yspeed = transitionSpeed
			g.f.transitionSpeed = transitionSpeed
			g.f.backgroundYSpeed = transitionSpeed
			if g.f.setFallingLevel() {
				g.state = stateFallDone
			} else {
				g.state = stateFallTransition
			}
		}
	case stateFallTransition:
		g.p.updateYPosition()
		g.f.updateYPosition()
		g.f.update()
		if g.p.transitionDone() {
			g.p.startFall()
			g.f.isTransition = false
			g.f.backgroundYShift = 0
			g.state = stateFallDanger
		}
	case stateFallDead:
	case stateFallDone:
		g.p.updateYPosition()
		g.f.updateYPosition()
		g.f.update()
		g.endYPosition += transitionSpeed
		if g.p.transitionDone() {
			g.p.startFall()
			g.f.isTransition = false
			g.f.backgroundYShift = 0
			g.endYPosition = 0
			g.state = stateEndGame
		}
	case stateEndGame:
		g.updateEnd()
		g.p.updateYPosition()
	case stateEndBis:
		g.updateEnd()
	}

	return nil
}
