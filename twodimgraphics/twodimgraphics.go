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

// Apparently it's not possible to name a package with a leading 2 
package twodimgraphics

import (
	// Sweet!
	"golang.org/x/image/font/gofont/goregular"
	"github.com/golang/freetype/truetype"
	// what is this and why do I need it? 
	"golang.org/x/image/math/fixed"
	"golang.org/x/image/font"
	"image/color"
	"image"
	"fmt")

var dst *image.RGBA
var fg color.RGBA
var cfont *truetype.Font
var err error;

func Setup(){
	cfont,err = truetype.Parse(goregular.TTF)
}

func Cleanup(){
	// FIXME is the font cleaned up 
}

func swap(a,b int) (int, int) {
	return b,a
}



func SetForeground(color color.RGBA){
	fg=color
}

// FIXME drawable - buffer - image, the change of names makes this a bit confusing
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

/* Minor complaint in C you could've written this one a bit shorter. 3 lines shorter 
void DrawGrid(...){
	for(int y=rect.Min.Y,int col=0; y<rect.Max.x; y+=step.Y, col++){
		HorLine(rect.Min.X,y,rect.Max.X)
		DrawString(rect.Min.X,y,strconv.Itoa(col))
	}
	for(int x=rect.Min.x,int row=0; y<rect.Min.x; x+=step.X, row++){
		VerLine(x,rect.Min.Y,rect.Max.Y)
		DrawString(x,rect.Min.Y,strconv.Itoa(row))
	}
}
*/

// FIXME: reverse the order of the parameters origo is more important than the 2nd argument offset
func DrawGrid(rect image.Rectangle, offset image.Point,step image.Point,origo image.Point){
	col,row := 0, 0
	format:="%03d"
	for y:=rect.Min.Y+offset.Y; y<=rect.Max.Y; y+=step.Y  {
		// FIXME implement line width and save 8 lines of code
		if (row-origo.Y) == 0 {
			HorLine(rect.Min.X,y-1,rect.Max.X)
			HorLine(rect.Min.X,y+1,rect.Max.X)
		}
		HorLine(rect.Min.X,y,rect.Max.X)
		DrawString(rect.Min.X,y,fmt.Sprintf(format,row-origo.Y),image.Point{-150,50})
		row++
	}
	for x:=rect.Min.X+offset.X; x<=rect.Max.X; x+=step.X {
		if (origo.X - col) == 0 {
			VerLine(x-1,rect.Min.Y,rect.Max.Y)
			VerLine(x+1,rect.Min.Y,rect.Max.Y)
		}
		VerLine(x,rect.Min.Y,rect.Max.Y)
		DrawString(x,rect.Min.Y,fmt.Sprintf(format,col-origo.X),image.Point{-50,-50})
		col++
	}
}


// Should we send points as an Image.Point instead? 
// Offset allows us to move the text in percent of the size that is total width x height
// which makes it possible to neatly center the text on the middle of the lines
func DrawString(x0 int, y0 int, text string, offset image.Point){
	f := truetype.NewFace(cfont,
			&truetype.Options{
				Size: 12.0,
				DPI: 72.0,
	// Who came up with all of this what on earth is hinting
				Hinting: font.HintingNone,
			})
// Create an infinite image
	u := image.NewUniform(fg)

	d := &font.Drawer{
		Dst: dst,
		Src: u, // how do I make an image from a color? 
		// here go complains over an error at the last line instead of the next from this one contrary to the way gcc does it. 
		Face: f,
	}

	m:=f.Metrics()
	d.Dot=fixed.Point26_6{
		X: d.MeasureString(text)*fixed.I(offset.X)/fixed.I(100)+fixed.I(x0),
		Y: m.Height*fixed.I(offset.Y)/fixed.I(100)+fixed.I(y0),
	}

	d.DrawString(text);
}
