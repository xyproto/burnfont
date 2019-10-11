# burnfont [![GoDoc](https://godoc.org/github.com/xyproto/burnfont?status.svg)](http://godoc.org/github.com/xyproto/burnfont) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/burnfont)](https://goreportcard.com/report/github.com/xyproto/burnfont)

This is a hand-crafted 6x6 pixel font, defined by code, created in 1997.

Here is a generated image with all available letters in this font:

![letters](img/letters.png)

Here is the same image, but scaled up 4x:

![scaled](img/scaled.png)

The font is a bit nostalgic to me, and was used in my small DOS drawing program named *Burn*. The palette that came with *Burn* is [available here](https://github.com/xyproto/burnpalette). The GUI was drawn with 320x200 pixels, 256 indexed colors (`mode 13h`), and looked like this:

![burn screenshot](img/burn.png)


The font definition looks like this:

```go
case 'k':
	fontLine("***", x, y)
	fontLine("-**", x, y+1)
	fontLine("**-**", x+1, y+2)
	fontLine("****-", x+1, y+3)
	fontLine("-**-**", x, y+4)
	fontLine("*** **", x, y+5)
```

`*` is an opaque pixel, while `-` is a 25% transparent one.

## Contents

This Go package has a slice `Available` that lists all available runes. There is also a `Draw` function that takes an `*image.RGBA` value, a rune, a position (x and y) and a color (r, g and b) and draws a letter at that position in the image.

* The `cmd/scaled` example outputs an image where each "pixel" of the font is 4x4 pixels.
* The `cmd/letters` example outputs an image where each pixel is a pixel.

## Generating an image

The `scaled` utility can be built and run like this:

    cd cmd/scaled
    go build
    ./scaled > letters.png

## General information

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
