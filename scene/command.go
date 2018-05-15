package scene

type commandtype int

const (
	cmdBalloonSpawn commandtype = iota
	cmdBalloonSpawned
	cmdBalloonOutOfScreen
	cmdDoAction
	cmdDraw
)

type command struct {
	ct      commandtype
	payload interface{}
}

func newCommand(ct commandtype, payload interface{}) *command {
	return &command{ct, payload}
}
