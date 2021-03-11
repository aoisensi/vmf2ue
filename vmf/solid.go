package vmf

import (
	"strconv"

	"github.com/galaco/vmf"
)

type Solid struct {
	ID    int
	Sides []*Side
}

func newSolid(node vmf.Node) *Solid {
	solid := new(Solid)
	var err error
	solid.ID, err = strconv.Atoi(node.GetProperty("id"))
	if err != nil {
		panic(err)
	}
	sides := node.GetChildrenByKey("side")
	solid.Sides = make([]*Side, len(sides))
	for i, side := range sides {
		solid.Sides[i] = newSide(side)
	}
	return solid
}
