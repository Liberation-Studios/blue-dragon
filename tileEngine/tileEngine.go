package tileEngine

import "fmt"

func BigTileArrayInput() {
	const tileChunkSize uint8 = 4
	tileChunk := [tileChunkSize * tileChunkSize]int{0}

	fmt.Println(tileChunk)
}
