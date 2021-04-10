package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

const (
	scale    = 3
	filename = "draw.png"
)

func main() {
	draw()
}

var w, h int = 1440, 720

type pair struct {
	x1, x2 int
}

func fill() map[int]pair {
	var point color.Color
	point = color.RGBA{170, 124, 134, 255}
	p := make(map[int]pair)
	for y := 0; y < h; y++ {
		x := 0
		_x := w
		for x != w/2 {
			if img.At(x, y) == point && img.At(_x, y) == point {
				p[y] = pair{x + 2, _x - 2}
				break
			}
			x++
			_x--
		}
	}

	return p
}

var (
	img image.Image
	col color.Color
)

func draw() {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	visible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}
	invisible := Color{float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255))}

	drawPrism(dc, float64(w)/2, float64(h)/2+200, 3, visible, invisible)

	col = color.RGBA{0, 0, 0, 255}
	f, err := os.Open("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}

	pairs := fill()
	for k, v := range pairs {
		dc.DrawLine(float64(v.x1), float64(k), float64(v.x2), float64(k))
	}
	dc.Stroke()

	dc.SavePNG(filename)
	// png.Encode(f, img)
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
		dc.SavePNG(fmt.Sprintf("image%v.png", i))
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
	o := origin{x, y}
	// dc.DrawPoint(o.x, o.y, 1)
	dc.SetLineWidth(6)

	// draw invisible lines
	dc.SetRGBA(invisible.r, invisible.g, invisible.b, 1)
	dc.DrawLine(o.x, o.y-30*s, o.x, o.y-160*s)
	dc.DrawLine(o.x-80*s, o.y, o.x, o.y-30*s)
	dc.DrawLine(o.x+80*s, o.y, o.x, o.y-30*s)
	dc.Stroke()

	dc.SetRGBA(visible.r, visible.g, visible.b, 1)
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

	// draw heights
	dc.DrawLine(o.x-80*s, o.y, o.x-80*s, o.y-130*s)
	dc.DrawLine(o.x+80*s, o.y, o.x+80*s, o.y-130*s)
	dc.Stroke()
}

func toRad(deg float64) float64 {
	return float64(deg) * (math.Pi / 180.0)
}

type origin struct {
	x, y float64
}

type Color struct {
	r, g, b float64
}
