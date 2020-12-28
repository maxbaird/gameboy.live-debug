package gb

import (
	"fmt"
	"strings"
)

const PIXEL_AMT = 160
const spriteTileAmt = 40
const tilePixelAmt = 7

type pixelData [spriteTileAmt][tilePixelAmt]int

type gpuState struct {
	scrollx          int
	scrolly          int
	windowx          int
	windowy          int
	tileData         int
	backgroundMemory int
	yPosition        int
	tileRow          int

	xPosition    [PIXEL_AMT]int
	tileColumn   [PIXEL_AMT]int
	tileNumber   [PIXEL_AMT]int
	tileAddress  [PIXEL_AMT]int
	tileLocation [PIXEL_AMT]int
	currentLine  [PIXEL_AMT]int
	pixelData1   [PIXEL_AMT]int
	pixelData2   [PIXEL_AMT]int
	colourBit    [PIXEL_AMT]int
	colourNumber [PIXEL_AMT]int
	red          [PIXEL_AMT]int
	green        [PIXEL_AMT]int
	blue         [PIXEL_AMT]int
}

type spriteState struct {
	use8x16    bool
	lcdControl int

	index        [spriteTileAmt]int
	yPosition    [spriteTileAmt]int
	xPosition    [spriteTileAmt]int
	tileLocation [spriteTileAmt]int
	attributes   [spriteTileAmt]int
	yFlip        [spriteTileAmt]int
	xFlip        [spriteTileAmt]int
	scanline     [spriteTileAmt]int
	ysize        [spriteTileAmt]int
	line         [spriteTileAmt]int
	dataAddress  [spriteTileAmt]int
	data1        [spriteTileAmt]int
	data2        [spriteTileAmt]int

	colorbit     pixelData
	colorNum     pixelData
	colorAddress pixelData
	pixel        pixelData
}

func dumpSpriteState(SpriteState *spriteState, filename string) {
	fmt.Println("Dumping sprite data: ", filename)

	var str strings.Builder
	str.WriteString(fmt.Sprintf("use8x16: %v\n", SpriteState.use8x16))
	str.WriteString(fmt.Sprintf("lcdControl: %d\n", SpriteState.lcdControl))
	str.WriteString(fmt.Sprintf("===========\n"))

	appendSpriteData(&str, "index", SpriteState.index[:])
	appendSpriteData(&str, "yPosition", SpriteState.yPosition[:])
	appendSpriteData(&str, "xPosition", SpriteState.xPosition[:])
	appendSpriteData(&str, "tileLocation", SpriteState.tileLocation[:])
	appendSpriteData(&str, "attributes", SpriteState.attributes[:])
	appendSpriteData(&str, "yFlip", SpriteState.yFlip[:])
	appendSpriteData(&str, "xFlip", SpriteState.xFlip[:])
	appendSpriteData(&str, "scanline", SpriteState.scanline[:])
	appendSpriteData(&str, "ysize", SpriteState.ysize[:])
	appendSpriteData(&str, "line", SpriteState.line[:])
	appendSpriteData(&str, "dataAddress", SpriteState.dataAddress[:])
	appendSpriteData(&str, "data1", SpriteState.data1[:])
	appendSpriteData(&str, "data2", SpriteState.data2[:])

	appendPixelSpriteData(&str, "colorbit", SpriteState.colorbit)
	appendPixelSpriteData(&str, "colorNum", SpriteState.colorNum)
	appendPixelSpriteData(&str, "colorAddress", SpriteState.colorAddress)
	appendPixelSpriteData(&str, "pixel", SpriteState.pixel)

	writeToFile(filename, str.String())

	fmt.Println("Dumped sprite data: ", filename)
}

func appendSpriteData(str *strings.Builder, name string, spriteData []int) {
	for i := 0; i < spriteTileAmt; i++ {
		str.WriteString(fmt.Sprintf("%s%d: %d\n", name, i, spriteData[i]))
	}
	str.WriteString(fmt.Sprintf("===========\n"))
}

func appendPixelSpriteData(str *strings.Builder, name string, p pixelData) {
	for i, data := range p {
		for j, data2 := range data {
			str.WriteString(fmt.Sprintf("%v%v-%v: %v\n", name, i, j, data2))
		}
	}
	str.WriteString(fmt.Sprintf("===========\n"))
}

func dumpGPUState(GPUState *gpuState, iteration int, filename string) {
	fmt.Println("Dumping GPU state: ", filename)
	var str strings.Builder
	str.WriteString(fmt.Sprintf("Iteration: %d\n", iteration))
	str.WriteString(fmt.Sprintf("scrollx: %d\n", GPUState.scrollx))
	str.WriteString(fmt.Sprintf("scrolly: %d\n", GPUState.scrolly))
	str.WriteString(fmt.Sprintf("windowx: %d\n", GPUState.windowx))
	str.WriteString(fmt.Sprintf("windowy: %d\n", GPUState.windowy))
	str.WriteString(fmt.Sprintf("tileData: %d\n", GPUState.tileData))
	str.WriteString(fmt.Sprintf("backgroundMemory: %d\n", GPUState.backgroundMemory))
	str.WriteString(fmt.Sprintf("yPosition: %d\n", GPUState.yPosition))
	str.WriteString(fmt.Sprintf("tileRow: %d\n", GPUState.tileRow))
	str.WriteString(fmt.Sprintf("===========\n"))

	appendPixelData(&str, "xPosition", GPUState.xPosition[:])
	appendPixelData(&str, "tileColumn", GPUState.tileColumn[:])
	appendPixelData(&str, "tileNumber", GPUState.tileNumber[:])
	appendPixelData(&str, "tileAddress", GPUState.tileAddress[:])
	appendPixelData(&str, "tileLocation", GPUState.tileLocation[:])
	appendPixelData(&str, "currentLine", GPUState.currentLine[:])
	appendPixelData(&str, "pixelData1", GPUState.pixelData1[:])
	appendPixelData(&str, "pixelData2", GPUState.pixelData2[:])
	appendPixelData(&str, "colourBit", GPUState.colourBit[:])
	appendPixelData(&str, "colourNumber", GPUState.colourNumber[:])
	appendColours(&str, GPUState)

	writeToFile(filename, str.String())
	fmt.Println("Dumped GPU state: ", filename)
}

func appendPixelData(str *strings.Builder, name string, pixelData []int) {
	for i := 0; i < PIXEL_AMT; i++ {
		str.WriteString(fmt.Sprintf("%s%d: %d\n", name, i, pixelData[i]))
	}
	str.WriteString(fmt.Sprintf("===========\n"))
}

func appendColours(str *strings.Builder, GPUState *gpuState) {
	for i := 0; i < PIXEL_AMT; i++ {
		str.WriteString(fmt.Sprintf("red: %d, green: %d, blue: %d\n", GPUState.red[i], GPUState.green[i], GPUState.blue[i]))
	}
}
