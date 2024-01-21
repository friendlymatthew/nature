package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/friendlymatthew/nature/pkg/botany"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH  = 500
	HEIGHT = 500
	ATTR   = 200
)

type Canvas struct {
	tree         *botany.Tree
	treeDoneTime time.Time
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *Canvas) Update() error {
	if !c.tree.IsDone {
		c.tree.Grow()
	} else {

		if c.treeDoneTime.IsZero() {
			c.treeDoneTime = time.Now()
		} else if time.Since(c.treeDoneTime).Seconds() >= 2 {
			// Tree is done, create a new one
			c.tree = botany.NewTree(ATTR, WIDTH, HEIGHT)
			c.treeDoneTime = time.Time{}
		}
	}

	return nil
}

func (c *Canvas) Draw(screen *ebiten.Image) {
	// Fill the entire window
	screen.Fill(color.RGBA{255, 255, 255, 255})

	canvas := ebiten.NewImage(WIDTH, HEIGHT)

	c.tree.Draw(canvas)

	// Draw the canvas onto the screen at the offset position
	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(canvas, opts)
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func main() {

	c := &Canvas{
		tree: botany.NewTree(ATTR, WIDTH, HEIGHT),
	}

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("trees")

	if err := ebiten.RunGame(c); err != nil {
		log.Fatal(err)
	}
}
