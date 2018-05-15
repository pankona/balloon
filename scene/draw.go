package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

type drawable interface {
	getPosition() (float32, float32)
	getScale() (float32, float32)
}

type drawer struct {
	simra   simra.Simraer
	drawers map[drawable]simra.Spriter
}

func (d *drawer) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case cmdBalloonSpawned:
		sp := d.simra.NewSprite()
		d.simra.AddSprite(sp)
		// TODO: this tex is temporary
		tex := d.simra.NewImageTexture(
			"kokeshi_64x64.png", image.Rect(0, 0, 64, 64))
		sp.ReplaceTexture(tex)
		b := cmd.payload.(drawable)
		d.drawers[b] = sp
	case cmdDraw:
		for k, v := range d.drawers {
			v.SetPosition(k.getPosition())
			v.SetScale(k.getScale())
		}
	}
}
