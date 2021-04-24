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

func (g *game) fallingObjectsCollision() bool {
	for objectID := 0; objectID < g.fOL.nextObjectID; objectID++ {
		if g.p.collide(
			g.fOL.objects[objectID].xposition,
			g.fOL.objects[objectID].yposition,
			g.fOL.objects[objectID].width,
			g.fOL.objects[objectID].height,
		) {
			return true
		}
	}
	return false
}

func (p *player) collide(x, y, w, h float64) bool {
	pxmax := p.xposition + p.width/2
	pxmin := p.xposition - p.width/2
	xmax := x + w/2
	xmin := x - w/2
	xcollide := pxmin < xmax && pxmax > xmin
	if !xcollide {
		return false
	}
	pymax := p.yposition + p.height/2
	pymin := p.yposition - p.height/2
	ymax := y + h/2
	ymin := y - h/2
	ycollide := pymin < ymax && pymax > ymin
	return ycollide
}
