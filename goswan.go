/*
 *    Copyright (C) 2016 Richard Jansson
 *
 *    This file is part of goswan.
 *
 *    goswan is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    goswan is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with goswan.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"goswan/colorscheme"
	"goswan/twodimgraphics"
	"image"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

var width int = 1280
var height int = 960 

var b screen.Buffer
var t screen.Texture
var err error


func paintevent(s screen.Screen, w screen.Window,bounds image.Rectangle) {
	w.Fill(bounds,colorscheme.Black,screen.Src)

	twodimgraphics.SetForeground(colorscheme.Grey)

	size := image.Rectangle{ image.Point{0,0},image.Point{bounds.Max.X/2,bounds.Max.Y/2}}

	// For any newbies out there this is something that you ought to understand! 
	for i:=0;i<3;i++ {
		var pos image.Rectangle;

		pos.Min.X = (i%2)*size.Max.X
		pos.Min.Y = (i/2)*size.Max.Y
		pos.Max.X = pos.Min.X + size.Max.X
		pos.Max.Y = pos.Min.Y + size.Max.Y

		twodimgraphics.DrawGrid(pos.Inset(60),image.Point{16,12},// offset 
					image.Point{32,24}, // step 
					image.Point{7,7} ) // origo 
	}


	t.Upload(image.Point{0,0},b,bounds)

	w.Copy(image.Point{0,0},t,bounds,screen.Over,nil)
}

func setupDrawing(s screen.Screen){
	winsize:=image.Point{width,height};

	b,err=s.NewBuffer(winsize);
	if err != nil {
		// FIXME: handle error
	}

	twodimgraphics.SetDrawable(b.RGBA())

	t,err=s.NewTexture(winsize)
	if err!= nil {
		// FIXME handle errors 
	}
}

func cleanupDrawing(){
	b.Release();
	t.Release();
}

func main() {
	colorscheme.Whatever()

	driver.Main(func(s screen.Screen) {
// FIXME: patch shiny so that we can set title and icon 
		opts := screen.NewWindowOptions{width,height}
		w, err := s.NewWindow(&opts)
		if err != nil {
			return
		}
		defer w.Release()

		twodimgraphics.Setup()
		defer twodimgraphics.Cleanup()

		setupDrawing(s)
		defer cleanupDrawing()

		var sz size.Event
		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				paintevent(s, w,sz.Bounds())
			case size.Event:
				sz=e
			}
		}
	})
}
