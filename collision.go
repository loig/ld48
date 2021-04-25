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
	"math"
)

func (g *game) fallingObjectsCollision() bool {
	for objectID := 0; objectID < len(g.fOL.objects); objectID++ {
		if g.fOL.objects[objectID].alive && g.fOL.objects[objectID].objectType != diamond && g.p.collide(
			g.fOL.objects[objectID].xposition,
			g.fOL.objects[objectID].yposition,
		) {
			return true
		}
	}
	return false
}

func (g *game) fallingDiamondCollision() bool {
	for objectID := 0; objectID < len(g.fOL.objects); objectID++ {
		if g.fOL.objects[objectID].alive && g.fOL.objects[objectID].objectType == diamond && g.p.collide(
			g.fOL.objects[objectID].xposition,
			g.fOL.objects[objectID].yposition,
		) {
			g.fOL.objects[objectID].alive = false
			return true
		}
	}
	return false
}

func (p *player) collide(x int, y float64) bool {
	return p.xposition == x && p.yposition-float64(cellSize) < y
}

func (g *game) fallingPlayerCollision() bool {
	ymin := int(math.Floor((g.p.yposition + 16) / float64(cellSize)))
	ymax := int(math.Ceil((g.p.yposition - 16) / float64(cellSize)))
	return (ymin > 0 && ymin < gridHeight && g.f.walls[ymin][g.p.xposition+leftMargin] == wallTile) ||
		(ymax > 0 && ymax < gridHeight && g.f.walls[ymax][g.p.xposition+leftMargin] == wallTile)
}

func (g *game) fallingPlayerDiamondCollision() bool {
	ymin := int(math.Floor(g.p.yposition / float64(cellSize)))
	ymax := int(math.Ceil(g.p.yposition / float64(cellSize)))
	if ymin > 0 && ymin < gridHeight && g.f.walls[ymin][g.p.xposition+leftMargin] == diamondTile {
		g.f.walls[ymin][g.p.xposition+leftMargin] = noWallTile
		return true
	}
	if ymax > 0 && ymax < gridHeight && g.f.walls[ymax][g.p.xposition+leftMargin] == diamondTile {
		g.f.walls[ymax][g.p.xposition+leftMargin] = noWallTile
		return true
	}
	return false
}
