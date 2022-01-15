package scene

import (
	"raytracing/geometry"
)

type Camera struct {
	ViewportWidth  float64
	ViewportHeight float64
	FocalLength    float64
	Origin         geometry.Vector
}

func Horizontal(camera Camera) geometry.Vector {
	return geometry.Vector{X: camera.ViewportWidth, Y: 0, Z: 0}
}

func Vertical(camera Camera) geometry.Vector {
	return geometry.Vector{X: 0, Y: camera.ViewportHeight, Z: 0}
}

func LowerLeftCorner(camera Camera) geometry.Vector {
	corner := geometry.Divide(geometry.Add(Horizontal(camera), Vertical(camera)), 2)
	depth := geometry.Vector{X: 0, Y: 0, Z: camera.FocalLength}
	return geometry.Subtract(geometry.Subtract(camera.Origin, corner), depth)
}
