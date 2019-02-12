package TextToImage

import (
	"flag"
	"os"
)

var (
	TextPtr 	*string
	SizePtr 	*int
	FontPtr 	*string
	OutputPtr 	*string
)

func StartApp() {
	TextPtr = flag.String("text", "", "Your text (required)")
	SizePtr = flag.Int("size", 0, "Size of the text in points (required)")
	FontPtr = flag.String("font", "Arial", "Name of font to be used")
	OutputPtr = flag.String("out", "", "Output path of file")

	flag.Parse()

	// check if required flags are empty
	if *TextPtr == "" && *SizePtr == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *OutputPtr == "" {
		path := *TextPtr + ".png"
		OutputPtr = &path
	}

	drawText(*TextPtr, *FontPtr, *SizePtr, *OutputPtr)
}