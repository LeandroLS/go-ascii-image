package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func Init() (image.Image, int) {
	width := flag.Int("w", 80, "Use -w <width>")
	fpath := flag.String("p", "test.jpg", "Use -p <filesource>")
	flag.Parse()

	f, err := os.Open(*fpath)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
	return img, *width
}

func getChar(grayScale int) string {
	densityCharsArr := []string{"@@", "@@", "@@", "@@", "@@", "@@", "@@", "@@", "%%", "%%", "%%", "%%", "%%", "%%", "%%", "%%", "##", "##", "##", "##", "##", "##", "##", "##", "##", "**", "**", "**", "**", "**", "**", "**", "**", "++", "++", "++", "++", "++", "++", "++", "++", "++", "==", "==", "==", "==", "==", "==", "==", "=="}
	interval := float64(len(densityCharsArr)) / float64(256)
	return densityCharsArr[int(math.Floor(float64(grayScale)*interval))]
}

func main() {

	//go:embed test.jpg

	userImg, width := Init()

	userImgResized := resize.Resize(uint(width), 0, userImg, resize.Lanczos3)

	var content string
	grayImg := image.NewGray(userImgResized.Bounds())
	for y := userImgResized.Bounds().Min.Y; y < userImgResized.Bounds().Max.Y; y++ {
		for x := userImgResized.Bounds().Min.X; x < userImgResized.Bounds().Max.X; x++ {
			grayImg.Set(x, y, userImgResized.At(x, y))
			r, g, b, _ := grayImg.At(x, y).RGBA()
			avg := uint8((r + g + b) / 3)
			char := getChar(int(avg))
			content += char
		}
		content += ("\n")
	}

	f, err := os.Create("ascii.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fmt.Println(content)

	f.WriteString(content)

	fmt.Printf("ASCII created in %s file", f.Name())
}
