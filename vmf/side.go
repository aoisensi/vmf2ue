package vmf

import (
	"fmt"
	"strconv"

	"github.com/galaco/vmf"
)

type Side struct {
	ID       int
	Plane    [3][3]float64
	Material string
	DispInfo *DispInfo
}

func newSide(node vmf.Node) *Side {
	side := new(Side)
	var err error
	side.ID, err = strconv.Atoi(node.GetProperty("id"))
	if err != nil {
		panic(err)
	}
	fmt.Sscanf(
		node.GetProperty("plane"),
		"(%f %f %f) (%f %f %f) (%f %f %f)",
		&side.Plane[0][0], &side.Plane[0][1], &side.Plane[0][2],
		&side.Plane[1][0], &side.Plane[1][1], &side.Plane[1][2],
		&side.Plane[2][0], &side.Plane[2][1], &side.Plane[2][2],
	)
	side.Material = node.GetProperty("material")
	if node.HasProperty("dispinfor") {
		side.DispInfo = newDispInfo(node.GetChildrenByKey("dispinfo")[0])
	}
	return side
}
