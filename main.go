// +build darwin linux

package main

import (
	"github.com/pankona/baloon/scene"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	sim := simra.NewSimra()
	sim.Start(&scene.Title{})
}
