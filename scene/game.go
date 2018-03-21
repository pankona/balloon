package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/simlog"
)

// Game represents a scene object for main game scene
type Game struct {
	simra simra.Simraer
}

// Initialize function for simra.Scener interface
func (ga *Game) Initialize(sim simra.Simraer) {
	simlog.FuncIn()
	ga.simra = sim
	simlog.FuncOut()
}

// Drive function for simra.Driver interface
func (ga *Game) Drive() {
}
