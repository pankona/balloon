package scene

import (
	"context"

	"github.com/pankona/gomo-simra/simra/fps"
)

// spawner emits spawn event in random
type spawner struct {
	eq    chan *command
	count int
}

func (s *spawner) run(context context.Context) {
	for {
		select {
		case <-context.Done():
			return
		case <-fps.After(60):
			// TODO: in random
			if s.count > 10 {
				return
			}
			s.eq <- newCommand(cmdBalloonSpawn, nil)
			s.count++
		}
	}
}

func (s *spawner) OnEvent(e interface{}) {
	_ = e.(*command)
	// TODO: implement
}
