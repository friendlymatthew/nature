package botany

import (
	"image/color"

	"github.com/friendlymatthew/nature/pkg/vec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Node struct {
	parent *Node

	position  vec.Position
	direction vec.Vec2
}

func NewNode(parent *Node, p vec.Position, d vec.Vec2) Node {
	return Node{parent: parent, position: p, direction: d}
}

func (a *Node) Draw(screen *ebiten.Image) {
	clr := color.RGBA{R: 255, G: 255, B: 0, A: 255}
	vector.DrawFilledCircle(screen, a.position.X, a.position.Y, 4.0, clr, false)
}
