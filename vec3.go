package main

import "math"

type vec3 struct {
	X float64
	Y float64
	Z float64
}

func newVec3(x, y, z float64) vec3 {
	return vec3{X: x, Y: y, Z: z}
}

func (v vec3) Neg() vec3 {
	return newVec3(-v.X, -v.Y, -v.Z)
}

func (v vec3) Add(o vec3) vec3 {
	return newVec3(v.X+o.X, v.Y+o.Y, v.Z+o.Z)
}

func (v vec3) Sub(o vec3) vec3 {
	return newVec3(v.X-o.X, v.Y-o.Y, v.Z-o.Z)
}

// Scalar Multiplication
func (v vec3) Mul(t float64) vec3 {
	return newVec3(v.X*t, v.Y*t, v.Z*t)
}

// Scalar Division
func (v vec3) Div(t float64) vec3 {
	return newVec3(v.X/t, v.Y/t, v.Z/t)
}

func (v vec3) Dot(o vec3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v vec3) Cross(o vec3) vec3 {
	return newVec3(
		v.Y*o.Z-v.Z*o.Y,
		v.Z*o.X-v.X-o.Z,
		v.X*o.Y-v.Y*o.X,
	)
}

func (v vec3) lengthSquared() float64 {
	return v.Dot(v)
}

func (v vec3) length() float64 {
	return math.Sqrt(v.lengthSquared())
}

func (v vec3) normalize() vec3 {
	return v.Div(v.length())
}
