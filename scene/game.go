package scene

import (
	"context"

	"github.com/pankona/gomo-simra/simra"
)

const (
	eqlen = 512
)

// Game represents a scene object for main game scene
type Game struct {
	simra          simra.Simraer
	currentRunLoop func()
	currentFrame   int64
	pubsub         *simra.PubSub
	eq             chan *command
	ctx            context.Context
	cancel         func()
}

// Initialize function for simra.Scener interface
func (ga *Game) Initialize(sim simra.Simraer) {
	ga.simra = sim
	ga.pubsub = simra.NewPubSub()
	ga.eq = make(chan *command, eqlen)
	ga.ctx, ga.cancel = context.WithCancel(context.Background())
	ga.simra.AddTouchListener(ga)
	ga.updateRunLoop(ga.runLoopReady)
}

func (ga *Game) updateRunLoop(runloop func()) {
	ga.currentRunLoop = runloop
}

// Drive function for simra.Driver interface
func (ga *Game) Drive() {
	ga.currentFrame++
	ga.currentRunLoop()
}

// OnTouchBegin is called at beginning of touch
func (ga *Game) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called at moving of touch
func (ga *Game) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called at releasing of touch
func (ga *Game) OnTouchEnd(x, y float32) {
	ga.updateRunLoop(ga.runLoopFinished)
}
