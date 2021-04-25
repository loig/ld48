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
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

//go:embed assets/sprites.png
var spriteSheetBytes []byte
var spriteSheetImage *ebiten.Image

//go:embed assets/end.png
var endBytes []byte
var endImage *ebiten.Image

//go:embed assets/title.png
var titleBytes []byte
var titleImage *ebiten.Image

var blackImage *ebiten.Image

//go:embed assets/bossassenceur.mp3
var elevatorMusic []byte

//go:embed assets/bossassenceur2.mp3
var fallingObjectsMusic []byte

//go:embed assets/bossassenceur3.mp3
var fallingMusic []byte

var infiniteMusic [3]*audio.InfiniteLoop

var soundsBytes [numSound][]byte

//go:embed assets/rockSound.mp3
var rockSoundBytes []byte

//go:embed assets/diamondCatchSound.mp3
var diamondCatchSoundBytes []byte

//go:embed assets/surpriseSound.mp3
var surpriseSoundBytes []byte

//go:embed assets/questionSound.mp3
var questionSoundBytes []byte

//go:embed assets/agreeSound.mp3
var agreeSoundBytes []byte

//go:embed assets/waterSound.mp3
var waterSoundBytes []byte

//go:embed assets/deathSound.mp3
var deathSoundBytes []byte

//go:embed assets/earthquakeSound.mp3
var earthquakeSoundBytes []byte

//go:embed assets/playerMoveSound.mp3
var playerMoveSoundBytes []byte

//go:embed assets/playerMoveSound2.mp3
var playerMoveSound2Bytes []byte

//go:embed assets/playerMoveSound3.mp3
var playerMoveSound3Bytes []byte

//go:embed assets/playerMoveSound4.mp3
var playerMoveSound4Bytes []byte

//go:embed assets/playerMoveSound5.mp3
var playerMoveSound5Bytes []byte

//go:embed assets/playerMoveSound6.mp3
var playerMoveSound6Bytes []byte

//go:embed assets/playerMoveSound7.mp3
var playerMoveSound7Bytes []byte

func loadAssets() {
	var err error
	spriteSheetDecoded, _, err := image.Decode(bytes.NewReader(spriteSheetBytes))
	if err != nil {
		log.Fatal(err)
	}
	spriteSheetImage = ebiten.NewImageFromImage(spriteSheetDecoded)

	blackImage = ebiten.NewImage(screenWidth, screenHeight)
	blackImage.Fill(color.Black)

	endDecoded, _, err := image.Decode(bytes.NewReader(endBytes))
	if err != nil {
		log.Fatal(err)
	}
	endImage = ebiten.NewImageFromImage(endDecoded)

	titleDecoded, _, err := image.Decode(bytes.NewReader(titleBytes))
	if err != nil {
		log.Fatal(err)
	}
	titleImage = ebiten.NewImageFromImage(titleDecoded)
}
