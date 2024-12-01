package main

import (
	"image/color"
	"os"
	"path/filepath"
)

const EXPORT bool = false

const EMPTY uint8 = 0
const KING uint8 = 1
const QUEEN uint8 = 2
const BISHOP uint8 = 3
const KNIGHT uint8 = 4
const ROOK uint8 = 5
const PAWN uint8 = 6

const WHITE uint8 = 8
const BLACK uint8 = 16

const typemask uint8 = 0b000_00_111
const colormask uint8 = 0b000_11_000

const STARTFEN string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var WHITERGB color.RGBA = color.RGBA{241, 217, 192, 255}
var BLACKRGB color.RGBA = color.RGBA{169, 122, 101, 255}
var SELECTRGB color.RGBA = color.RGBA{249, 224, 118, 255}

func GetPath() string {
	filename, err := os.Executable()
	if err != nil {
		return ""
	}

	filename = filepath.Dir(filename) // exit path twice because i would else be in /src
	if !EXPORT {
		filename = filepath.Dir(filename)
	}
	return filename
}

func AppendFilePath(path string, toAppend []string) string {
	for _, appendage := range toAppend {
		path += string(os.PathSeparator) + appendage
	}
	return path
}

type Move struct {
	from    int8
	to      int8
	special int8
}
