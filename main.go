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
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := initGame()

	ebiten.SetWindowTitle("LD48")
	//ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)

	err := ebiten.RunGame(g)
	if err != nil {
		log.Print(err)
	}
}
