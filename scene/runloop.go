package scene

import "fmt"

func (ga *Game) runLoopReady() {
	// TODO: implement to show like "Ready Go"
	fmt.Println("runLoopReady!")
	ga.updateRunLoop(ga.runLoopPrepareRunning)
}

func (ga *Game) runLoopPrepareRunning() {
	s := &spawner{ga.eq}
	ga.pubsub.Subscribe("spawner", s)
	f := &factory{}
	ga.pubsub.Subscribe("factory", f)
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
