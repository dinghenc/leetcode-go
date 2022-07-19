package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"unicode/utf8"
)

func parseImage() {
	file, err := os.Open("D:\\Downloads\\quiz5-abc.png")
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	var imgBytes []byte
	for y := 0; y < height; y++ {
		for x := 0; x < width; x += 8 {
			var b byte
			for i := 0; i < 8; i++ {
				r, _, _, _ := img.At(x+i, y).RGBA()
				if r == 0 {
					b |= 1 << (8 - i - 1)
				}
			}
			imgBytes = append(imgBytes, b)
		}
	}

	b := imgBytes
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c", r)

		b = b[size:]
	}
}
