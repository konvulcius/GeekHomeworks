package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func switchColor(color *string) {
	switch *color {
	case "white":
		*color = "black"
	default:
		*color = "white"
	}
}

func main() {
	boardSize := 400
	colors := make(map[string]color.RGBA, 2)

	colors["white"] = color.RGBA{255, 246, 195, 255}
	colors["black"] = color.RGBA{167, 79, 0, 255}

	board := image.NewRGBA(image.Rect(0, 0, boardSize, boardSize))
	boxSize := boardSize / 8
	curColor := "white"

	for x := 0; x < boardSize; x += boxSize {
		for y := 0; y < boardSize; y += boxSize {
			draw.Draw(board, image.Rect(x, y, x+boxSize, y+boxSize), &image.Uniform{colors[curColor]}, image.ZP, draw.Src)
			switchColor(&curColor)
		}
		switchColor(&curColor)
	}

	file, err := os.Create("check_mate.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, board)
}
