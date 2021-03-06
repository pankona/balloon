package scene

import (
	"math/rand"
	"time"
)

// TODO: declare on more suitable place
const maxZIndex = 5

func init() {
	rand.Seed(time.Now().UnixNano())
}

// factory represents factory of balloons
type factory struct {
	eq chan *command
}

func (f *factory) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case cmdBalloonSpawn:
		x := rand.Intn(configScreenWidth)
		z := rand.Intn(maxZIndex)
		f.eq <- newCommand(cmdBalloonSpawned, newBalloon(f.eq, x, 0, z))
	default:
		// nop
	}
}
