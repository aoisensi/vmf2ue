package vmf

import "github.com/galaco/vmf"

func newDispInfo(node vmf.Node) *DispInfo {
	return new(DispInfo)
}

type DispInfo struct {
	Power         int
	StartPosition [3]float64
	Flags         int
	Elevation     float64
	SubDiv        bool
	Normals       map[string][]float64 // float64 count is (3 * power * power)

}
