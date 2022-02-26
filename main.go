package main

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func getChar(grayScale int) string {
	// densityCharsArr := []string{"  ", "--", "oo", "WW", "##"}
	// densityCharsArr := []string{"==","==","==","==","==","==","==","--","--","--","--","--","--","--","--","::","::","::","::","::","::","::","::","::","..","..","..","..","..","..","..",".."}
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
	imgOpened.Seek(0, 0)
	img, _, err := image.Decode(imgOpened)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func resizeImg(image image.Image) *os.File {
	userImgResized := resize.Resize(100, 0, image, resize.Lanczos3)
	newUserImgResized, err := os.Create("resized-img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer newUserImgResized.Close()
	newUserImgResized.Seek(0, 0)

	jpeg.Encode(newUserImgResized, userImgResized, nil)
	imgResized, err := os.Open("resized-img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	return imgResized

}

func main() {
	userImg := openImgFile("./cat.jpg")
	imgResized := resizeImg(userImg)

	f, err := os.Create("ascii.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	img, _, _ := image.Decode(imgResized)
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
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
