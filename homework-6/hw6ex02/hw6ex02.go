package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{255, 0, 0, 150}
	blue := color.RGBA{0, 0, 255, 150}
	rectImg := image.NewRGBA(image.Rect(0, 0, 300, 300))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	for setY := 39; setY < 300; setY += 60 {
		for x := 0; x < 300; x++ {
			for fat := setY; fat < setY+3; fat++ {
				rectImg.Set(x, fat, red)
			}
		}
	}

	for setX := 49; setX < 300; setX += 100 {
		for y := 0; y < 300; y++ {
			for fat := setX; fat < setX+3; fat++ {
				rectImg.Set(fat, y, blue)
			}
		}
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
