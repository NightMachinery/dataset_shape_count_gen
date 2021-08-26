package main

import "github.com/fogleman/gg"

func main() {
	width := 200
	height := width
    dc := gg.NewContext(width, height)

	dc.DrawRectangle(0., 0., float64(width), float64(height))
    dc.SetRGB(1, 1, 1) // white
    dc.Fill()

	circle_radius := float64(5)
    dc.DrawCircle(float64(width/2), float64(height/2), circle_radius)
    dc.SetRGB(0, 0, 0)
    dc.Fill()

    dc.SavePNG("out.png")
}
