package vmf

import (
	"strconv"

	"github.com/galaco/vmf"
)

type World struct {
	ID     int
	Solids []*Solid
}

func newWorld(node vmf.Node) *World {
	world := new(World)
	var err error
	world.ID, err = strconv.Atoi(node.GetProperty("id"))
	if err != nil {
		panic(err)
	}
	solids := node.GetChildrenByKey("solid")
	world.Solids = make([]*Solid, len(solids))
	for i, solid := range solids {
		world.Solids[i] = newSolid(solid)
	}
	return world
}
