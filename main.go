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
	densityCharsArr := []string{"  ", "--", "oo", "WW", "##"}

	interval := float64(len(densityCharsArr)) / float64(256)
	return densityCharsArr[int(math.Floor(float64(grayScale)*interval))]
}

func main() {
	imgfile, err := os.Open("./lion.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()
	imgfile.Seek(0, 0)

	img, _, _ := image.Decode(imgfile)
	imgResized := resize.Resize(50, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	out.Seek(0, 0)
	// write new image to file
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

	imgCfg, _, err := image.DecodeConfig(imgfile2)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgfile2.Seek(0, 0)

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width : ", width)
	fmt.Println("Height : ", height)

	img2, _, _ := image.Decode(imgfile2)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			fmt.Println(r, g, b)
			avg := uint8((r + g + b) / 3)
			fmt.Println(avg)
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
