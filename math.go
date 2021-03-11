package main

import "math"

const EPS = 0.000001

type Vec3 [3]float64

func (x Vec3) Add(y Vec3) Vec3 {
	return Vec3{x[0] + y[0], x[1] + y[1], x[2] + y[2]}
}
func (x Vec3) Sub(y Vec3) Vec3 {
	return Vec3{x[0] - y[0], x[1] - y[1], x[2] - y[2]}
}

func (x Vec3) Mul(y float64) Vec3 {
	return Vec3{x[0] * y, x[1] * y, x[2] * y}
}

func (x Vec3) Dot(y Vec3) float64 {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2]
}

func (x Vec3) Cross(y Vec3) Vec3 {
	return Vec3{x[1]*y[2] - x[2]*y[1], x[2]*y[0] - x[0]*y[2], x[0]*y[1] - x[1]*y[0]}
}

func (x Vec3) Norm() float64 {
	return math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2])
}

func (x Vec3) Normalize() Vec3 {
	return x.Mul(1.0 / x.Norm())
}

type Plane struct {
	V Vec3
	D float64
}

func PlaneFromPoints(a, b, c Vec3) Plane {
	plane := Plane{}
	plane.V = b.Sub(a).Cross(c.Sub(a)).Normalize()
	plane.D = -plane.V.Dot(a)
	return plane
}

func (p Plane) Classify(v Vec3) float64 {
	return p.V.Dot(v) + p.D
}
