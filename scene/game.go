package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

// Game represents a scene object for main game scene
type Game struct {
	simra      simra.Simraer
	kokeshi    simra.Spriter
	kokeshiTex *simra.Texture
}

// Initialize function for simra.Scener interface
func (ga *Game) Initialize(sim simra.Simraer) {
	ga.simra = sim
	ga.loadTextures()
	ga.prepareSprites()
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
}
