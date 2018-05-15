package scene

import "fmt"

const upspeed = 10

// balloon represents balloon itself
type balloon struct {
	eq     chan *command
	x, y   int // position
	zindex int
}

func newBalloon(eq chan *command, x, y, z int) *balloon {
	return &balloon{
		eq:     eq,
		x:      x,
		y:      y,
		zindex: z,
	}
}

func (b *balloon) isOutOfScreen() bool {
	return false
}

func (b *balloon) doAction() {
	b.y += upspeed
	fmt.Printf("i'm %p. position = (%d, %d)\n", b, b.x, b.y)
	if b.isOutOfScreen() {
		b.eq <- newCommand(cmdBalloonOutOfScreen, b)
	}
}
