package TextToImage

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"log"
)


func drawText(text string, fontFamily string, size int, path string) {

	fontBytes, err := ioutil.ReadFile("/Library/Fonts/" + fontFamily + ".ttf")
	if err != nil {
		log.Fatal("Error, could not find the .ttf file for the font specified")
		return
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	opts := truetype.Options{}
	opts.Size = float64(size)

	ff := truetype.NewFace(f, &opts)
	w := font.MeasureString(ff, text)

	dc := gg.NewContext(w.Round(), int(1.2*float64(size)))

	dc.SetRGB(0,0,0)
	dc.SetFontFace(ff)
	dc.DrawStringAnchored(text, 0, float64(size)/2, 0, 0.5)

	err = dc.SavePNG(path)
	if err != nil {
		log.Fatal("Error, could not generate image")
		return
	}

}