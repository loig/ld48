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

type speedHandler struct {
	framesPerElevatorStep      int
	elevatorFrame              int
	framesPerFallingObjectStep int
	fallingObjectFrame         int
	framesPerFallingPlayerStep int
	fallingPlayerFrame         int
}

func (s *speedHandler) isNextElevatorStep() bool {
	s.elevatorFrame++
	if s.elevatorFrame >= s.framesPerElevatorStep {
		s.elevatorFrame = 0
		return true
	}
	return false
}

func (s *speedHandler) isNextFallingObjectStep() bool {
	s.fallingObjectFrame++
	if s.fallingObjectFrame >= s.framesPerFallingObjectStep {
		s.fallingObjectFrame = 0
		return true
	}
	return false
}

func (s *speedHandler) isNextFallingPlayerStep() bool {
	s.fallingPlayerFrame++
	if s.fallingPlayerFrame >= s.framesPerFallingPlayerStep {
		s.fallingPlayerFrame = 0
		return true
	}
	return false
}
