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

var fallingLevel1 [gridHeight][gridWidth + leftMargin + rightMargin]int = [12][12]int{
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, wallTile, wallTile, diamondTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, wallTile, wallTile, noWallTile, noWallTile, wallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
}

var fallingLevel2 [gridHeight][gridWidth + leftMargin + rightMargin]int = [12][12]int{
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, wallTile, wallTile, noWallTile, noWallTile, wallTile, noWallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile, wallTile, wallTile},
	[12]int{wallTile, wallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, noWallTile, wallTile, wallTile},
}
