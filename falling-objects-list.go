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
}

func (fOL *fallingObjectsList) update(mayAddObject bool) {
	for objectID := 0; objectID < len(fOL.objects); objectID++ {
		fOL.objects[objectID].update()
	}
	if mayAddObject {
		if fOL.objectsToAdd > 0 {
			fOL.addFallingObjects()
		} else {
			if fOL.noAlive() {
				fOL.setLevel()
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
	}
}

func initFallingObjectsList(numObjects int) fallingObjectsList {
	fOL := fallingObjectsList{}
	fOL.objects = make([]fallingObject, 0, numObjects)
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
		for spawned := 0; spawned < fOL.batchSize; spawned++ {
			xposition := fOL.spawnPositions[fOL.spawnID%len(fOL.spawnPositions)]
			fOL.spawnID++
			if fOL.spawnID >= len(fOL.spawnPositions) {
				needShuffle = true
			}
			objectID := fOL.nextAvailable()
			if objectID < len(fOL.objects) {
				fOL.objects[objectID].reset(xposition, fOL.getYSpeed())
			} else {
				fOL.objects = append(fOL.objects, newFallingObject(xposition, fOL.getYSpeed()))
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

func (fOL *fallingObjectsList) setLevel() {
	if fOL.levelNum < elevatorNumLevelsPhase1+elevatorNumLevelsPhase2 {
		fOL.levelNum++
		switch fOL.levelNum {
		case 1:
			fOL.spawnChances = 2
			fOL.maxSpawnInterval = 2
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
			fOL.objectSpeed = 9
		case 2:
			fOL.spawnChances = 1
			fOL.maxSpawnInterval = 1
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
			fOL.objectSpeed = 11
		case 3:
			fOL.batchSize = 3
			fOL.minSpawnInterval = 2
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
		case 4:
			fOL.batchSize = 5
			fOL.minSpawnInterval = 3
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
		case 5:
			fOL.batchSize = 7
			fOL.minSpawnInterval = 4
			fOL.objectsToAdd = elevatorNumObjectsPerLevel
		}
	}
}
