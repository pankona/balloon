package scene

// balloon represents balloon itself
type balloon struct {
	x, y   int // position
	zindex int
}

func newBalloon(x, y, z int) *balloon {
	return &balloon{
		x:      x,
		y:      y,
		zindex: z,
	}
}
