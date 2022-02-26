package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func getChar(grayScale int) string {
	densityCharsArr := []string{"@@", "@@", "@@", "@@", "@@", "@@", "@@", "@@", "%%", "%%", "%%", "%%", "%%", "%%", "%%", "%%", "##", "##", "##", "##", "##", "##", "##", "##", "##", "**", "**", "**", "**", "**", "**", "**", "**", "++", "++", "++", "++", "++", "++", "++", "++", "++", "==", "==", "==", "==", "==", "==", "==", "=="}
	interval := float64(len(densityCharsArr)) / float64(256)
	return densityCharsArr[int(math.Floor(float64(grayScale)*interval))]
}

func openImgFile(imgFileName string) image.Image {
	imgOpened, err := os.Open(imgFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgOpened.Close()

	img, _, err := image.Decode(imgOpened)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	userImg := openImgFile("./cat.jpg")

	userImgResized := resize.Resize(100, 0, userImg, resize.Lanczos3)

	f, err := os.Create("ascii.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	grayImg := image.NewGray(userImgResized.Bounds())
	for y := userImgResized.Bounds().Min.Y; y < userImgResized.Bounds().Max.Y; y++ {
		for x := userImgResized.Bounds().Min.X; x < userImgResized.Bounds().Max.X; x++ {
			grayImg.Set(x, y, userImgResized.At(x, y))
			r, g, b, _ := grayImg.At(x, y).RGBA()
			avg := uint8((r + g + b) / 3)
			_, err := f.WriteString(getChar(int(avg)))
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err := f.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
