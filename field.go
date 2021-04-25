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
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type field struct {
	background       [gridHeight + 2][gridWidth + leftMargin + rightMargin]int
	backgroundYShift float64
	backgroundYSpeed float64
	elevator         [gridHeight][gridWidth + leftMargin + rightMargin]int
	walls            [gridHeight][gridWidth + leftMargin + rightMargin]int
	oldWalls         [gridHeight][gridWidth + leftMargin + rightMargin]int
	wallTiles        [gridHeight * 2][(gridWidth + leftMargin + rightMargin) * 2]int
	oldWallTiles     [gridHeight * 2][(gridWidth + leftMargin + rightMargin) * 2]int
	fallingLevelNum  int
	currentOddity    int
	isTransition     bool
	transitionSpeed  float64
	yposition        float64
}

const (
	sideTile1 int = iota
	sideTile2
	leftBorderTile1
	leftBorderTile2
	leftBorderTile3
	rightBorderTile1
	rightBorderTile2
	rightBorderTile3
	mainBackTile
	otherBackTile1
	otherBackTile2
	otherBackTile3
	otherBackTile4
	backBackTile1
	backBackTile2
)

const (
	emptyTile int = iota
	elevatorLeftTile
	elevatorTile
	elevatorRightTile
	elevatorChainTile
	chainTile
)

const (
	noWallTile int = iota
	wallTile
	diamondTile
)

const (
	topLeftCornerTile int = iota
	topRightCornerTile
	topLeftSmallCornerTile
	topRightSmallCornerTile
	fullTile1
	fullTile2
	topEdgeTile1
	topEdgeTile2
	leftEdgeTile1
	rightEdgeTile1
	bottomLeftCornerTile
	bottomRightCornerTile
	bottomLeftSmallCornerTile
	bottomRightSmallCornerTile
	fullTile3
	fullTile4
	bottomEdgeTile1
	bottomEdgeTile2
	leftEdgeTile2
	rightEdgeTile2
	emptySmallTile
)

func initField() field {
	f := field{}
	for line := 0; line < gridHeight; line++ {
		f.genLine(line)
		f.genElevator(line)
	}
	f.genLine(gridHeight)
	f.genLine(gridHeight + 1)
	f.backgroundYSpeed = initialBackgroundYSpeed
	return f
}

func (f *field) genLine(line int) {
	for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
		if row < leftMargin-1 {
			if rand.Intn(10) != 0 {
				f.background[line][row] = sideTile1
			} else {
				f.background[line][row] = sideTile2
			}
		} else if row < leftMargin {
			if rand.Intn(10) != 0 {
				f.background[line][row] = leftBorderTile1
			} else {
				if rand.Intn(2) == 0 {
					f.background[line][row] = leftBorderTile2
				} else {
					f.background[line][row] = leftBorderTile3
				}
			}
		} else if row < gridWidth+leftMargin {
			if rand.Intn(25) == 0 {
				switch rand.Intn(4) {
				case 0:
					f.background[line][row] = otherBackTile1
				case 1:
					f.background[line][row] = otherBackTile2
				case 2:
					f.background[line][row] = otherBackTile3
				case 3:
					f.background[line][row] = otherBackTile4
				}
			} else {
				f.background[line][row] = mainBackTile
			}
		} else if row == gridWidth+leftMargin {
			if rand.Intn(10) != 0 {
				f.background[line][row] = rightBorderTile1
			} else {
				if rand.Intn(2) == 0 {
					f.background[line][row] = rightBorderTile2
				} else {
					f.background[line][row] = rightBorderTile3
				}
			}
		} else {
			if rand.Intn(10) != 0 {
				f.background[line][row] = sideTile1
			} else {
				f.background[line][row] = sideTile2
			}
		}
	}
}

func (f *field) genElevator(line int) {
	if line == elevatorLevel {
		f.elevator[line][leftMargin] = elevatorLeftTile
		f.elevator[line][leftMargin+1] = elevatorChainTile
		for row := leftMargin + 2; row < leftMargin+gridWidth-2; row++ {
			f.elevator[line][row] = elevatorTile
		}
		f.elevator[line][leftMargin+gridWidth-2] = elevatorChainTile
		f.elevator[line][leftMargin+gridWidth-1] = elevatorRightTile
	} else if line < elevatorLevel {
		f.elevator[line][leftMargin+1] = chainTile
		f.elevator[line][leftMargin+gridWidth-2] = chainTile
	}
}

func (f *field) update() {
	f.backgroundYShift += f.backgroundYSpeed
	if -f.backgroundYShift >= float64(cellSize) {
		for line := 0; line < gridHeight+1; line++ {
			for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
				f.background[line][row] = f.background[line+1][row]
			}
		}
		f.genLine(gridHeight + 1)
		f.backgroundYShift = f.backgroundYShift + float64(cellSize)
		f.currentOddity = (f.currentOddity + 1) % 2
	}
}

func (f *field) drawBackground(screen *ebiten.Image, withSide bool) {
	for line := 0; line < gridHeight+2; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize-cellSize)+f.backgroundYShift)
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			backTile := backBackTile1
			if (line+row+f.currentOddity)%2 == 0 {
				backTile = backBackTile2
			}
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(backTile*cellSize, 0, backTile*cellSize+cellSize, cellSize)).(*ebiten.Image), &options)
			if withSide || (row != 1 && row != gridWidth+leftMargin+rightMargin-2) {
				screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.background[line][row]*cellSize, 0, f.background[line][row]*cellSize+cellSize, cellSize)).(*ebiten.Image), &options)
			}
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}

func (f *field) drawElevator(screen *ebiten.Image) {
	for line := 0; line < gridHeight; line++ {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, float64(line*cellSize))
		for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
			screen.DrawImage(spriteSheetImage.SubImage(image.Rect(f.elevator[line][row]*cellSize, 2*cellSize, f.elevator[line][row]*cellSize+cellSize, 3*cellSize)).(*ebiten.Image), &options)
			options.GeoM.Translate(float64(cellSize), 0)
		}
	}
}

func (f *field) setFallingLevel() bool {
	f.yposition = 0

	for line := 0; line < len(f.walls); line++ {
		for row := 0; row < len(f.walls[0]); row++ {
			f.oldWalls[line][row] = f.walls[line][row]
			f.walls[line][row] = noWallTile
		}
	}

	for line := 0; line < len(f.wallTiles); line++ {
		for row := 0; row < len(f.wallTiles[0]); row++ {
			f.oldWallTiles[line][row] = f.wallTiles[line][row]
			f.wallTiles[line][row] = emptySmallTile
		}
	}

	f.fallingLevelNum++

	switch f.fallingLevelNum {
	case 1:
		f.walls = fallingLevel1
		f.fallingLevelNum = 4
	case 2:
		f.walls = fallingLevel2
	case 3:
		f.walls = fallingLevel3
	case 4:
		f.walls = fallingLevel4
	case 5:
		f.walls = fallingLevel5
	default:
		return true
	}

	for line := 0; line < gridHeight; line++ {
		f.wallTiles[line*2][0] = emptySmallTile
		f.wallTiles[line*2][1] = emptySmallTile
		f.wallTiles[line*2+1][0] = emptySmallTile
		f.wallTiles[line*2+1][1] = emptySmallTile
		f.wallTiles[line*2][((gridWidth+leftMargin+rightMargin)-1)*2] = emptySmallTile
		f.wallTiles[line*2][((gridWidth+leftMargin+rightMargin)-1)*2+1] = emptySmallTile
		f.wallTiles[line*2+1][((gridWidth+leftMargin+rightMargin)-1)*2] = emptySmallTile
		f.wallTiles[line*2+1][((gridWidth+leftMargin+rightMargin)-1)*2+1] = emptySmallTile
		for row := 1; row < gridWidth+leftMargin+rightMargin-1; row++ {

			if f.walls[line][row] == wallTile {
				var topLeftTile int
				if line == 0 || f.walls[line-1][row] == wallTile {
					if f.walls[line][row-1] == wallTile {
						if line == 0 || f.walls[line-1][row-1] == wallTile {
							topLeftTile = fullTile1
						} else {
							topLeftTile = topLeftSmallCornerTile
						}
					} else {
						topLeftTile = leftEdgeTile1
					}
				} else {
					if f.walls[line][row-1] == wallTile {
						topLeftTile = topEdgeTile1
					} else {
						topLeftTile = topLeftCornerTile
					}
				}
				f.wallTiles[line*2][row*2] = topLeftTile

				var topRightTile int
				if line == 0 || f.walls[line-1][row] == wallTile {
					if f.walls[line][row+1] == wallTile {
						if line == 0 || f.walls[line-1][row+1] == wallTile {
							topRightTile = fullTile2
						} else {
							topRightTile = topRightSmallCornerTile
						}
					} else {
						topRightTile = rightEdgeTile1
					}
				} else {
					if f.walls[line][row+1] == wallTile {
						topRightTile = topEdgeTile2
					} else {
						topRightTile = topRightCornerTile
					}
				}
				f.wallTiles[line*2][row*2+1] = topRightTile

				var bottomRightTile int
				if line == gridHeight-1 || f.walls[line+1][row] == wallTile {
					if f.walls[line][row+1] == wallTile {
						if line == gridHeight-1 || f.walls[line+1][row+1] == wallTile {
							bottomRightTile = fullTile2
						} else {
							bottomRightTile = bottomRightSmallCornerTile
						}
					} else {
						bottomRightTile = rightEdgeTile2
					}
				} else {
					if f.walls[line][row+1] == wallTile {
						bottomRightTile = bottomEdgeTile2
					} else {
						bottomRightTile = bottomRightCornerTile
					}
				}
				f.wallTiles[line*2+1][row*2+1] = bottomRightTile

				var bottomLeftTile int
				if line == gridHeight-1 || f.walls[line+1][row] == wallTile {
					if f.walls[line][row-1] == wallTile {
						if line == gridHeight-1 || f.walls[line+1][row-1] == wallTile {
							bottomLeftTile = fullTile2
						} else {
							bottomLeftTile = bottomLeftSmallCornerTile
						}
					} else {
						bottomLeftTile = leftEdgeTile2
					}
				} else {
					if f.walls[line][row-1] == wallTile {
						bottomLeftTile = bottomEdgeTile1
					} else {
						bottomLeftTile = bottomLeftCornerTile
					}
				}
				f.wallTiles[line*2+1][row*2] = bottomLeftTile
			} else {
				f.wallTiles[line*2][row*2] = emptySmallTile
				f.wallTiles[line*2][row*2+1] = emptySmallTile
				f.wallTiles[line*2+1][row*2] = emptySmallTile
				f.wallTiles[line*2+1][row*2+1] = emptySmallTile
			}
		}
	}
	return false
}

func (f *field) drawWalls(screen *ebiten.Image) {
	if !f.isTransition {
		for line := 0; line < gridHeight*2; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize/2))
			for row := 0; row < (gridWidth+leftMargin+rightMargin)*2; row++ {
				if f.wallTiles[line][row] != emptySmallTile {
					xstart := (f.wallTiles[line][row] % (emptySmallTile / 2)) * cellSize / 2
					ystart := 4*cellSize + (f.wallTiles[line][row]/(emptySmallTile/2))*cellSize/2
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(xstart, ystart, xstart+cellSize/2, ystart+cellSize/2)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize/2), 0)
			}
		}
		for line := 0; line < gridHeight; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize))
			for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
				if f.walls[line][row] == diamondTile {
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(diamond*cellSize, cellSize, diamond*cellSize+cellSize, 2*cellSize)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize), 0)
			}
		}
	} else {
		for line := 0; line < gridHeight*2; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize/2+gridHeight*cellSize-cellSize)+f.yposition)
			for row := 0; row < (gridWidth+leftMargin+rightMargin)*2; row++ {
				if f.wallTiles[line][row] != emptySmallTile {
					xstart := (f.wallTiles[line][row] % (emptySmallTile / 2)) * cellSize / 2
					ystart := 4*cellSize + (f.wallTiles[line][row]/(emptySmallTile/2))*cellSize/2
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(xstart, ystart, xstart+cellSize/2, ystart+cellSize/2)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize/2), 0)
			}
		}
		for line := 0; line < gridHeight*2; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize/2)+f.yposition)
			for row := 0; row < (gridWidth+leftMargin+rightMargin)*2; row++ {
				if f.oldWallTiles[line][row] != emptySmallTile {
					xstart := (f.oldWallTiles[line][row] % (emptySmallTile / 2)) * cellSize / 2
					ystart := 4*cellSize + (f.oldWallTiles[line][row]/(emptySmallTile/2))*cellSize/2
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(xstart, ystart, xstart+cellSize/2, ystart+cellSize/2)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize/2), 0)
			}
		}
		for line := 0; line < gridHeight; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize)+f.yposition)
			for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
				if f.oldWalls[line][row] == diamondTile {
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(diamond*cellSize, cellSize, diamond*cellSize+cellSize, 2*cellSize)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize), 0)
			}
		}
		for line := 0; line < gridHeight; line++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(0, float64(line*cellSize+gridHeight*cellSize-cellSize)+f.yposition)
			for row := 0; row < gridWidth+leftMargin+rightMargin; row++ {
				if f.walls[line][row] == diamondTile {
					screen.DrawImage(spriteSheetImage.SubImage(image.Rect(diamond*cellSize, cellSize, diamond*cellSize+cellSize, 2*cellSize)).(*ebiten.Image), &options)
				}
				options.GeoM.Translate(float64(cellSize), 0)
			}
		}
	}
}

func (f *field) updateYPosition() {
	f.yposition += f.transitionSpeed
}

func (f *field) moveDone() bool {
	return false
}
