package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

const (
	scale    = 3
	filename = "image.png"
)

func main() {
	// draw()
	drawDiagonal()
	drawAbsOrd()
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
	s := float64(rand.Intn(4))
	var x, y float64 = float64(w / 10), float64(h / 10)
	for i := 0; i < 10; i++ {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, y, s, visible, invisible)
		x += float64(w / 10)
		y += float64(h / 10)
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
	s := float64(rand.Intn(4))
	var x float64 = float64(w / 10)
	for i := 0; i < 9; i++ {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, 500, s, visible, invisible)
		x += float64(w / 10)
		dc.SavePNG(fmt.Sprintf("4image%v.png", i))
	}
	var y float64 = float64(h / 10)
	for i := 9; i < 19; i++ {
		dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		drawPrism(dc, x, y, s, visible, invisible)
		y += float64(h / 10)
		dc.SavePNG(fmt.Sprintf("4image%v.png", i))
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
