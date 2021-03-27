package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/andybons/gogif"

	"github.com/fogleman/gg"
)

const (
	scale    = 3
	filename = "image.png"
)

var files []string

func main() {
	// draw()
	// drawDiagonal()
	// files = make([]string, 0)
	// drawAbsOrd()
	// {
	// 	outGif := &gif.GIF{}
	// 	for _, name := range files {
	// 		f, _ := os.Open(name)
	// 		inGif, _ := png.Decode(f)
	// 		bounds := inGif.Bounds()
	// 		palettedImage := image.NewPaletted(bounds, nil)
	// 		quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	// 		quantizer.Quantize(palettedImage, bounds, inGif, image.ZP)
	// 		f.Close()
	
	// 		outGif.Image = append(outGif.Image, palettedImage)
	// 		outGif.Delay = append(outGif.Delay, 1)
	// 	}
	// 	// save to out.gif
	// 	f, _ := os.OpenFile("out1.gif", os.O_WRONLY|os.O_CREATE, 0600)
	// 	defer f.Close()
	// 	gif.EncodeAll(f, outGif)
	// }
	//
	// files = make([]string, 0)
	// drawDiagonal()
	// {
	// 	outGif := &gif.GIF{}
	// 	for _, name := range files {
	// 		f, _ := os.Open(name)
	// 		inGif, _ := png.Decode(f)
	// 		bounds := inGif.Bounds()
	// 		palettedImage := image.NewPaletted(bounds, nil)
	// 		quantizer := gogif.MedianCutQuantizer{NumColor: 64}
	// 		quantizer.Quantize(palettedImage, bounds, inGif, image.ZP)
	// 		f.Close()
	
	// 		outGif.Image = append(outGif.Image, palettedImage)
	// 		outGif.Delay = append(outGif.Delay, 1)
	// 	}
	// 	// save to out.gif
	// 	f, _ := os.OpenFile("out2.gif", os.O_WRONLY|os.O_CREATE, 0600)
	// 	defer f.Close()
	// 	gif.EncodeAll(f, outGif)
	// }

	files = make([]string, 18)
	drawElipse()
	{
		outGif := &gif.GIF{}
		for _, name := range files {
			f, _ := os.Open(name)
			inGif, _ := png.Decode(f)
			bounds := inGif.Bounds()
			palettedImage := image.NewPaletted(bounds, nil)
			quantizer := gogif.MedianCutQuantizer{NumColor: 64}
			quantizer.Quantize(palettedImage, bounds, inGif, image.ZP)
			f.Close()

			outGif.Image = append(outGif.Image, palettedImage)
			outGif.Delay = append(outGif.Delay, 1)
		}
		// save to out.gif
		f, _ := os.OpenFile("out3.gif", os.O_WRONLY|os.O_CREATE, 0600)
		defer f.Close()
		gif.EncodeAll(f, outGif)
	}
}

var w, h int = 2543, 1344

func draw() {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	drawDiagonal()
	dc.SavePNG(filename)
}

func drawDiagonal() {
	dc := gg.NewContext(w+100, h+100)
	rand.Seed(time.Now().UnixNano())
	visible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}
	invisible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}
	var x, y float64 = float64(w / 10), float64(h)
	for i := 40; i > 0; i-- {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, y, 0.1*float64(i), visible, invisible)
		x += float64(w / 40)
		y -= float64(h / 40)
		files = append(files, fmt.Sprintf("5image%v.png", i))
		dc.SavePNG(fmt.Sprintf("5image%v.png", i))
	}
}

func drawAbsOrd() {
	dc := gg.NewContext(w+100, h+100)
	rand.Seed(time.Now().UnixNano())
	visible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}
	invisible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}
	var x float64 = float64(w / 50)
	for i := 20; i > 0; i-- {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, 500, 2, visible, invisible)
		x += float64(w / 50)
		files = append(files, fmt.Sprintf("4image%v.gif", i))
		dc.SavePNG(fmt.Sprintf("4image%v.gif", i))
	}
	var y float64 = 500
	for i := 0; i < 20; i++ {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, y, 2, visible, invisible)
		y += float64(h / 50)
		files = append(files, fmt.Sprintf("4image%v.gif", 40+i))
		dc.SavePNG(fmt.Sprintf("4image%v.gif", 40+i))
	}
}

func drawRandom50(dc *gg.Context) {
	for i := 0; i < 50; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float64(rand.Intn(w - scale*300))
		y := float64(rand.Intn(h - scale*200))
		visible := Color{float64(rand.Intn(255)),
			float64(rand.Intn(255)),
			float64(rand.Intn(255))}
		invisible := Color{float64(rand.Intn(255)),
			float64(rand.Intn(255)),
			float64(rand.Intn(255))}
		s := float64(rand.Intn(4))
		drawPrism(dc, 200+x, 450+y, s, visible, invisible)
	}
}

func drawPrismWithChild(dc *gg.Context) {
	rand.Seed(time.Now().UnixNano())
	x := float64(rand.Intn(w - scale*300))
	y := float64(rand.Intn(h - scale*300))
	drawPrism(dc, 200+x, 600+y, scale+1, Color{226, 106, 106}, Color{100, 100, 100})
	drawPrism(dc, 200+x, 600+y, scale, Color{226, 106, 106}, Color{100, 100, 100})
}

func drawPrism(dc *gg.Context, x, y, s float64, visible, invisible Color) {
	dc.SetRGBA(visible.r, visible.g, visible.b, 1)
	o := origin{x, y}
	// dc.DrawPoint(o.x, o.y, 1)
	dc.SetLineWidth(6)
	// draw top and bottom lines
	dc.DrawLine(o.x-50*s, o.y-100*s, o.x+50*s, o.y-100*s)
	dc.DrawLine(o.x-50*s, o.y+30*s, o.x+50*s, o.y+30*s)

	// draw left and right
	dc.DrawLine(o.x-50*s, o.y-100*s, o.x-50*s, o.y+30*s)
	dc.DrawLine(o.x+50*s, o.y-100*s, o.x+50*s, o.y+30*s)

	// draw top and bottom bases
	dc.DrawLine(o.x-50*s, o.y-100*s, o.x-80*s, o.y-130*s)
	dc.DrawLine(o.x+50*s, o.y-100*s, o.x+80*s, o.y-130*s)
	dc.DrawLine(o.x-80*s, o.y-130*s, o.x, o.y-160*s)
	dc.DrawLine(o.x+80*s, o.y-130*s, o.x, o.y-160*s)

	dc.DrawLine(o.x-50*s, o.y+30*s, o.x-80*s, o.y)
	dc.DrawLine(o.x+50*s, o.y+30*s, o.x+80*s, o.y)
	dc.DrawLine(o.x-80*s, o.y, o.x, o.y-30*s)
	dc.DrawLine(o.x+80*s, o.y, o.x, o.y-30*s)

	// draw heights
	dc.DrawLine(o.x-80*s, o.y, o.x-80*s, o.y-130*s)
	dc.DrawLine(o.x+80*s, o.y, o.x+80*s, o.y-130*s)
	dc.DrawLine(o.x, o.y-30*s, o.x, o.y-160*s)
	dc.Stroke()

	// redraw invisible lines
	dc.SetRGBA(invisible.r, invisible.g, invisible.b, 1)
	dc.DrawLine(o.x, o.y-30*s, o.x, o.y-160*s)
	dc.DrawLine(o.x-80*s, o.y, o.x, o.y-30*s)
	dc.DrawLine(o.x+80*s, o.y, o.x, o.y-30*s)

	dc.Stroke()
}

type origin struct {
	x, y float64
}

type Color struct {
	r, g, b float64
}

func drawElipse() {
	ch := make(chan int)

	var a, b float64 = 320, 160
	var dx, dy float64 = float64(w/2), float64(h/2)
	scale := 2.0
	//for i := 0; i < 3; i++ {
	for d := 0; d < 360; d += 20 {
		go func(s float64, deg int) {
			dc := gg.NewContext(w+100, h+100)
			println(deg)
			rad := toRad(deg)
			x := a*math.Sin(rad) + dx
			y := b*math.Cos(rad) + dy
			dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
			dc.SetRGB(0, 0, 0)
			dc.Fill()
			drawPrism(dc, x, y, s, Color{226, 106, 106}, Color{100, 100, 100})
			dc.SavePNG(fmt.Sprintf("6image%v.png", deg))
			ch <- deg
		}(scale, d)

		if 0 < d && d < 180 {
			scale -= 0.15
		} else {
			scale += 0.15
		}
	}
	for i := 0; i < 18; i++ {
		d := <-ch
		files[d/20] = fmt.Sprintf("6image%v.png", d)
	}
	//}
}

func toRad(deg int) float64 {
	return float64(deg) * (math.Pi / 180.0)
}
