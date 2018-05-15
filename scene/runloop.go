package scene

import (
	"fmt"
)

func (ga *Game) runLoopReady() {
	// TODO: implement to show like "Ready Go"
	fmt.Println("runLoopReady!")
	ga.updateRunLoop(ga.runLoopPrepareRunning)
}

func mustNoError(m map[error]bool) {
	for k := range m {
		if k != nil {
			panic("failed to subscribe. " + k.Error())
		}
	}
}

func (ga *Game) runLoopPrepareRunning() {
	s := &spawner{eq: ga.eq}
	f := &factory{eq: ga.eq}
	p := &preserve{
		eq:       ga.eq,
		balloons: make(map[*balloon]bool),
	}

	m := make(map[error]bool)
	m[ga.pubsub.Subscribe("spawner", s)] = true
	m[ga.pubsub.Subscribe("factory", f)] = true
	m[ga.pubsub.Subscribe("preserve", p)] = true
	mustNoError(m)

	go s.run(ga.ctx)
	ga.updateRunLoop(ga.runLoopRunning)
}

func (ga *Game) runLoopRunning() {
	len := len(ga.eq)
	for i := 0; i < len; i++ {
		e := <-ga.eq
		ga.pubsub.Publish(e)
	}
}

func (ga *Game) runLoopFinished() {
	ga.cancel()
	close(ga.eq)
	ga.simra.SetScene(&Title{})
}
