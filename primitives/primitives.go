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
// Since we no longer pass a window we can remove a dependency from here 
//	"golang.org/x/exp/shiny/screen"
	"image"
	// To skip imported but not used error, fmt is handy 
	// For what I refer to as norweigan debugging
	_"fmt")

func swap(a,b int) (int, int) {
	return b,a
}

var dst *image.RGBA
var fg color.RGBA

// Moving this here makes the code to appear less cluttered I believe 
func SetForeground(color color.RGBA){
	fg=color
}

// set the buffer / image that we are working on
func SetDrawable(drawable *image.RGBA){
	dst = drawable
}

func HorLine(x0,y0,x1 int) {
	if x0 > x1 {
		x0,x1 = swap(x0,x1)
	}
	for x:=x0; x<=x1; x++ {
		dst.SetRGBA(x,y0,fg)
	}
}

func VerLine(x0,y0,y1 int) {
	if y0 > y1 {
		y0,y1 = swap(y0,y1)
	}
	for y:=y0; y<=y1; y++ {
		dst.SetRGBA(x0,y,fg)
	}
}

// NOTE:  Makes more sense to have this here. 
// Although it's only semi primitive 
// Also added a nice offset and the ability to have different step in the 
// horizontal and vertical direction
func DrawGrid(rect image.Rectangle, offset image.Point,step image.Point){
	for y:=rect.Min.Y+offset.Y; y<=rect.Max.Y; y+=step.Y {
		HorLine(rect.Min.X,y,rect.Max.X)
	}

	for x:=rect.Min.X+offset.X; x<=rect.Max.X; x+=step.X {
		VerLine(x,rect.Min.Y,rect.Max.Y)
	}
}
