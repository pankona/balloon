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
	fmt.Println("runLoopPrepareRunning!")
	// TODO: implement
}

func (ga *Game) runLoopRunning() {
	fmt.Println("runLoopRunning!")
	// TODO: implement
}

func (ga *Game) runLoopPrepareFinished() {
	// TODO: implement
}

func (ga *Game) runLoopFinished() {
	// TODO: implement
}
