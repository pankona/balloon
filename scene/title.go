package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/simlog"
)

// Title represents a scene object for title scene
type Title struct {
	simra         simra.Simraer
	sceneChangeCh chan struct{}
}

// Initialize function for simra.Scener interface
func (t *Title) Initialize(sim simra.Simraer) {
	simlog.FuncIn()

	t.simra = sim
	t.simra.SetDesiredScreenSize((float32)(configScreenHeight), (float32)(configScreenWidth))
	t.simra.AddTouchListener(t)

	t.sceneChangeCh = make(chan struct{}, 1)

	simlog.FuncOut()
}

// Drive function for simra.Driver interface
func (t *Title) Drive() {
	select {
	case <-t.sceneChangeCh:
		t.simra.SetScene(&Game{})
	default:
		// nop
	}
}

// OnTouchBegin function for simra.Toucher interface
func (t *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove function for simra.Toucher interface
func (t *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd function for simra.Toucher interface
func (t *Title) OnTouchEnd(x, y float32) {
	simlog.FuncIn()
	t.sceneChangeCh <- struct{}{}
	simlog.FuncOut()
}
