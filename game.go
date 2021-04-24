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

type game struct {
	state int
	p     player
	fOL   fallingObjectsList
	sH    speedHandler
	f     field
}

const (
	stateElevatorDanger int = iota
	stateElevatorDone
	stateElevatorDead
)

func initGame() *game {
	g := game{}

	loadAssets()

	g.p = player{
		xposition: 3,
		yposition: 9,
	}

	g.fOL = initFallingObjectsList(100)

	g.sH = speedHandler{
		framesPerElevatorStep:      20,
		framesPerFallingObjectStep: 15,
	}

	g.f = initField()

	return &g
}
