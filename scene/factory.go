package scene

import "fmt"

// factory represents factory of balloons
type factory struct {
}

func (f *factory) OnEvent(e interface{}) {
	cmd := e.(*command)
	switch cmd.ct {
	case commandSpawn:
		// TODO: implement
		fmt.Println("i'm balloon factory. spawn!")
	default:
		// nop
	}
}
