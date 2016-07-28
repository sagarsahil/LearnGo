//This program creates a gif image on the fly.
// use the program by redirecting the output to a file with extension gif.
//.createimg > imagename.gif

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	whiteIndex = 0 //Color 1
	blackIndex = 1 //Color 2
)

func main() {
	createimage(os.Stdout)
}

func createimage(out io.Writer) {
	const (
		cycles = 5     // number of ossillator
		res    = 0.001 // angular resolution
		size   = 100   // image canvas covers
		nframe = 64    // number of animation frame
		delay  = 8     // delay between frame in 10 ms unites
	)

	freq := rand.Float64() * 3.0 //  y oscillator frequency
	anim := gif.GIF{LoopCount: nframe}
	phase := 0.0
	for counter := 0; counter < nframe; counter++ {
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}
