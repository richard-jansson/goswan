package colorscheme 

import (
	"image/color"
	"fmt"
)


/* 
 * If we treat R,G,B as binary we end up with these colors. 
 * And if you are not female this is what you ought to end 
 * up with. 
 * 
 * 2^3 colors 
 * 
 * http://www.thedoghousediaries.com/1406
 */
var (
	Black	= color.RGBA{0x00,0x00,0x00,0x00}

	Red	= color.RGBA{0xff,0x00,0x00,0x00}
	Green	= color.RGBA{0x00,0xff,0x00,0x00}
	Blue	= color.RGBA{0x00,0x00,0xff,0x00}

	Yellow	= color.RGBA{0xff,0xff,0x00,0x00}
	Magenta	= color.RGBA{0xff,0x00,0xff,0x00} // WTF is Magenta?
	Cyan	= color.RGBA{0x00,0xff,0xff,0x00} // WTF is cyan? 

	White	= color.RGBA{0xff,0xff,0xff,0xff}
)

// But yeah we'll continue 0x80 is 50% of full brightness
var Grey = color.RGBA{0x80,0x80,0x80,0x80}


/* 
 * Fun GO fact apparently the function has to be Capitalized. 
 * i.e. Making the first letter uppercase for the function to 
 * be exported. :w
 */
func Whatever() {
	a := 10
	fmt.Printf("%i\n",a);
}
