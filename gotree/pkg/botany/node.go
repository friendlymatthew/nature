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

	avgDir vec.Vec2
	count  int
}

func NewNode(parent *Node, p vec.Position, d vec.Vec2) Node {
	return Node{parent: parent, position: p, direction: d}
}

func (a *Node) Draw(screen *ebiten.Image) {
	// Set the color for the current node (fully opaque white in this case)
	nodeColor := color.RGBA{R: 26, G: 26, B: 26, A: 255}

	// Draw the current node as a filled circle
	vector.DrawFilledCircle(screen, a.position.X, a.position.Y, 1.3, nodeColor, false)
}
