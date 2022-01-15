package main

import (
	"fmt"
	"log"
	"os"
	"raytracing/geometry"
	"raytracing/scene"
	"strings"
)

func main() {
	imageWidth := 400
	var aspectRatio float64 = 16 / 9
	image := scene.NewImage(imageWidth, int(float64(imageWidth)/aspectRatio))

	var viewportHeight float64 = 2
	camera := scene.Camera{ViewportWidth: viewportHeight * aspectRatio, ViewportHeight: viewportHeight, FocalLength: 1.0, Origin: geometry.Vector{X: 0, Y: 0, Z: 0}}

	renderImage(camera, image)
	writeToFile("image.ppm", imageToPpm(image))

	a := geometry.Vector{X: 1, Y: 2, Z: 3}
	geometry.Add(a, a)
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

func renderImage(camera scene.Camera, image scene.Image) {
	for j := 0; j < image.ImageHeight; j++ {
		log.Printf("Scanlines remaining: %d", image.ImageHeight-j)
		for i := 0; i < image.ImageWidth; i++ {
			image.Pixels[j][i] = normaliseColours(renderPixel(image, i, j))
		}
	}
	log.Print("Done.")
}

func renderPixel(image scene.Image, i, j int) (pixel [3]float64) {
	pixel[0] = float64(i) / (float64(image.ImageWidth) - 1)
	pixel[1] = (float64(image.ImageHeight) - float64(j) - 1) / (float64(image.ImageHeight) - 1)
	pixel[2] = 0.25
	return
}

func normaliseColours(pixel [3]float64) (normalisedPixel [3]uint8) {
	for i := 0; i < 3; i++ {
		normalisedPixel[i] = uint8(pixel[i] * 255.99)
	}
	return
}

func imageToPpm(image scene.Image) (ppm string) {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", image.ImageWidth, image.ImageHeight))
	for _, row := range image.Pixels {
		for _, pixel := range row {
			sb.WriteString(fmt.Sprintf("%d %d %d\n", pixel[0], pixel[1], pixel[2]))
		}
	}
	return sb.String()
}
