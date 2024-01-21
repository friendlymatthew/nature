package botany

import (
	"github.com/friendlymatthew/nature/pkg/vec"
)

type Quad struct {
	Boundary *Rectangle
	Nodes    []Node
	Capacity int

	TopLeft     *Quad
	TopRight    *Quad
	BottomLeft  *Quad
	BottomRight *Quad
}

type Rectangle struct {
	Center vec.Position
	Width  float32
	Height float32
}

func NewRectangle(tl vec.Position, br vec.Position) Rectangle {
	w := br.X - tl.X
	h := br.Y - tl.Y

	center := vec.Position{
		X: (tl.X) + (w / 2),
		Y: (tl.Y) + (h / 2),
	}

	return Rectangle{
		Center: center,
		Width:  w,
		Height: h,
	}
}

func (r *Rectangle) Contains(point vec.Position) bool {
	return point.X >= r.Center.X-r.Width/2 &&
		point.X <= r.Center.X+r.Width/2 &&
		point.Y >= r.Center.Y-r.Height/2 &&
		point.Y <= r.Center.Y+r.Height/2
}

func NewQuad(boundary *Rectangle, capacity int) *Quad {
	return &Quad{
		Boundary: boundary,
		Nodes:    make([]Node, 0, capacity),
		Capacity: capacity,
	}
}

func (q *Quad) Insert(node Node) bool {
	if !q.Boundary.Contains(node.position) {
		return false
	}

	if len(q.Nodes) < q.Capacity {
		q.Nodes = append(q.Nodes, node)
		return true
	}

	if q.TopLeft == nil {
		q.subdivide()
	}

	if q.TopLeft.Insert(node) {
		return true
	}
	if q.TopRight.Insert(node) {
		return true
	}
	if q.BottomLeft.Insert(node) {
		return true
	}
	if q.BottomRight.Insert(node) {
		return true
	}

	return false
}

func (q *Quad) subdivide() {
	x := q.Boundary.Center.X
	y := q.Boundary.Center.Y
	newWidth := q.Boundary.Width / 2
	newHeight := q.Boundary.Height / 2

	q.TopLeft = NewQuad(&Rectangle{vec.Position{X: x - newWidth/2, Y: y - newHeight/2}, newWidth, newHeight}, q.Capacity)
	q.TopRight = NewQuad(&Rectangle{vec.Position{X: x + newWidth/2, Y: y - newHeight/2}, newWidth, newHeight}, q.Capacity)
	q.BottomLeft = NewQuad(&Rectangle{vec.Position{X: x - newWidth/2, Y: y + newHeight/2}, newWidth, newHeight}, q.Capacity)
	q.BottomRight = NewQuad(&Rectangle{vec.Position{X: x + newWidth/2, Y: y + newHeight/2}, newWidth, newHeight}, q.Capacity)
}
