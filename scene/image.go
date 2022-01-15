package scene

type Image struct {
	Pixels      [][][3]uint8
	ImageWidth  int
	ImageHeight int
}

func NewImage(imageWidth, imageHeight int) (image Image) {
	image.ImageWidth = imageWidth
	image.ImageHeight = imageHeight

	image.Pixels = make([][][3]uint8, imageHeight)
	for height := 0; height < imageHeight; height++ {
		image.Pixels[height] = make([][3]uint8, imageWidth)
	}

	return
}
