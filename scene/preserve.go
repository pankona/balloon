package scene

// preserve is a pool of actioner
type preserve struct {
	eq        chan *command
	actioners map[actioner]bool
}

func (p *preserve) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case commandSpawned:
		b := cmd.payload.(actioner)
		p.actioners[b] = true
	case commandDoAction:
		for k := range p.actioners {
			k.doAction()
		}
	}
}
