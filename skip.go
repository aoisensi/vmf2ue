package main

var (
	unknownMaterials = make(map[string]struct{})
	unknownMeshes    = make(map[string]struct{})
	unknownClasses   = make(map[string]struct{})
	skipClasses      = map[string]struct{}{}
)
