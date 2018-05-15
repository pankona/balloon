package scene

type preserve struct {
	eq       chan *command
	balloons map[*balloon]bool
}

func (p *preserve) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case commandSpawned:
		b := cmd.payload.(*balloon)
		p.balloons[b] = true
	}
}
