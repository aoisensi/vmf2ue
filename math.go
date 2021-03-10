package main

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

func (x Vec3) Center(y Vec3) Vec3 {
	return x.Add(y).Mul(0.5)
}
