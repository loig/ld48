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
		if g.earthShakingFrame == 0 {
			g.playSound(earthquakeSound, true)
		}
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
		g.updateMusic(2)
		g.updateTitleScreen()
	case stateIntro:
		g.stopMusic(2)
		g.updateIntro()
	case stateElevatorDanger:
		g.updateMusic(1)
		g.f.update()
		_, move := g.p.update()
		if move {
			g.playSound(playerMoveSound, true)
		}
		if g.fOL.update(g.sH.isNextFallingObjectStep(), &g.earthShaking) {
			g.playSound(rockSound, false)
		}
		if g.fOL.doneFalling() {
			g.state = stateElevatorDone
			g.animationStep = 0
			g.animationFrame = 0
		}
		if g.fallingObjectsCollision() {
			g.state = stateElevatorDead
			g.playSound(deathSound, false)
		}
		if g.fallingDiamondCollision() {
			g.score++
			g.playSound(diamondCatchSound, true)
		}
	case stateElevatorDone:
		g.stopMusic(1)
		if g.updateElevatorBreak() {
			g.state = statePrepareFall
			g.p.startFall()
			g.f.setFallingLevel()
			g.animationStep = 0
			g.animationFrame = 0
		}
	case statePrepareFall:
		g.updateMusic(2)
		if g.updateAfterElevator() {
			g.state = stateFallDanger
		}
	case stateElevatorDead:
		g.stopMusic(1)
		g.f.update()
		g.fOL.update(false, &g.earthShaking)
		g.deathUpdate()
	case stateFallDanger:
		g.updateMusic(2)
		dead, move := g.p.update()
		if move {
			g.playSound(playerMoveSound, true)
		}
		if dead || g.fallingPlayerCollision() {
			g.state = stateFallDead
			g.playSound(deathSound, false)
		}
		if g.fallingPlayerDiamondCollision() {
			g.score++
			g.playSound(diamondCatchSound, false)
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
		g.updateMusic(2)
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
		g.stopMusic(2)
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
			g.stopMusic(2)
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
