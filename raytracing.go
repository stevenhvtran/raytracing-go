package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const imageWidth = 256
	const imageHeight = 256

	image := makeImageSlice(imageWidth, imageHeight)
	renderImage(imageWidth, imageHeight, image)
	writeToFile("image.ppm", imageToPpm(imageWidth, imageHeight, image))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(fileName, contents string) {
	err := os.WriteFile(fileName, []byte(contents), 0666)
	check(err)
}

func makeImageSlice(imageWidth, imageHeight int) (image [][][3]uint8) {
	image = make([][][3]uint8, imageHeight)
	for height := 0; height < imageHeight; height++ {
		image[height] = make([][3]uint8, imageWidth)
	}
	return
}

func renderImage(imageWidth, imageHeight int, image [][][3]uint8) {
	for j := 0; j < imageHeight; j++ {
		log.Printf("Scanlines remaining: %d", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (float64(imageWidth) - 1)
			g := (float64(imageHeight) - float64(j) - 1) / (float64(imageHeight) - 1)
			b := 0.25

			image[j][i][0] = uint8(r * 255.99)
			image[j][i][1] = uint8(g * 255.99)
			image[j][i][2] = uint8(b * 255.99)
		}
	}
	log.Print("Done.")
}

func imageToPpm(imageWidth int, imageHeight int, image [][][3]uint8) (ppm string) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", imageWidth, imageHeight))
	for _, row := range image {
		for _, pixel := range row {
			sb.WriteString(fmt.Sprintf("%d %d %d\n", pixel[0], pixel[1], pixel[2]))
		}
	}
	return sb.String()
}
