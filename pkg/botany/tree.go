package botany

import (
	"fmt"

	"github.com/friendlymatthew/nature/pkg/vec"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	attractionDistance float32 = 100.0
	killDistance       float32 = 40.0
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
			Y: sH,
		},
		vec.NewVec2(0, -1),
	)

	nodes := []Node{root}

	return &Tree{
		nodes:      nodes,
		attractors: attractors,
	}
}

// http://algorithmicbotany.org/papers/colonization.egwnp2007.large.pdf
func (t *Tree) Grow() {
	if len(t.attractors) == 0 {
		fmt.Println("END")
	}

	for i := range t.nodes {
		t.nodes[i].avgDir = vec.Vec2{}
		t.nodes[i].count = 0
	}

	for i, a := range t.attractors {
		closestNodeIndex := t.findClosestNode(a)

		if closestNodeIndex == -1 {
			lastNode := t.nodes[len(t.nodes)-1]
			newPos := vec.Position{X: lastNode.position.X + 1, Y: lastNode.position.Y - 1}
			newNode := NewNode(&lastNode, newPos, vec.NewVec2(1, -1))
			t.nodes = append(t.nodes, newNode)
			continue
		}

		cNode := &t.nodes[closestNodeIndex]
		if cNode.position.Dist(&a.position) < killDistance {
			// reached the end of the road
			fmt.Printf("\npruning leaf %v\n", i)
			t.attractors = append(t.attractors[:i], t.attractors[i+1:]...)
		} else if cNode.position.Dist(&a.position) < attractionDistance {
			direction := a.position.Sub(&cNode.position)
			cNode.avgDir = cNode.avgDir.Add(&direction)
			cNode.count += 1
			fmt.Printf("summing, new count is %v", cNode.count)
		}
	}

	fmt.Printf("\nnodes left: %v\n", len(t.nodes))

	// Create new nodes based on average direction
	for _, n := range t.nodes {
		if n.count > 0 {
			avgDir := n.avgDir.Div(float32(n.count)).Norm()
			newPos := n.position.Add(&avgDir)

			fmt.Printf("avgDir x: %v, y:%v", avgDir.Dx, avgDir.Dy)

			// create new node
			nn := NewNode(
				&n,
				newPos,
				avgDir, // we may need to scale
			)

			t.nodes = append(t.nodes, nn)
			fmt.Println("adding")
		} else {
			fmt.Println("count none")
		}
	}
}

// Helper function to find the closest node to an attractor within attraction distance
func (t *Tree) findClosestNode(attractor Attractor) int {
	var closestNodeIndex int = -1
	minDist := attractionDistance

	for i, node := range t.nodes {
		dist := node.position.Dist(&attractor.position)
		if dist < minDist {
			closestNodeIndex = i
			minDist = dist
		}
	}
	return closestNodeIndex
}

func (t *Tree) Draw(screen *ebiten.Image) {
	for _, a := range t.attractors {
		a.Draw(screen)
	}

	for _, n := range t.nodes {
		n.Draw(screen)
	}
}
