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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type fallingObjectsList struct {
	objects          []fallingObject
	explosions       []explosion
	objectsToAdd     int
	spawnChances     int
	maxSpawnInterval int
	minSpawnInterval int
	sinceLastSpawn   int
	spawnPositions   [gridWidth]int
	spawnID          int
	levelNum         int
	objectSpeed      float64
	batchSize        int
	numDiamond       int
}

func (fOL *fallingObjectsList) update(mayAddObject bool, earthShake *bool) {
	for objectID := 0; objectID < len(fOL.objects); objectID++ {
		currentlyAlive := fOL.objects[objectID].alive
		fOL.objects[objectID].update()
		if currentlyAlive && !fOL.objects[objectID].alive {
			fOL.explosions[objectID].reset(fOL.objects[objectID].objectType, float64(fOL.objects[objectID].xposition*cellSize), fOL.objects[objectID].yposition)
		}
		fOL.explosions[objectID].update()
	}
	if mayAddObject {
		if fOL.objectsToAdd > 0 {
			fOL.addFallingObjects()
		} else {
			if fOL.noAlive() {
				fOL.setLevel(earthShake)
			}
		}
	}
}

func (fOL *fallingObjectsList) doneFalling() bool {
	return fOL.levelNum >= elevatorNumLevelsPhase1+elevatorNumLevelsPhase2 &&
		fOL.objectsToAdd <= 0 && fOL.noAlive()
}

func (fOL *fallingObjectsList) draw(screen *ebiten.Image) {
	for objectID := 0; objectID < len(fOL.objects); objectID++ {
		fOL.objects[objectID].draw(screen)
		fOL.explosions[objectID].draw(screen)
	}
}

func initFallingObjectsList(numObjects int) fallingObjectsList {
	fOL := fallingObjectsList{}
	fOL.objects = make([]fallingObject, 0, numObjects)
	fOL.explosions = make([]explosion, 0, numObjects)
	fOL.objectsToAdd = numObjects
	fOL.spawnChances = initialSpawnChances
	fOL.maxSpawnInterval = initialSpawnInterval
	fOL.minSpawnInterval = 1
	fOL.sinceLastSpawn = fOL.maxSpawnInterval
	for spawnID := 0; spawnID < len(fOL.spawnPositions); spawnID++ {
		fOL.spawnPositions[spawnID] = spawnID
	}
	rand.Shuffle(len(fOL.spawnPositions), func(i, j int) {
		fOL.spawnPositions[i], fOL.spawnPositions[j] = fOL.spawnPositions[j], fOL.spawnPositions[i]
	})
	fOL.spawnID = 0
	fOL.objectSpeed = initialObjectSpeed
	fOL.batchSize = 1
	fOL.numDiamond = 1
	return fOL
}

func (fOL *fallingObjectsList) nextAvailable() int {
	objectID := 0
	for objectID < len(fOL.objects) && fOL.objects[objectID].alive {
		objectID++
	}
	return objectID
}

func (fOL *fallingObjectsList) noAlive() bool {
	for objectID := 0; objectID < len(fOL.objects); objectID++ {
		if fOL.objects[objectID].alive {
			return false
		}
	}
	return true
}

func (fOL *fallingObjectsList) addFallingObjects() {
	if fOL.sinceLastSpawn >= fOL.minSpawnInterval &&
		(fOL.sinceLastSpawn >= fOL.maxSpawnInterval || rand.Intn(fOL.spawnChances) == 0) {
		needShuffle := false
		diamondInBatch := false
		for spawned := 0; spawned < fOL.batchSize; spawned++ {
			xposition := fOL.spawnPositions[fOL.spawnID%len(fOL.spawnPositions)]
			fOL.spawnID++
			if fOL.spawnID >= len(fOL.spawnPositions) {
				needShuffle = true
			}
			isDiamond := !diamondInBatch && fOL.numDiamond > 0 &&
				((fOL.objectsToAdd <= fOL.numDiamond) || rand.Intn(5) == 0)
			if isDiamond {
				fOL.numDiamond--
				diamondInBatch = true
			}
			objectID := fOL.nextAvailable()
			if objectID < len(fOL.objects) {
				fOL.objects[objectID].reset(xposition, fOL.getYSpeed(), isDiamond)
			} else {
				fOL.objects = append(fOL.objects, newFallingObject(xposition, fOL.getYSpeed(), isDiamond))
				fOL.explosions = append(fOL.explosions, explosion{})
			}
		}
		if needShuffle {
			fOL.spawnID = 0
			rand.Shuffle(len(fOL.spawnPositions), func(i, j int) {
				fOL.spawnPositions[i], fOL.spawnPositions[j] = fOL.spawnPositions[j], fOL.spawnPositions[i]
			})
		}
		fOL.objectsToAdd--
		fOL.sinceLastSpawn = 0
	} else {
		fOL.sinceLastSpawn++
	}
}

func (fOL *fallingObjectsList) getYSpeed() float64 {
	return fOL.objectSpeed
}

func (fOL *fallingObjectsList) setLevel(earthShake *bool) {
	if fOL.levelNum < elevatorNumLevelsPhase1+elevatorNumLevelsPhase2 {
		fOL.levelNum++
		if fOL.levelNum < 6 {
			*earthShake = true
		}
		switch fOL.levelNum {
		case 1:
			fOL.spawnChances = 2
			fOL.maxSpawnInterval = 2
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
			fOL.objectSpeed = 9
			fOL.numDiamond = 1
		case 2:
			fOL.batchSize = 2
			fOL.spawnChances = 1
			fOL.maxSpawnInterval = 1
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
			fOL.objectSpeed = 11
			fOL.numDiamond = 1
		case 3:
			fOL.batchSize = 3
			fOL.minSpawnInterval = 2
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
			fOL.numDiamond = 1
		case 4:
			fOL.batchSize = 5
			fOL.minSpawnInterval = 3
			fOL.objectsToAdd = elevatorNumObjectsPerLevel / 2
			fOL.numDiamond = 1
		case 5:
			fOL.batchSize = 7
			fOL.minSpawnInterval = 5
			fOL.objectSpeed = 8
			fOL.objectsToAdd = elevatorNumObjectsPerLevel / 3
			fOL.numDiamond = 0
		}
	}
}
