package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/friendlymatthew/nature/pkg/botany"
)

const (
	WIDTH  = 500
	HEIGHT = 500
)

type Canvas struct {
	tree *botany.Tree
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (c *Canvas) Update() error {
	c.tree.Grow()

	return nil
}

func (c *Canvas) Draw(screen *ebiten.Image) {

	c.tree.Draw(screen)
}

func (c *Canvas) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func main() {

	log.Println("Starting the game...")

	attrCount := 400

	c := &Canvas{
		tree: botany.NewTree(attrCount, WIDTH, HEIGHT),
	}
	log.Printf("Initialized tree with %d attractors\n", attrCount)

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("trees")
	log.Println("Window setup complete. Running the game...")

	if err := ebiten.RunGame(c); err != nil {
		log.Fatal(err)
	}
}
