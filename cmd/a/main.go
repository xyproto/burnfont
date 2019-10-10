package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/rune2image"
	"image/png"
	"os"
)

// Output an image of a red "a"

func main() {

	var (
		outputFilename string
		version        bool
	)

	flag.StringVar(&outputFilename, "o", "-", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")

	flag.Parse()

	if version {
		fmt.Println("a 1.0.0")
		os.Exit(0)
	}

	// Prepare to output the new PNG data to either stdout or to file
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

	// Convert the image to only use the given palette
	m, err := rune2image.Image('a')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Output the image
	if err := png.Encode(f, m); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
