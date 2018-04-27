package scene

import (
	"context"
	"sync"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

const (
	eqlen = 512
)

// Game represents a scene object for main game scene
type Game struct {
	simra          simra.Simraer
	kokeshi        simra.Spriter
	kokeshiTex     *simra.Texture
	progress       progress
	currentRunLoop func()
	currentFrame   int64
	prepareFunc    func()
	prepare        sync.Once
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
	ga.loadTextures()
	ga.prepareSprites()
	ga.updateRunLoop(ga.runLoopPrepareReady, ga.runLoopReady)
}

type progress int

const (
	progressReady progress = iota
	progressRunning
	progressFinished
)

// TODO: refactoring
func (ga *Game) doOncePrepare() {
	if ga.prepareFunc != nil {
		ga.prepare.Do(ga.prepareFunc)
	}
}

func (ga *Game) updateRunLoop(prepare, runloop func()) {
	ga.prepare = sync.Once{}
	ga.prepareFunc = prepare
	ga.currentRunLoop = runloop
}

func (ga *Game) loadTextures() {
	ga.kokeshiTex = ga.simra.NewImageTexture("kokeshi_64x64.png", image.Rect(0, 0, 64, 64))
}

func (ga *Game) prepareSprites() {
	ga.kokeshi = ga.simra.NewSprite()
	ga.kokeshi.SetScale(64, 64)
	ga.kokeshi.SetPosition(configScreenWidth/2, configScreenHeight/2)
	ga.simra.AddSprite(ga.kokeshi)
	ga.kokeshi.ReplaceTexture(ga.kokeshiTex)
}

// Drive function for simra.Driver interface
func (ga *Game) Drive() {
	ga.currentFrame++
	ga.doOncePrepare()
	ga.currentRunLoop()
}

func (ga *Game) OnTouchBegin(x, y float32) {
}

func (ga *Game) OnTouchMove(x, y float32) {
}

func (ga *Game) OnTouchEnd(x, y float32) {
	ga.updateRunLoop(ga.runLoopPrepareFinished, ga.runLoopFinished)
}
