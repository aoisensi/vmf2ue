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
	UAxis    [5]float64
	VAxis    [5]float64
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
	dispInfo := node.GetChildrenByKey("dispinfo")
	if len(dispInfo) == 1 {
		side.DispInfo = newDispInfo(dispInfo[0])
	}
	fmt.Sscanf(
		node.GetProperty("uaxis"),
		"[%v %v %v %v] %v",
		&side.UAxis[0], &side.UAxis[1], &side.UAxis[2], &side.UAxis[3], &side.UAxis[4],
	)
	fmt.Sscanf(
		node.GetProperty("vaxis"),
		"[%v %v %v %v] %v",
		&side.VAxis[0], &side.VAxis[1], &side.VAxis[2], &side.VAxis[3], &side.VAxis[4],
	)
	return side
}
