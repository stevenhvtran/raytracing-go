package geometry

type Ray struct {
	Origin    Vector
	Direction Vector
}

func At(ray Ray, t float64) Vector {
	return Add(ray.Origin, Multiply(ray.Direction, t))
}
