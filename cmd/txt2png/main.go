package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xyproto/burnfont"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "syntax: txt2png TEXTFILE [PNGFILE]")
		os.Exit(1)
	}
	text, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	pngFilename := "output.png"
	if len(os.Args) > 2 {
		pngFilename = os.Args[2]
	}
	burnfont.SavePNG(string(text), pngFilename)
}
