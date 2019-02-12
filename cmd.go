package TextToImage

import (
	"flag"
	"os"
)

var (
	textPtr 	*string
	sizePtr 	*int
	fontPtr 	*string
	outputPtr 	*string
)

func StartApp() {
	textPtr = flag.String("text", "", "Your text (required)")
	sizePtr = flag.Int("size", 0, "Size of the text in points (required)")
	fontPtr = flag.String("font", "Arial", "Name of font to be used")
	outputPtr = flag.String("out", "", "Output path of file")

	flag.Parse()

	// check if required flags are empty
	if *textPtr == "" && *sizePtr == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *outputPtr == "" {
		path := *textPtr + ".png"
		outputPtr = &path
	}

	drawText(*TextPtr, *FontPtr, *SizePtr, *OutputPtr)
}