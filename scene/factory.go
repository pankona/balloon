package scene

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: declare on more suitable place
const maxZIndex = 5

// factory represents factory of balloons
type factory struct {
	eq chan *command
}

func (f *factory) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case commandSpawn:
		// TODO: implement
		fmt.Println("i'm balloon factory. spawn!")
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(configScreenWidth)
		z := rand.Intn(maxZIndex)
		f.eq <- newCommand(commandSpawned, newBalloon(x, 0, z))
	default:
		// nop
	}
}
