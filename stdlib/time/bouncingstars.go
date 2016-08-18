package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width    = 76
	height   = 10
	numBalls = 10
)

type screen [height][width]byte

func (s *screen) print() {
	fmt.Print("\x0c")
	for _, row := range *s {
		for i, b := range row {
			if b == 0 {
				row[i] = ' '
			}
		}
		fmt.Println(string(row[:]))
	}
}

func (s *screen) paint(x, y int, b byte) {
	s[y][x] = b
}

type ball struct {
	x, y   int
	xd, yd int
}

func newBall() *ball {
	return &ball{
		x:  rand.Intn(width),
		y:  rand.Intn(height),
		xd: (rand.Intn(2) * 2) - 1,
		yd: (rand.Intn(2) * 2) - 1,
	}
}

func (b *ball) move() {
	moveBounce(&b.x, &b.xd, width)
	moveBounce(&b.y, &b.yd, height)
}

func moveBounce(v, d *int, max int) {
	*v += *d
	if *v < 0 {
		*v = -*v
		*d = -*d
	} else if *v >= max {
		*v = max - 2 - (*v - max)
		*d = -*d
	}
}

func main() {
	balls := []*ball{}
	for i := 0; i < numBalls; i++ {
		balls = append(balls, newBall())
	}
	for i := 200; i >= 0; i-- {
		var s screen
		for _, b := range balls {
			b.move()
			s.paint(b.x, b.y, '*')
		}
		s.print()
		time.Sleep(time.Second / 20)
	}
}
