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
	"goswan/primitives"
	"image"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

// 640x480 for that retro feel
// FIXME: make this configurable without having to recompile
var width int = 640
var height int = 480

var b screen.Buffer
var t screen.Texture
var err error


func paintevent(s screen.Screen, w screen.Window,bounds image.Rectangle) {
	w.Fill(bounds,colorscheme.Black,screen.Src)

	primitives.SetForeground(colorscheme.White)

	primitives.DrawGrid(bounds.Inset(60),image.Point{3,3},image.Point{8,6})
//	drawGrid(w,bounds,36)
//	primitives.HorLine(w,100,10,200,colorscheme.White)
//	prim_2d.SetForeground(colorscheme.White)
//	prim_2d.HorLine(100,10,200, colorscheme.White)

	// 3. Create a texture 

	// 4. Put the buffer in the texture 
	t.Upload(image.Point{0,0},b,bounds)

	// 5. Copy the texture to the window
	// screen.Over => reference
	// What is the last argument?
	w.Copy(image.Point{0,0},t,bounds,screen.Over,nil)
}


// Re run on resize or whatever  
func setupDrawing(s screen.Screen){
	winsize:=image.Point{width,height};

	// Create a buffer, on this we can use image operations 
	b,err=s.NewBuffer(winsize);
	if err != nil {
		// FIXME: handle error
	}

	// give our primitives package a pointer to the buffer image
	primitives.SetDrawable(b.RGBA())

	// Create the texture, this is not necesserily accessible by the CPU  
	t,err=s.NewTexture(winsize)
	if err!= nil {
		// FIXME handle errors 
	}
}

// Don't forget to cleanup
func cleanupDrawing(){
	b.Release();
	t.Release();
}

func main() {
	colorscheme.Whatever()

	driver.Main(func(s screen.Screen) {
		// TODO wouldn't it be fun to patch golang.org/x/exp/shiny/screen
		// so that it takes other arguments fullscreen, title, icon, cursorHidden
		opts := screen.NewWindowOptions{width,height}
		w, err := s.NewWindow(&opts)
		if err != nil {
			return
		}
		defer w.Release()

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
