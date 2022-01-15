package geometry

import (
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func Add(u, v Vector) Vector {
	return Vector{u.X + v.X, u.Y + v.Y, u.X + v.X}
}

func Subtract(u, v Vector) Vector {
	return Vector{u.X - v.X, u.Y - v.Y, u.X - v.X}
}

func Dot(u, v Vector) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*u.Z
}

func Multiply(u Vector, t float64) Vector {
	return Vector{u.X * t, u.Y * t, u.Z * t}
}

func Divide(u Vector, t float64) Vector {
	return Multiply(u, 1/t)
}

func LengthSquared(u Vector) float64 {
	return u.X*u.X + u.Y*u.Y + u.Z*u.Z
}

func Length(u Vector) float64 {
	return math.Sqrt(u.X*u.X + u.Y*u.Y + u.Z*u.Z)
}

func Cross(u, v Vector) Vector {
	return Vector{u.Y*v.Z - u.Z*v.Y, -(u.X*v.Z - u.Z*v.X), u.X*v.Y - u.Y*v.X}
}
