package vmf

import (
	"strconv"
	"strings"

	"github.com/galaco/vmf"
)

type Entity map[string]interface{}

func newEntity(node vmf.Node) Entity {
	entity := make(Entity)
	for _, node := range *node.GetAllValues() {
		node := node.(vmf.Node)
		key := *node.GetKey()
		entity[key] = (*node.GetAllValues())[0]
	}
	return entity
}

func (e Entity) Has(key string) bool {
	_, ok := e[key]
	return ok
}

func (e Entity) String(key string) string {
	return e[key].(string)
}

func (e Entity) Float(key string) float64 {
	v, _ := strconv.ParseFloat(e.String(key), 64)
	return v
}

func (e Entity) FloatSlice(key string) []float64 {
	vv := strings.Split(e.String(key), " ")
	v := make([]float64, len(vv))
	for i := range v {
		v[i], _ = strconv.ParseFloat(vv[i], 64)
	}
	return v
}
