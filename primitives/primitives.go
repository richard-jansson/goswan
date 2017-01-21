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
 *
*/

// Provides for basic 2d primitive functions 
package primitives

import (
	"image/color"
	"golang.org/x/exp/shiny/screen"
	"image"
	// To skip imported but not used error, fmt is handy 
	// For what I refer to as norweigan debugging
	_"fmt")


func swap(a,b int) (int, int) {
	return b,a
}

func putPixel(w screen.Window,x int, y int,c color.RGBA) {
	w.Fill(image.Rect(x,y,x+1,y+1),c,screen.Src)
}

func HorLine(w screen.Window, x0,y0,x1 int,fg color.RGBA) {
	if x0 > x1 {
		x0,x1 = swap(x0,x1)
	}
	for x:=x0; x<=x1; x++ {
		putPixel(w,x,y0,fg)
	}
}

func VerLine(w screen.Window, x0,y0,y1 int,fg color.RGBA) {
	if y0 > y1 {
		y0,y1 = swap(y0,y1)
	}
	for y:=y0; y<=y1; y++ {
		putPixel(w,x0,y,fg)
	}
}
