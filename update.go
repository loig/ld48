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

import "math/rand"

func (g *game) Update() error {

	if g.earthShaking {
		g.earthShakingFrame++
		if g.earthShakingFrame%5 == 0 {
			g.earthShakingXShift = float64(rand.Intn(10)) - 5
			g.earthShakingYShift = float64(rand.Intn(10)) - 5
		}
		if g.earthShakingFrame > 60 {
			g.earthShakingXShift = 0
			g.earthShakingYShift = 0
			g.earthShaking = false
			g.earthShakingFrame = 0
		}
	}

	switch g.state {
	case stateTitle:
		g.updateTitleScreen()
	case stateIntro:
		g.updateIntro()
	case stateElevatorDanger:
		g.f.update()
		g.p.update()
		g.fOL.update(g.sH.isNextFallingObjectStep(), &g.earthShaking)
		if g.fOL.doneFalling() {
			g.state = stateElevatorDone
			g.animationStep = 0
			g.animationFrame = 0
		}
		if g.fallingObjectsCollision() {
			g.state = stateElevatorDead
		}
		if g.fallingDiamondCollision() {
			g.score++
		}
	case stateElevatorDone:
		if g.updateElevatorBreak() {
			g.state = statePrepareFall
			g.p.startFall()
			g.f.setFallingLevel()
			g.animationStep = 0
			g.animationFrame = 0
		}
	case statePrepareFall:
		if g.updateAfterElevator() {
			g.state = stateFallDanger
		}
	case stateElevatorDead:
		g.f.update()
		g.fOL.update(false, &g.earthShaking)
		g.deathUpdate()
	case stateFallDanger:
		dead := g.p.update()
		if dead || g.fallingPlayerCollision() {
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
		g.deathUpdate()
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
