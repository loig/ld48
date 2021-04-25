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
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

type soundManager struct {
	audioContext     *audio.Context
	soundPlayer      *audio.Player
	musicPlayers     [3]*audio.Player
	currentMoveSound int
}

const (
	rockSound int = iota
	diamondCatchSound
	surpriseSound
	questionSound
	agreeSound
	waterSound
	deathSound
	earthquakeSound
	playerMoveSound
	playerMoveSound2
	playerMoveSound3
	playerMoveSound4
	playerMoveSound5
	playerMoveSound6
	playerMoveSound7
	numSound
)

// loop the music
func (g *game) updateMusic(pID int) {
	if g.audio.musicPlayers[pID] != nil {
		if !g.audio.musicPlayers[pID].IsPlaying() {
			g.audio.musicPlayers[pID].Rewind()
			g.audio.musicPlayers[pID].Play()
		}
	} else {
		var error error
		g.audio.musicPlayers[pID], error = audio.NewPlayer(g.audio.audioContext, infiniteMusic[pID])
		if error != nil {
			log.Panic("Audio problem:", error)
		}
		g.audio.musicPlayers[pID].Play()
	}
}

// stop the music
func (g *game) stopMusic(pID int) {
	if g.audio.musicPlayers[pID] != nil && g.audio.musicPlayers[pID].IsPlaying() {
		g.audio.musicPlayers[pID].Pause()
	}
}

// stop the current non-overlaying sound
func (g *game) stopSound() {
	if g.audio.soundPlayer != nil && g.audio.soundPlayer.IsPlaying() {
		error := g.audio.soundPlayer.Close()
		if error != nil {
			log.Panic("Sound problem:", error)
		}
	}
}

// play a sound, telling if it should overlay
// with other sounds or not
func (g *game) playSound(sound int, overlaying bool) {
	if sound == playerMoveSound {
		sound = g.audio.currentMoveSound
		g.audio.currentMoveSound++
		if g.audio.currentMoveSound > playerMoveSound7 {
			g.audio.currentMoveSound = playerMoveSound
		}
	}
	var soundBytes []byte = soundsBytes[sound]
	var error error
	if overlaying {
		soundPlayer := audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
		soundPlayer.Play()
	} else {
		if g.audio.soundPlayer != nil && g.audio.soundPlayer.IsPlaying() {
			error = g.audio.soundPlayer.Close()
			if error != nil {
				log.Panic("Sound problem:", error)
			}
		}
		g.audio.soundPlayer = audio.NewPlayerFromBytes(g.audio.audioContext, soundBytes)
		g.audio.soundPlayer.Play()
	}
}

// load all audio assets
func (g *game) initAudio() {

	var error error
	g.audio.audioContext = audio.NewContext(44100)

	g.audio.currentMoveSound = playerMoveSound

	// music
	sound, error := mp3.Decode(g.audio.audioContext, bytes.NewReader(elevatorMusic))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("20s")
	duration := tduration.Seconds()
	theBytes := int64(math.Round(duration * 4 * float64(44100)))
	infiniteMusic[0] = audio.NewInfiniteLoop(sound, theBytes)

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(fallingObjectsMusic))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ = time.ParseDuration("9.6s")
	duration = tduration.Seconds()
	theBytes = int64(math.Round(duration * 4 * float64(44100)))
	infiniteMusic[1] = audio.NewInfiniteLoop(sound, theBytes)

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(fallingMusic))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	infiniteMusic[2] = audio.NewInfiniteLoop(sound, theBytes)

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(rockSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[rockSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(diamondCatchSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[diamondCatchSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(surpriseSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[surpriseSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(questionSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[questionSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(agreeSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[agreeSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(waterSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[waterSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(deathSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[deathSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(earthquakeSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[earthquakeSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSoundBytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound2Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound2], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound3Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound3], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound4Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound4], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound5Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound5], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound6Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound6], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}

	sound, error = mp3.Decode(g.audio.audioContext, bytes.NewReader(playerMoveSound7Bytes))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	soundsBytes[playerMoveSound7], error = ioutil.ReadAll(sound)
	if error != nil {
		log.Panic("Audio problem:", error)
	}
}
