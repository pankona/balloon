package scene

import "fmt"

func (ga *Game) runLoopReady() {
	// TODO: implement to show like "Ready Go"
	fmt.Println("runLoopReady!")
	ga.updateRunLoop(ga.runLoopPrepareRunning)
}

func (ga *Game) runLoopPrepareRunning() {
	s := &spawner{eq: ga.eq}
	f := &factory{eq: ga.eq}
	p := &preserve{
		eq:       ga.eq,
		balloons: make(map[*balloon]bool),
	}
	ga.pubsub.Subscribe("spawner", s)
	ga.pubsub.Subscribe("factory", f)
	ga.pubsub.Subscribe("preserve", p)
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
