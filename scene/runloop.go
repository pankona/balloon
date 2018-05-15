package scene

import (
	"fmt"

	"github.com/pankona/gomo-simra/simra"
)

func (ga *Game) runLoopReady() {
	// TODO: implement to show like "Ready Go"
	fmt.Println("runLoopReady!")
	ga.updateRunLoop(ga.runLoopPrepareRunning)
}

func (ga *Game) runLoopPrepareRunning() {
	s := &spawner{eq: ga.eq}
	f := &factory{eq: ga.eq}
	p := &preserve{
		eq:        ga.eq,
		actioners: make(map[actioner]bool),
	}

	subscribers := []struct {
		id string
		s  simra.Subscriber
	}{
		{"spawner", s},
		{"factory", f},
		{"preserve", p},
	}

	for _, v := range subscribers {
		err := ga.pubsub.Subscribe(v.id, v.s)
		if err != nil {
			panic("failed to subscribe. " + err.Error())
		}
	}

	go s.run(ga.ctx)
	ga.updateRunLoop(ga.runLoopRunning)
}

func (ga *Game) runLoopRunning() {
	len := len(ga.eq)
	for i := 0; i < len; i++ {
		e := <-ga.eq
		ga.pubsub.Publish(e)
	}
	ga.pubsub.Publish(newCommand(commandDoAction, nil))
}

func (ga *Game) runLoopFinished() {
	ga.cancel()
	close(ga.eq)
	ga.simra.SetScene(&Title{})
}
