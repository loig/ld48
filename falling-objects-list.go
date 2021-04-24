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

type fallingObjectsList struct {
	objects      []fallingObject
	frames       []int
	nextObjectID int
	frameCount   int
}

func (fOL *fallingObjectsList) update() {
	if fOL.nextObjectID < len(fOL.frames) {
		fOL.frameCount++
		for fOL.nextObjectID < len(fOL.frames) &&
			fOL.frameCount >= fOL.frames[fOL.nextObjectID] {
			fOL.nextObjectID++
			fOL.frameCount = 0
		}
	}
	for objectID := 0; objectID < fOL.nextObjectID; objectID++ {
		fOL.objects[objectID].update()
	}
}

func (fOL *fallingObjectsList) draw(screen *ebiten.Image) {
	for objectID := 0; objectID < fOL.nextObjectID; objectID++ {
		fOL.objects[objectID].draw(screen)
	}
}

func initFallingObjectsList() fallingObjectsList {
	fOL := fallingObjectsList{}
	fOL.objects = []fallingObject{
		newFallingObject(100, 16, 16),
		newFallingObject(500, 32, 32),
		newFallingObject(300, 12, 12),
	}
	fOL.frames = []int{
		20, 100, 0,
	}
	return fOL
}
