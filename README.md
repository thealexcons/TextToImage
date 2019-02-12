# Text to Image Converter
Simple text to image converter command line application written in Go

### Installation
You can get this package using the `go get` command:
```sh
$ go get github.com/thealexcons/TextToImage
```

### Usage
You can specify the text, font, size of text and output path of the image using command line flags:

```sh
$ ./TextToImage -text="Hello there!!" -size=120 -font="Impact" -out="hello.png"
```
The `-text` and `-size` are required flags. If you choose to omit the font and output path arguments, the image will be rendered using the Helvetica font and the image will be saved to the directory of the app binary.
