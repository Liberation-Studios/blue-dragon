package tileEngine

import (
	"fmt"
)

func BigTileArrayInput() {
	const tileChunkSize uint8 = 4
	tileChunk := [16]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 720, 11, 12, 13, 14, 15}

	TranslateCoordinates(tileChunkSize)
	fmt.Println(GetTile(tileChunk, 2, 2, tileChunkSize))
}

func TranslateCoordinates(tileChunkSize uint8) {
	var x uint8 = 2
	var y uint8 = 2
	var width uint8 = 4 // Width of the grid

	if x < 0 || x >= width || y < 0 || y >= width {
		fmt.Println("Coordinates are out of bounds")
		return
	}

	index := convertTo1DIndex(x, y, width)
	fmt.Println("1D Index:", index)

}

func convertTo1DIndex(x, y, columns uint8) uint8 {
	return x + y*columns
}

func GetTile(tileChunk [16]int, x, y, columns uint8) int {
	return tileChunk[convertTo1DIndex(x, y, columns)]

}
