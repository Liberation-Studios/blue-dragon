package tileEngine

import "fmt"

func BigTileArrayInput() {
	const tileChunkSize uint8 = 4
	tileChunk := [tileChunkSize * tileChunkSize]int{0}

	fmt.Println(tileChunk)
}

func TranslateCoordinates() {
	var x = 4
	var y = 2
	var tileChunkSize = 4

}
