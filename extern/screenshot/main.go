package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/kbinani/screenshot"
)

func main() {
	n := screenshot.NumActiveDisplays()
	log.Printf("Active Displays: %d\n", n)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		log.Printf("Bounds: %#v\n", bounds)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		png.Encode(file, img)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}
