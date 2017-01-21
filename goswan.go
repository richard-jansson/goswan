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

	"fmt"
)

func drawGrid(w screen.Window,bounds image.Rectangle, step int) {
	bounds=bounds.Inset(60)

	fmt.Printf("%s\n",bounds)

	for y:=bounds.Min.Y; y<=bounds.Max.Y; y+=step {
		primitives.HorLine(w,bounds.Min.X,y,bounds.Max.X,colorscheme.White)
	}

	for x:=bounds.Min.X; x<=bounds.Max.X; x+=step {
		primitives.VerLine(w,x,bounds.Min.Y,bounds.Max.Y,colorscheme.White)
	}
}

func paintevent(s screen.Screen, w screen.Window,bounds image.Rectangle) {
	w.Fill(bounds,colorscheme.Black,screen.Src)

	drawGrid(w,bounds,36)
//	primitives.HorLine(w,100,10,200,colorscheme.White)
//	prim_2d.SetForeground(colorscheme.White)
//	prim_2d.HorLine(100,10,200, colorscheme.White)
}

func main() {
	colorscheme.Whatever()

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			return
		}
		defer w.Release()

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
