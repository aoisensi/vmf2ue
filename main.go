package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aoisensi/vmf2ue/vmf"
)

var out *os.File
var nest = make([]string, 0, 16)

const scale = 2.0

func main() {
	file, _ := os.Open(os.Args[1])
	out, _ = os.Create("out.txt")
	readBind()

	level, _ := vmf.Read(file)
	writeBegin("Map", "")
	writeBegin("Level", "")

	//brush
	for _, solid := range level.World.Solids {
		writeBegin("Actor", "Class=/Script/Engine.Brush Name=Solid%v", solid.ID)
		writeBegin("Brush", "Name=Model_%v", solid.ID)
		writePolyList(solid)
		writeEnd()
		writef("Brush=Model'\"Model_%v\"'", solid.ID)
		writef("ActorLabel=\"Solid%v\"", solid.ID)
		writeEnd()
	}
	writeEnd()
	writeEnd()
}

func writePolyList(solid *vmf.Solid) {
	intersection := func(a, b, c Plane) *Vec3 {
		denom := a.V.Dot(b.V.Cross(c.V))
		if -0.000001 < denom && denom < 0.000001 {
			return nil
		}
		ad := b.V.Cross(c.V).Mul(-a.D)
		bd := c.V.Cross(a.V).Mul(-b.D)
		cd := a.V.Cross(b.V).Mul(-c.D)
		v := ad.Add(bd).Add(cd).Mul(1.0 / denom)
		return &v
	}

	writeBegin("PolyList", "")
	faces := make([]Plane, len(solid.Sides))
	for i, side := range solid.Sides {
		faces[i] = PlaneFromPoints(side.Plane[0], side.Plane[1], side.Plane[2])
	}
	for i, faceI := range faces {
		writeBegin("Polygon", "")
		verteces := make([]Vec3, 0, 16)
		for j, faceJ := range faces {
			if i == j {
				continue
			}
			for k, faceK := range faces {
				if i == k || j <= k {
					continue
				}
				v := intersection(faceI, faceJ, faceK)
				if v == nil {
					continue
				}
				ok := true
				for _, faceL := range faces {
					if faceL.V.Dot(*v)+faceL.D < -EPS {
						ok = false
						break
					}
				}
				if ok {
					verteces = append(verteces, *v)
				}
			}
		}

		center := Vec3{}
		for _, v := range verteces {
			center = center.Add(v)
		}
		center = center.Mul(1.0 / float64(len(verteces)))

		for n := 0; n < len(verteces)-2; n++ {
			a := verteces[n].Sub(center).Normalize()
			p := PlaneFromPoints(verteces[n], center, center.Add(faces[i].V))
			smallestAngle := -1.0
			smallest := -1
			for m := n + 1; m < len(verteces); m++ {
				side := p.Classify(verteces[m])
				if side < EPS {
					continue
				}
				b := verteces[m].Sub(center).Normalize()
				angle := a.Dot(b)
				if angle > smallestAngle {
					smallestAngle = angle
					smallest = m
				}
			}
			verteces[n+1], verteces[smallest] = verteces[smallest], verteces[n+1]
		}

		for _, v := range verteces {
			writef("Vertex   %+013.6f,%+013.6f,%+013.6f", v[1]*scale, v[0]*scale, v[2]*scale)
		}
		writeEnd()
	}
	writeEnd()
}

func writef(format string, a ...interface{}) {
	writeNest()
	fmt.Fprintf(out, format, a...)
	fmt.Fprintln(out)
}

func writeNest() {
	fmt.Fprint(out, strings.Repeat("   ", len(nest)))
}

func writeBegin(name string, format string, a ...interface{}) {
	writeNest()
	fmt.Fprint(out, "Begin "+name)
	if format != "" {
		fmt.Fprint(out, " ")
		fmt.Fprintf(out, format, a...)
	}
	fmt.Fprintln(out)
	nest = append(nest, name)
}

func writeEnd() {
	l := len(nest) - 1
	name := nest[l]
	nest = nest[:l]
	writeNest()
	fmt.Fprintln(out, "End "+name)
}

func calcPlane(plane [3][3]float64) (r [4]Vec3) {
	r[0] = plane[0]
	r[1] = plane[1]
	r[2] = plane[2]
	r[0][1] *= -1
	r[1][1] *= -1
	r[2][1] *= -1
	r[3] = r[0].Sub(r[1]).Add(r[2])
	return
}
