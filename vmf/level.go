package vmf

import "github.com/galaco/vmf"

type Level struct {
	World *World
}

func newLevel(vmf vmf.Vmf) *Level {
	level := new(Level)
	level.World = newWorld(vmf.World)
	return level
}
