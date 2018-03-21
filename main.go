// +build darwin linux

package main

import (
	"github.com/pankona/balloon/scene"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	sim := simra.NewSimra()
	sim.Start(&scene.Title{})
}
