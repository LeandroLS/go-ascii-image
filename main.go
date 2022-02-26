package main

import (
	"fmt"
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
	densityCharsArr := []string{"@@","@@","@@","@@","@@","@@","@@","@@","%%","%%","%%","%%","%%","%%","%%","%%","##","##","##","##","##","##","##","##","##","**","**","**","**","**","**","**","**","++","++","++","++","++","++","++","++","++","==","==","==","==","==","==","==","=="}
	interval := float64(len(densityCharsArr)) / float64(256)
	return densityCharsArr[int(math.Floor(float64(grayScale)*interval))]
}

func main() {
	imgfile, err := os.Open("./woman3.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()
	imgfile.Seek(0, 0)

	img, _, _ := image.Decode(imgfile)
	imgResized := resize.Resize(100, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	out.Seek(0, 0)

	jpeg.Encode(out, imgResized, nil)
	f, _ := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	imgfile2, err := os.Open("./test_resized.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile2.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgfile2.Seek(0, 0)

	img2, _, _ := image.Decode(imgfile2)
	grayImg := image.NewGray(img2.Bounds())
	for y := img2.Bounds().Min.Y; y < img2.Bounds().Max.Y; y++ {
		for x := img2.Bounds().Min.X; x < img2.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img2.At(x, y))
			r, g, b, _ := grayImg.At(x, y).RGBA()
			avg := uint8((r + g + b) / 3)
			_, err2 := f.WriteString(getChar(int(avg)))
			if err2 != nil {
				log.Fatal(err2)
			}
		}
		_, err2 := f.WriteString("\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}

}
