package botany

import (
	"github.com/friendlymatthew/nature/pkg/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tree struct {

	// a fancy term for leaves
	attractors []Attractor
	nodes      []Node
}

func NewTree(attrCount int, sW, sH float32) *Tree {
	attractors := make([]Attractor, attrCount)

	for i := range attractors {
		attractors[i] = *NewAttractor(sW, sH)
	}

	root := NewNode(
		nil,
		vec.Position{
			X: sW / 2,
			Y: sH / 2,
		},
		vec.NewVec2(0, -1),
	)

	nodes := []Node{root}

	return &Tree{
		nodes:      nodes,
		attractors: attractors,
	}
}

func (t *Tree) Draw(screen *ebiten.Image) {
	for _, a := range t.attractors {
		a.Draw(screen)
	}

	for _, n := range t.nodes {
		n.Draw(screen)
	}
}
