package vmf

import (
	"github.com/galaco/vmf"
)

type Level struct {
	World    *World
	Entities []Entity
}

func newLevel(data vmf.Vmf) *Level {
	level := new(Level)
	level.World = newWorld(data.World)
	entities := *data.Entities.GetAllValues()
	level.Entities = make([]Entity, len(entities))
	for i, entity := range entities {
		level.Entities[i] = newEntity(entity.(vmf.Node))
	}
	return level
}
