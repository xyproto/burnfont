package rune2image

import (
	"errors"
	"fmt"
	"image"
	"image/color"
)

// Available is a slice with all available runes, for this package
var Available = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'Æ', 'Ø', 'Å', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'æ', 'ø', 'å', '.', ';', ',', '\'', '"', '*', '+', '!', '?', '-', '=', '_', '/', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '(', ')'}

func plotFontLine(s string, x, y int, m *image.RGBA, r, g, b byte) {
	for _, l := range s {
		if l == '*' {
			m.Set(x, y, color.NRGBA{r, g, b, 255})
		} else if l == '-' {
			m.Set(x, y, color.NRGBA{r, g, b, 64})
		}
		x++
	}
	return
}

// Draw will draw an image of the selected rune at (x,y).
// If the rune is not available, an error will be returned.
// r,g,b is the main color of the rune.
func Draw(m *image.RGBA, l rune, x, y int, r, g, b byte) error {
	switch l {
	case 'a':
		plotFontLine("***-", x+1, y+1, m, r, g, b)
		plotFontLine("--**", x+1, y+2, m, r, g, b)
		plotFontLine("-****", x, y+3, m, r, g, b)
		plotFontLine("**-**-", x, y+4, m, r, g, b)
		plotFontLine("-**-**", x, y+5, m, r, g, b)
	case 'A':
		plotFontLine("-**-", x+1, y, m, r, g, b)
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("**--**", x, y+2, m, r, g, b)
		plotFontLine("******", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("**  **", x, y+5, m, r, g, b)
	case 'b':
		plotFontLine("***", x, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("****-", x+1, y+2, m, r, g, b)
		plotFontLine("** **", x+1, y+3, m, r, g, b)
		plotFontLine("-**-**", x, y+4, m, r, g, b)
		plotFontLine("**-**-", x, y+5, m, r, g, b)
	case 'B':
		plotFontLine("*****-", x, y, m, r, g, b)
		plotFontLine("-**-**", x, y+1, m, r, g, b)
		plotFontLine("****-", x+1, y+2, m, r, g, b)
		plotFontLine("** **", x+1, y+3, m, r, g, b)
		plotFontLine("-** **", x, y+4, m, r, g, b)
		plotFontLine("*****-", x, y+5, m, r, g, b)
	case 'c':
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("**--**", x, y+2, m, r, g, b)
		plotFontLine("** ---", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'C':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**- -*", x, y+1, m, r, g, b)
		plotFontLine("**", x, y+2, m, r, g, b)
		plotFontLine("**", x, y+3, m, r, g, b)
		plotFontLine("**- -*", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'd':
		plotFontLine("***", x+3, y, m, r, g, b)
		plotFontLine("-**", x+3, y+1, m, r, g, b)
		plotFontLine("-*****", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("**  **", x, y+4, m, r, g, b)
		plotFontLine("-***-*", x, y+5, m, r, g, b)
	case 'D':
		plotFontLine("****-", x, y, m, r, g, b)
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("** -**", x, y+2, m, r, g, b)
		plotFontLine("** -**", x, y+3, m, r, g, b)
		plotFontLine("**-**-", x, y+4, m, r, g, b)
		plotFontLine("****-", x, y+5, m, r, g, b)
	case 'e':
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("** -**", x, y+2, m, r, g, b)
		plotFontLine("******", x, y+3, m, r, g, b)
		plotFontLine("**-", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'E':
		plotFontLine("-*****", x, y, m, r, g, b)
		plotFontLine("**---*", x, y+1, m, r, g, b)
		plotFontLine("****-", x, y+2, m, r, g, b)
		plotFontLine("**--", x, y+3, m, r, g, b)
		plotFontLine("**- **", x, y+4, m, r, g, b)
		plotFontLine("*****-", x, y+5, m, r, g, b)
	case 'f':
		plotFontLine("-***-", x+1, y, m, r, g, b)
		plotFontLine("-**--*", x, y+1, m, r, g, b)
		plotFontLine("****", x, y+2, m, r, g, b)
		plotFontLine("-**-", x, y+3, m, r, g, b)
		plotFontLine("**", x+1, y+4, m, r, g, b)
		plotFontLine("****", x, y+5, m, r, g, b)
	case 'F':
		plotFontLine("*****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("**-", x, y+2, m, r, g, b)
		plotFontLine("****", x, y+3, m, r, g, b)
		plotFontLine("**-", x, y+4, m, r, g, b)
		plotFontLine("*-", x, y+5, m, r, g, b)
	case 'g':
		plotFontLine("-**-**", x, y+1, m, r, g, b)
		plotFontLine("** **-", x, y+2, m, r, g, b)
		plotFontLine("-****", x, y+3, m, r, g, b)
		plotFontLine("-**", x+2, y+4, m, r, g, b)
		plotFontLine("****-", x, y+5, m, r, g, b)
	case 'G':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("**", x, y+2, m, r, g, b)
		plotFontLine("** ***", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-***-*", x, y+5, m, r, g, b)
	case 'h':
		plotFontLine("***", x, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("****-", x+1, y+2, m, r, g, b)
		plotFontLine("**-**", x+1, y+3, m, r, g, b)
		plotFontLine("-** **", x, y+4, m, r, g, b)
		plotFontLine("*** **", x, y+5, m, r, g, b)
	case 'H':
		plotFontLine("**  **", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("******", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("**  **", x, y+4, m, r, g, b)
		plotFontLine("**  **", x, y+5, m, r, g, b)
	case 'i':
		plotFontLine("**", x+2, y, m, r, g, b)
		plotFontLine("--", x+2, y+1, m, r, g, b)
		plotFontLine("***", x+1, y+2, m, r, g, b)
		plotFontLine("-**", x+1, y+3, m, r, g, b)
		plotFontLine("-**-", x+1, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case 'I':
		plotFontLine("****", x+1, y, m, r, g, b)
		plotFontLine("-**-", x+1, y+1, m, r, g, b)
		plotFontLine("**", x+2, y+2, m, r, g, b)
		plotFontLine("**", x+2, y+3, m, r, g, b)
		plotFontLine("-**-", x+1, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case 'j':
		plotFontLine("**", x+3, y, m, r, g, b)
		plotFontLine("--", x+3, y+1, m, r, g, b)
		plotFontLine("***", x+2, y+2, m, r, g, b)
		plotFontLine("-**", x+2, y+3, m, r, g, b)
		plotFontLine("**-**", x, y+4, m, r, g, b)
		plotFontLine("-***-", x, y+5, m, r, g, b)
	case 'J':
		plotFontLine("****", x+2, y, m, r, g, b)
		plotFontLine("-**-", x+2, y+1, m, r, g, b)
		plotFontLine("**", x+3, y+2, m, r, g, b)
		plotFontLine("** **", x, y+3, m, r, g, b)
		plotFontLine("**-**", x, y+4, m, r, g, b)
		plotFontLine("-***-", x, y+5, m, r, g, b)
	case 'k':
		plotFontLine("***", x, y, m, r, g, b)
		plotFontLine("-**", x, y+1, m, r, g, b)
		plotFontLine("**-**", x+1, y+2, m, r, g, b)
		plotFontLine("****-", x+1, y+3, m, r, g, b)
		plotFontLine("-**-**", x, y+4, m, r, g, b)
		plotFontLine("*** **", x, y+5, m, r, g, b)
	case 'K':
		plotFontLine("***-**", x, y, m, r, g, b)
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("***-", x+1, y+2, m, r, g, b)
		plotFontLine("****-", x+1, y+3, m, r, g, b)
		plotFontLine("-**-**", x, y+4, m, r, g, b)
		plotFontLine("*** **", x, y+5, m, r, g, b)
	case 'l':
		plotFontLine("***", x+1, y, m, r, g, b)
		plotFontLine("-**", x+1, y+1, m, r, g, b)
		plotFontLine("**", x+2, y+2, m, r, g, b)
		plotFontLine("**", x+2, y+3, m, r, g, b)
		plotFontLine("-**-", x+1, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case 'L':
		plotFontLine("****", x, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("**", x+1, y+2, m, r, g, b)
		plotFontLine("**", x+1, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("******", x, y+5, m, r, g, b)
	case 'm':
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("-*****", x, y+2, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+3, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+4, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+5, m, r, g, b)
	case 'M':
		plotFontLine("**-**", x+1, y, m, r, g, b)
		plotFontLine("**-**", x+1, y+1, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+2, m, r, g, b)
		plotFontLine("*- -*", x+1, y+3, m, r, g, b)
		plotFontLine("*- -*", x+1, y+4, m, r, g, b)
		plotFontLine("*   *", x+1, y+5, m, r, g, b)
	case 'n':
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("-*****", x, y+2, m, r, g, b)
		plotFontLine("**-**", x+1, y+3, m, r, g, b)
		plotFontLine("** **", x+1, y+4, m, r, g, b)
		plotFontLine("** **", x+1, y+5, m, r, g, b)
	case 'N':
		plotFontLine("**  **", x, y, m, r, g, b)
		plotFontLine("**- **", x, y+1, m, r, g, b)
		plotFontLine("***-**", x, y+2, m, r, g, b)
		plotFontLine("**-***", x, y+3, m, r, g, b)
		plotFontLine("** -**", x, y+4, m, r, g, b)
		plotFontLine("**  **", x, y+5, m, r, g, b)
	case 'o':
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("**--**", x, y+2, m, r, g, b)
		plotFontLine("**  **", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'O':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("**  **", x, y+2, m, r, g, b)
		plotFontLine("**  **", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'p':
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("-**--*", x, y+2, m, r, g, b)
		plotFontLine("****", x+1, y+3, m, r, g, b)
		plotFontLine("-**-", x, y+4, m, r, g, b)
		plotFontLine("****", x, y+5, m, r, g, b)
	case 'P':
		plotFontLine("*****-", x, y, m, r, g, b)
		plotFontLine("-**-**", x, y+1, m, r, g, b)
		plotFontLine("****-", x+1, y+2, m, r, g, b)
		plotFontLine("**-", x+1, y+3, m, r, g, b)
		plotFontLine("-**-", x, y+4, m, r, g, b)
		plotFontLine("****", x, y+5, m, r, g, b)
	case 'q':
		plotFontLine("-**-**", x, y+1, m, r, g, b)
		plotFontLine("*--**-", x, y+2, m, r, g, b)
		plotFontLine("-****", x, y+3, m, r, g, b)
		plotFontLine("-**-", x+2, y+4, m, r, g, b)
		plotFontLine("****", x+2, y+5, m, r, g, b)
	case 'Q':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("** -**", x, y+2, m, r, g, b)
		plotFontLine("**-***", x, y+3, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
		plotFontLine("-**", x+3, y+5, m, r, g, b)
	case 'r':
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("-*****", x, y+2, m, r, g, b)
		plotFontLine("**--*", x+1, y+3, m, r, g, b)
		plotFontLine("-**-", x, y+4, m, r, g, b)
		plotFontLine("****", x, y+5, m, r, g, b)
	case 'R':
		plotFontLine("****-", x, y, m, r, g, b)
		plotFontLine("-*--*-", x, y+1, m, r, g, b)
		plotFontLine("*--*-", x+1, y+2, m, r, g, b)
		plotFontLine("***-", x+1, y+3, m, r, g, b)
		plotFontLine("-*-**-", x, y+4, m, r, g, b)
		plotFontLine("**--**", x, y+5, m, r, g, b)
	case 's':
		plotFontLine("-*****", x, y+1, m, r, g, b)
		plotFontLine("**-", x, y+2, m, r, g, b)
		plotFontLine("-****-", x, y+3, m, r, g, b)
		plotFontLine("-**", x+3, y+4, m, r, g, b)
		plotFontLine("*****-", x, y+5, m, r, g, b)
	case 'S':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-***-", x, y+2, m, r, g, b)
		plotFontLine("-**-", x+2, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 't':
		plotFontLine("-*", x+1, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("****", x, y+2, m, r, g, b)
		plotFontLine("-**-", x, y+3, m, r, g, b)
		plotFontLine("**-*", x+1, y+4, m, r, g, b)
		plotFontLine("-**-", x+1, y+5, m, r, g, b)
	case 'T':
		plotFontLine("******", x, y, m, r, g, b)
		plotFontLine("*-**-*", x, y+1, m, r, g, b)
		plotFontLine("**", x+2, y+2, m, r, g, b)
		plotFontLine("**", x+2, y+3, m, r, g, b)
		plotFontLine("-**-", x+1, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case 'u':
		plotFontLine("** **", x, y+1, m, r, g, b)
		plotFontLine("** **", x, y+2, m, r, g, b)
		plotFontLine("** **", x, y+3, m, r, g, b)
		plotFontLine("**-**-", x, y+4, m, r, g, b)
		plotFontLine("-**-**", x, y+5, m, r, g, b)
	case 'U':
		plotFontLine("**  **", x, y, m, r, g, b)
		plotFontLine("**  **", x, y+1, m, r, g, b)
		plotFontLine("**  **", x, y+2, m, r, g, b)
		plotFontLine("**  **", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case 'v':
		plotFontLine("**  **", x, y+1, m, r, g, b)
		plotFontLine("**  **", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
		plotFontLine("-**-", x+1, y+5, m, r, g, b)
	case 'V':
		plotFontLine("**  **", x, y, m, r, g, b)
		plotFontLine("**  **", x, y+1, m, r, g, b)
		plotFontLine("**  **", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
		plotFontLine("-**-", x+1, y+5, m, r, g, b)
	case 'w':
		plotFontLine("*   *", x+1, y+1, m, r, g, b)
		plotFontLine("*- -*", x+1, y+2, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+3, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+4, m, r, g, b)
		plotFontLine("-*-*-", x+1, y+5, m, r, g, b)
	case 'W':
		plotFontLine("*   *", x+1, y, m, r, g, b)
		plotFontLine("*- -*", x+1, y+1, m, r, g, b)
		plotFontLine("*- -*", x+1, y+2, m, r, g, b)
		plotFontLine("*-*-*", x+1, y+3, m, r, g, b)
		plotFontLine("**-**", x+1, y+4, m, r, g, b)
		plotFontLine("-*-*-", x+1, y+5, m, r, g, b)
	case 'x':
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-****-", x, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
		plotFontLine("**--**", x, y+5, m, r, g, b)
	case 'X':
		plotFontLine("**--**", x, y, m, r, g, b)
		plotFontLine("-****-", x, y+1, m, r, g, b)
		plotFontLine("-**-", x+1, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
		plotFontLine("**--**", x, y+5, m, r, g, b)
	case 'y':
		plotFontLine("**  **", x, y+1, m, r, g, b)
		plotFontLine("**--**", x, y+2, m, r, g, b)
		plotFontLine("-****-", x, y+3, m, r, g, b)
		plotFontLine("-*", x+4, y+4, m, r, g, b)
		plotFontLine("*****-", x, y+5, m, r, g, b)
	case 'Y':
		plotFontLine("**  **", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-****-", x, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-**-", x+1, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case 'z':
		plotFontLine("******", x, y+1, m, r, g, b)
		plotFontLine("*--**-", x, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("******", x, y+5, m, r, g, b)
	case 'Z':
		plotFontLine("******", x, y, m, r, g, b)
		plotFontLine("*- -**", x, y+1, m, r, g, b)
		plotFontLine("-**-", x+2, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("******", x, y+5, m, r, g, b)
	case 'æ':
		plotFontLine("****-", x, y+1, m, r, g, b)
		plotFontLine("-*-*", x+1, y+2, m, r, g, b)
		plotFontLine("-***-", x, y+3, m, r, g, b)
		plotFontLine("*-*-", x, y+4, m, r, g, b)
		plotFontLine("-****", x, y+5, m, r, g, b)
	case 'Æ':
		plotFontLine("-*****", x, y, m, r, g, b)
		plotFontLine("**-**-", x, y+1, m, r, g, b)
		plotFontLine("******", x, y+2, m, r, g, b)
		plotFontLine("**-**-", x, y+3, m, r, g, b)
		plotFontLine("** **-", x, y+4, m, r, g, b)
		plotFontLine("** ***", x, y+5, m, r, g, b)
	case 'ø':
		plotFontLine("-***-*", x, y+1, m, r, g, b)
		plotFontLine("*--**-", x, y+2, m, r, g, b)
		plotFontLine("*-**-*", x, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("*-***-", x, y+5, m, r, g, b)
	case 'Ø':
		plotFontLine("-***-*", x, y, m, r, g, b)
		plotFontLine("*--**-", x, y+1, m, r, g, b)
		plotFontLine("*-**-*", x, y+2, m, r, g, b)
		plotFontLine("*-**-*", x, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("*-***-", x, y+5, m, r, g, b)
	case 'å':
		plotFontLine("-**-", x+1, y, m, r, g, b)
		plotFontLine("***-", x+1, y+1, m, r, g, b)
		plotFontLine("--**", x+1, y+2, m, r, g, b)
		plotFontLine("-****", x, y+3, m, r, g, b)
		plotFontLine("**-**-", x, y+4, m, r, g, b)
		plotFontLine("-**-**", x, y+5, m, r, g, b)
	case 'Å':
		plotFontLine("**", x+2, y, m, r, g, b)
		plotFontLine("--", x+2, y+1, m, r, g, b)
		plotFontLine("-****-", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("******", x, y+4, m, r, g, b)
		plotFontLine("**--**", x, y+5, m, r, g, b)
	case '.':
		plotFontLine("**-", x+1, y+4, m, r, g, b)
		plotFontLine("-**", x, y+5, m, r, g, b)
	case ':':
		plotFontLine("**-", x+1, y, m, r, g, b)
		plotFontLine("-**", x, y+1, m, r, g, b)
		plotFontLine("**-", x+1, y+4, m, r, g, b)
		plotFontLine("-**", x, y+5, m, r, g, b)
	case ';':
		plotFontLine("**-", x+1, y, m, r, g, b)
		plotFontLine("-**", x, y+1, m, r, g, b)
		plotFontLine("-*", x+1, y+4, m, r, g, b)
		plotFontLine("*-", x+1, y+5, m, r, g, b)
	case ',':
		plotFontLine("-*", x+1, y+4, m, r, g, b)
		plotFontLine("*-", x+1, y+5, m, r, g, b)
	case '\'':
		plotFontLine("-*", x+1, y, m, r, g, b)
		plotFontLine("*-", x+1, y+1, m, r, g, b)
	case '"':
		plotFontLine("-* -*", x, y, m, r, g, b)
		plotFontLine("*- *-", x, y+1, m, r, g, b)
	case '*':
		plotFontLine("-*-", x, y, m, r, g, b)
		plotFontLine("-*-", x+1, y+1, m, r, g, b)
		plotFontLine("-******", x-3, y+2, m, r, g, b)
		plotFontLine("-*-", x+1, y+3, m, r, g, b)
		plotFontLine("-*-", x, y+4, m, r, g, b)
	case '+':
		plotFontLine("  **   ", x, y+2, m, r, g, b)
		plotFontLine("-****-", x, y+3, m, r, g, b)
		plotFontLine("  **   ", x, y+4, m, r, g, b)
	case '!':
		plotFontLine("**", x+1, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("-**-", x, y+2, m, r, g, b)
		plotFontLine("**", x+1, y+3, m, r, g, b)
		plotFontLine("--", x+1, y+4, m, r, g, b)
		plotFontLine("**", x+1, y+5, m, r, g, b)
	case '?':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-**-", x+2, y+2, m, r, g, b)
		plotFontLine("**-", x+2, y+3, m, r, g, b)
		plotFontLine("--", x+2, y+4, m, r, g, b)
		plotFontLine("**", x+2, y+5, m, r, g, b)
	case '-':
		plotFontLine("-****-", x, y+3, m, r, g, b)
	case '=':
		plotFontLine("-****-", x, y+2, m, r, g, b)
		plotFontLine("-****-", x, y+4, m, r, g, b)
	case '_':
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '/':
		plotFontLine("*", x+4, y+1, m, r, g, b)
		plotFontLine("*-", x+3, y+2, m, r, g, b)
		plotFontLine("*-", x+2, y+3, m, r, g, b)
		plotFontLine("*-", x+1, y+4, m, r, g, b)
		plotFontLine("*-", x+0, y+5, m, r, g, b)
	case '1':
		plotFontLine("-**", x+1, y, m, r, g, b)
		plotFontLine("***", x+1, y+1, m, r, g, b)
		plotFontLine("-**", x+1, y+2, m, r, g, b)
		plotFontLine("**", x+2, y+3, m, r, g, b)
		plotFontLine("**-", x+2, y+4, m, r, g, b)
		plotFontLine("****", x+1, y+5, m, r, g, b)
	case '2':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("*- -**", x, y+1, m, r, g, b)
		plotFontLine("-**-", x+2, y+2, m, r, g, b)
		plotFontLine("-**-", x+1, y+3, m, r, g, b)
		plotFontLine("-**--*", x, y+4, m, r, g, b)
		plotFontLine("******", x, y+5, m, r, g, b)
	case '3':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("**-", x+3, y+2, m, r, g, b)
		plotFontLine("-**", x+3, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '4':
		plotFontLine("-***", x+1, y, m, r, g, b)
		plotFontLine("-*-**", x, y+1, m, r, g, b)
		plotFontLine("*--**-", x, y+2, m, r, g, b)
		plotFontLine("******", x, y+3, m, r, g, b)
		plotFontLine("-**-", x+2, y+4, m, r, g, b)
		plotFontLine("****", x+2, y+5, m, r, g, b)
	case '5':
		plotFontLine("******", x, y, m, r, g, b)
		plotFontLine("**-", x, y+1, m, r, g, b)
		plotFontLine("*****-", x, y+2, m, r, g, b)
		plotFontLine("-**", x+3, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '6':
		plotFontLine("-***", x+1, y, m, r, g, b)
		plotFontLine("-**-", x, y+1, m, r, g, b)
		plotFontLine("**-", x, y+2, m, r, g, b)
		plotFontLine("*****-", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '7':
		plotFontLine("******", x, y, m, r, g, b)
		plotFontLine("*- -**", x, y+1, m, r, g, b)
		plotFontLine("-**", x+3, y+2, m, r, g, b)
		plotFontLine("-**-", x+2, y+3, m, r, g, b)
		plotFontLine("**-", x+2, y+4, m, r, g, b)
		plotFontLine("**", x+2, y+5, m, r, g, b)
	case '8':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-****-", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '9':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("-*****", x, y+2, m, r, g, b)
		plotFontLine("-**", x+3, y+3, m, r, g, b)
		plotFontLine("-**-", x+2, y+4, m, r, g, b)
		plotFontLine("-***-", x, y+5, m, r, g, b)
	case '0':
		plotFontLine("-****-", x, y, m, r, g, b)
		plotFontLine("**--**", x, y+1, m, r, g, b)
		plotFontLine("**--**", x, y+2, m, r, g, b)
		plotFontLine("**--**", x, y+3, m, r, g, b)
		plotFontLine("**--**", x, y+4, m, r, g, b)
		plotFontLine("-****-", x, y+5, m, r, g, b)
	case '(':
		plotFontLine("-***", x+2, y, m, r, g, b)
		plotFontLine("**-", x+2, y+1, m, r, g, b)
		plotFontLine("**-", x+1, y+2, m, r, g, b)
		plotFontLine("**-", x+1, y+3, m, r, g, b)
		plotFontLine("**-", x+2, y+4, m, r, g, b)
		plotFontLine("-***", x+2, y+5, m, r, g, b)
	case ')':
		plotFontLine("***-", x, y, m, r, g, b)
		plotFontLine("-**", x+1, y+1, m, r, g, b)
		plotFontLine("-**", x+2, y+2, m, r, g, b)
		plotFontLine("-**", x+2, y+3, m, r, g, b)
		plotFontLine("-**", x+1, y+4, m, r, g, b)
		plotFontLine("***-", x, y+5, m, r, g, b)
	case 0:
		return errors.New("the rune was 0. Did you pass a coordinate instead of a rune?")
	default:
		return fmt.Errorf("rune %s is not available", string(l))
	}
	return nil
}
