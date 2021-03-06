package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/burnfont"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	var (
		outputFilename string
		version        bool
	)

	flag.StringVar(&outputFilename, "o", "out.png", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")

	flag.Parse()

	if version {
		fmt.Println("scaled 0.0.1")
		os.Exit(0)
	}

	// Prepare to output the new PNG data to either stdout or to file
	var (
		f   *os.File
		err error
	)
	if outputFilename == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(outputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	w := 104
	h := 64
	scale := 4

	// Create a gray image
	grayImage := image.NewRGBA(image.Rect(0, 0, w*scale, h*scale))

	gray := color.RGBA{128, 128, 128, 255}
	draw.Draw(grayImage, grayImage.Bounds(), &image.Uniform{gray}, image.ZP, draw.Src)

	// Create an image where the text will be drawn
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	var (
		r byte = 255
		g byte
		b byte
		i int
	)
	alen := len(burnfont.Available)
OUT:
	for y := 0; y < 100; y += 8 {
		for x := 0; x < 100; x += 8 {
			// Convert the image to only use the given palette
			l := burnfont.Available[i]
			if err = burnfont.Draw(m, l, x, y, r, g, b); err != nil {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
				os.Exit(1)
			}
			i++
			if i >= alen {
				break OUT
			}
		}
	}

	// Draw the text image on top of the gray image
	draw.NearestNeighbor.Scale(grayImage, grayImage.Bounds(), m, m.Bounds(), draw.Over, nil)

	// Output the grayImage, with the text overlay
	if err := png.Encode(f, grayImage); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
