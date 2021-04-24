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
	case stateElevatorDanger:
		g.f.update()
		g.p.update(false)
		g.fOL.update(g.sH.isNextFallingObjectStep())
		if g.fOL.doneFalling() {
			g.state = stateElevatorDone
		}
		if g.fallingObjectsCollision() {
			g.state = stateElevatorDead
		}
	case stateElevatorDone:
		g.state = stateFallDanger
		g.p.startFall()
		g.f.setFallingLevel()
	case stateElevatorDead:
	case stateFallDanger:
		g.p.update(g.sH.isNextFallingPlayerStep())
		if g.fallingPlayerCollision() {
			g.state = stateFallDead
		}
		if g.p.fallingDone() {
			g.p.startFall()
			g.f.setFallingLevel()
		}
	case stateFallDead:
	}

	return nil
}
