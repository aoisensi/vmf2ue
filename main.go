package main

import (
	"fmt"
	"os"

	"github.com/galaco/vmf"
)

var out *os.File

func main() {
	file, _ := os.Open("itemtest.vmf")
	out, _ = os.Create("out.txt")

	reader := vmf.NewReader(file)
	level, _ := reader.Read()

	fmt.Fprintf(out, "Begin Map\n")
	fmt.Fprintf(out, "   Begin Level\n")
	//brush
	for solidID, solid := range level.World.GetChildrenByKey("solid") {
		fmt.Fprintf(out, "      Begin Actor Class=/Script/Engine.Brush\n")
		fmt.Fprintf(out, "         Begin Brush Name=Model_%v\n", solidID)
		fmt.Fprintf(out, "            Begin PolyList\n")
		for sideID, side := range solid.GetChildrenByKey("side") {
			fmt.Fprintf(out, "               Begin Polygon Link=%v\n", sideID)
			plane := parsePlane(side.GetProperty("plane"))
			for _, v := range plane {
				fmt.Fprintf(out, "                  Vertex   %+013.6f,%+013.6f,%+013.6f\n",
					v[0], v[1], v[2])

			}
			fmt.Fprintf(out, "               End Polygon\n")
		}
		fmt.Fprintf(out, "            End PolyList\n")
		fmt.Fprintf(out, "         End Brush\n")
		fmt.Fprintf(out, "      Brush=Model'\"Model_%v\"'\n", solidID)
		fmt.Fprintf(out, "      End Actor\n")
	}
	fmt.Fprintf(out, "   End Level\n")
	fmt.Fprintf(out, "End Map\n")
}

func parsePlane(plane string) (r [4]Vec3) {
	fmt.Sscanf(plane, "(%v %v %v) (%v %v %v) (%v %v %v)",
		&r[0][0], &r[0][1], &r[0][2],
		&r[1][0], &r[1][1], &r[1][2],
		&r[2][0], &r[2][1], &r[2][2])
	r[0][1] *= -1
	r[1][1] *= -1
	r[2][1] *= -1
	c := r[0].Center(r[2])
	r[3] = c.Add(c.Sub(r[1]))
	return
}
