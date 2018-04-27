package scene

import "fmt"

func (ga *Game) runLoopPrepareReady() {
	// TODO: implement
	fmt.Println("runLoopPrepareReady!")
}

func (ga *Game) runLoopReady() {
	// TODO: implement
	fmt.Println("runLoopReady!")
	ga.updateRunLoop(ga.runLoopPrepareRunning, ga.runLoopRunning)
}

func (ga *Game) runLoopPrepareRunning() {
	s := &spawner{ga.eq}
	ga.pubsub.Subscribe("spawner", s)
	f := &factory{}
	ga.pubsub.Subscribe("factory", f)
	go s.run(ga.ctx)
}

func (ga *Game) runLoopRunning() {
	len := len(ga.eq)
	for i := 0; i < len; i++ {
		e := <-ga.eq
		ga.pubsub.Publish(e)
	}
}

func (ga *Game) runLoopPrepareFinished() {
	ga.cancel()
	close(ga.eq)
	ga.simra.SetScene(&Title{})
}

func (ga *Game) runLoopFinished() {
	// TODO: implement
}
