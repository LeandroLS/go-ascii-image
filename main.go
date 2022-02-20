package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

func getChar(grayScale int) string {
	// densityCharsArr := []string{"$", "@", "B", "%", "8", "&", "W", "M", "#", "*", "o", "a", "h", "k", "b", "d", "p", "q", "w", "m", "Z", "O", "0", "Q", "L", "C", "J", "U", "Y", "X", "z", "c", "v", "u", "n", "x", "r", "j", "f", "t", "/", "\\", "|", "(", ")", "1", "{", "}", "[", "]", "?", "-", "_", "+", "~", "<", ">", "i", "!", "l", "I", ";", ":", ",", ",", "^", "`", "'", "."}
	// densityCharsArr := []string{"#", "W", "o", "-", " "}
	densityCharsArr := []string{" ", "-", "o", "W", "#"}

	interval := float64(len(densityCharsArr)) / float64(256)
	return densityCharsArr[int(math.Floor(float64(grayScale)*interval))]
}

func main() {
	oneCharWidth := 8
	oneCharHeight := 18
	imgfile, err := os.Open("./lion.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	// get image height and width with image/jpeg
	// change accordinly if file is png or gif

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width : ", width)
	fmt.Println("Height : ", height)

	// we need to reset the io.Reader again for image.Decode() function below to work
	// otherwise we will  - panic: runtime error: invalid memory address or nil pointer dereference
	// there is no build in rewind for io.Reader, use Seek(0,0)
	imgfile.Seek(0, 0)

	// get the image
	img, _, _ := image.Decode(imgfile)
	imgResized := resize.Resize(uint(width), uint(height*(oneCharWidth/oneCharHeight)), img, resize.Lanczos3)
	f, _ := os.Create("data.txt")
	fmt.Println("Width : ", width)
	fmt.Println("Height : ", height)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := imgResized.At(x, y).RGBA()
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
