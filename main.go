package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoisensi/vmf2ue/vmf"
)

var out *os.File
var nest = make([]string, 0, 16)
var only = ""

const SCALE = 2.0
const BRIGHTNESS = 16.0

func init() {
	flag.StringVar(&only, "only", "", "only")
	flag.Parse()
}

func main() {
	file, _ := os.Open(flag.Arg(0))
	defer file.Close()
	out, _ = os.Create("out.txt")
	defer out.Close()
	readBind()

	level, _ := vmf.Read(file)
	writeBegin("Map", "")
	writeBegin("Level", "")

	for _, entity := range level.Entities {
		writeEntity(entity)
	}

	writeEnd()
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

func hasOnly(tag string) bool {
	if only == "" {
		return true
	}
	for _, t := range strings.Split(only, ",") {
		if t == tag {
			return true
		}
	}
	return false
}
