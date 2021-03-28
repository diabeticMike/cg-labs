package main

import (
	"fmt"
	"math"

	"github.com/fogleman/gg"
)

const (
	scale    = 30
	filename = "image.png"
)

var files []string

func line(x1, y1, x2, y2, c float64) {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SetLineWidth(2)
	dc.SetRGB(226, 106, 106)
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
		tmp = y1
		y1 = y2
		y2 = tmp
	}
	dy := 1
	if y1 > y2 {
		dy = -1
	}
	m := math.Abs(y2 - y1)
	n := x2 - x1
	S := m - n
	l := n
	if m > n {
		l = m
	}
	m += m
	n += n
	dc.DrawRectangle(x1, y1, x1+c, y1+c)
	for i := 1; i < int(l); i++ {
		if S <= 0 {
			S += m
			x1 += 1
		} else {
			S -= n
			y1 += float64(dy)
		}
		dc.DrawRectangle(x1, y1, x1+c, y1+c)
	}
	dc.Fill()
	dc.SavePNG("line.png")
}

func main() {
	// drawParab()
	// drawHiperb()
	// ax3(5)
	// autoRectangle(500, 500, 100, 200)
	// drawElipse()
	// drawElipseLine(0, 90)
	// drawCircle()
	line(200, 200, 200, 400, 100)
}

func drawParab() {
	c := 11
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for i := -5.0; i < float64(c)/2; i++ {
		x = append(x, i)
		y = append(y, i*i)
	}

	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(3)
	for i := 0; i < c-1; i++ {
		dc.DrawLine(x[i]*scale+float64(w/2), float64(h/2)-y[i]*scale, x[i+1]*scale+float64(w/2), float64(h/2)-y[i+1]*scale)
		dc.Stroke()
	}
	dc.SavePNG("parab.png")
}

func drawHiperb() {
	c := 11
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for i := 1.0; i < float64(c)+1; i++ {
		x = append(x, i)
		y = append(y, 1/i)
	}

	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(3)
	for i := 0; i < c-1; i++ {
		dc.DrawLine(float64(w/2)+x[i]*scale*2, float64(h/2)+y[i]*scale*2, float64(w/2)+x[i+1]*scale*2, float64(h/2)+y[i+1]*scale*2)
		dc.Stroke()
	}
	dc.SavePNG("hiperb.png")
}

func ax3(a float64) {
	c := 11
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for i := 1.0; i < float64(c)+1; i++ {
		x = append(x, i)
		y = append(y, a*i*3)
		fmt.Println(i, a*3*i)
	}

	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(3)
	for i := 0; i < c-1; i++ {
		dc.DrawLine(float64(w/2)+x[i]*5, float64(h/2)+y[i]*5, float64(w/2)+x[i+1]*5, float64(h/2)+y[i+1]*5)
		dc.Stroke()
	}
	dc.SavePNG("ax3.png")
}

func autoRectangle(x1, y1, a, b float64) {
	x2, y2 := x1+a, y1
	x3, y3 := x1, y1+b
	x4, y4 := x1+a, y1+b

	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(3)
	dc.DrawLine(x1, y1, x2, y2)
	dc.DrawLine(x1, y1, x3, y3)
	dc.DrawLine(x4, y4, x2, y2)
	dc.DrawLine(x4, y4, x3, y3)
	dc.Stroke()

	dc.SavePNG("rect.png")
}

var w, h int = 2543, 1344

func drawElipse() {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	// var a float64 = 320
	var dx, dy float64 = float64(w / 2), float64(h / 2)
	c := 360
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for d := 0; d < c; d++ {
		rad := toRad(d)
		x = append(x, 320*math.Sin(rad)+dx)
		y = append(y, 160*math.Cos(rad)+dy)
	}
	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(2)
	dc.DrawLine(x[0], y[0], x[len(x)-1], y[len(y)-1])
	for i := 0; i < c-1; i++ {
		dc.DrawLine(x[i], y[i], x[i+1], y[i+1])
		dc.Stroke()
	}
	dc.SavePNG("elipse.png")
}

func drawElipseLine(s, e float64) {
	s += 90
	e += 90
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	// var a float64 = 320
	var dx, dy float64 = float64(w / 2), float64(h / 2)
	c := int(e - s)
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for d := 0; d < c; d++ {
		if s == e {
			break
		}
		rad := toRad(int(s))
		x = append(x, 320*math.Sin(rad)+dx)
		y = append(y, 160*math.Cos(rad)+dy)
		s++
	}
	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(2)
	for i := 0; i < c-1; i++ {
		dc.DrawLine(x[i], y[i], x[i+1], y[i+1])
		dc.Stroke()
	}
	dc.SavePNG("elipseLine.png")
}

func drawCircle() {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	// var a float64 = 320
	var dx, dy float64 = float64(w / 2), float64(h / 2)
	c := 360
	x := make([]float64, 0, c)
	y := make([]float64, 0, c)
	for d := 0; d < c; d++ {
		rad := toRad(d)
		x = append(x, 160*math.Sin(rad)+dx)
		y = append(y, 160*math.Cos(rad)+dy)
	}
	dc.SetRGB(226, 106, 106)
	dc.SetLineWidth(2)
	dc.DrawLine(x[0], y[0], x[len(x)-1], y[len(y)-1])
	for i := 0; i < c-1; i++ {
		dc.DrawLine(x[i], y[i], x[i+1], y[i+1])
		dc.Stroke()
	}
	dc.SavePNG("circle.png")
}

func toRad(deg int) float64 {
	return float64(deg) * (math.Pi / 180.0)
}
