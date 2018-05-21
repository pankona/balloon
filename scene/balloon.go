package scene

import (
	"fmt"
)

const (
	balloonUpspeed = 60
	balloonWidth   = 64 * 2
	balloonHeight  = 64 * 2
)

// balloon represents balloon itself
type balloon struct {
	eq     chan *command
	x, y   float32
	zindex int
}

func newBalloon(eq chan *command, x, y, z int) *balloon {
	return &balloon{
		eq:     eq,
		x:      float32(x),
		y:      float32(y),
		zindex: z,
	}
}

func (b *balloon) isOutOfScreen(screenWidth, screenHeight int) bool {
	return 0 > b.x || float32(screenWidth) < b.x ||
		0 > b.y || float32(screenHeight) < b.y
}

func (b *balloon) doAction() {
	b.y += float32(balloonUpspeed) / float32(60) / float32(b.zindex+1)
	fmt.Printf("i'm %p. position = (%010f, %010f)\n", b, b.x, b.y)
	if b.isOutOfScreen(configScreenWidth, configScreenHeight) {
		fmt.Println("i'm out of screen!")
		b.eq <- newCommand(cmdBalloonOutOfScreen, b)
	}
}

func (b *balloon) finalize() {
	// TODO: implement
}

func (b *balloon) getPosition() (float32, float32) {
	return b.x, b.y
}

func (b *balloon) getScale() (float32, float32) {
	// TODO: this is temporary implementation.
	return balloonWidth / float32(b.zindex), balloonHeight / float32(b.zindex)
}

func (b *balloon) getZIndex() int {
	return b.zindex
}
