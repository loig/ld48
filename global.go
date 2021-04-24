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

const (
	cellSize                   int     = 64
	gridWidth                  int     = 8
	gridHeight                 int     = 12
	leftMargin                 int     = 2
	rightMargin                int     = 2
	elevatorLevel              int     = 10
	elevatorNumLevelsPhase1    int     = 3
	elevatorNumLevelsPhase2    int     = 3
	elevatorNumObjectsPerLevel int     = 15
	initialSpawnChances        int     = 3
	initialSpawnInterval       int     = 3
	initialBackgroundYSpeed    float64 = 1
	initialObjectSpeed         float64 = 7
	screenWidth                        = (gridWidth + leftMargin + rightMargin) * cellSize
	screenHeight                       = gridHeight * cellSize
)
