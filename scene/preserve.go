package scene

// preserve is a pool of actioner
type preserve struct {
	eq        chan *command
	actioners map[actioner]bool
}

func (p *preserve) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case cmdBalloonSpawned:
		b := cmd.payload.(actioner)
		p.actioners[b] = true
	case cmdBalloonOutOfScreen:
		b := cmd.payload.(actioner)
		b.finalize()
		delete(p.actioners, b)
	case cmdDoAction:
		for k := range p.actioners {
			k.doAction()
		}
	}
}
