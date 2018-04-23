package scene

import (
	"sync"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
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
}

// Initialize function for simra.Scener interface
func (ga *Game) Initialize(sim simra.Simraer) {
	ga.simra = sim
	ga.loadTextures()
	ga.prepareSprites()
	ga.updateRunLoop(nil, ga.runLoopReady)
}

type progress int

const (
	progressReady progress = iota
	progressRunning
	progressFinished
)

func (ga *Game) doOncePrepare() {
	if ga.prepareFunc != nil {
		ga.prepare.Do(ga.prepareFunc)
	}
}

func (ga *Game) updateRunLoop(prepare, runloop func()) {
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
