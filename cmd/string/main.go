package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/burnfont"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// Output an image of a red "a"

func main() {

	var (
		outputFilename string
		version        bool
	)

	flag.StringVar(&outputFilename, "o", "out.png", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")

	flag.Parse()

	if version {
		fmt.Println("string 0.0.1")
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
	h := 56

	// Create a gray image
	grayImage := image.NewRGBA(image.Rect(0, 0, w, h))

	gray := color.RGBA{128, 128, 128, 255}
	draw.Draw(grayImage, grayImage.Bounds(), &image.Uniform{gray}, image.ZP, draw.Src)

	// Create an image where the text will be drawn
	m := image.NewRGBA(image.Rect(0, 0, w, h))

	// Draw a string
	burnfont.DrawString(m, 0, 0, "HELLO", color.NRGBA{0, 0, 0, 255})

	// Draw the text image on top of the gray image
	draw.Draw(grayImage, grayImage.Bounds(), m, m.Bounds().Min, draw.Over)

	// Output the grayImage, with the text overlay
	if err := png.Encode(f, grayImage); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
