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
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/background.png
var spriteSheetBytes []byte
var spriteSheetImage *ebiten.Image

func loadAssets() {
	var err error
	spriteSheetDecoded, _, err := image.Decode(bytes.NewReader(spriteSheetBytes))
	if err != nil {
		log.Fatal(err)
	}
	spriteSheetImage = ebiten.NewImageFromImage(spriteSheetDecoded)
}