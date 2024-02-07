package botany

import (
	"fmt"
	"math/rand"

	"github.com/friendlymatthew/nature/pkg/vec"
)

const (
	attractionDistance float32 = 60.0
	killDistance       float32 = 10.0
	marginWidth        float32 = 50
	marginHeightTop    float32 = 25
	marginHeightBottom float32 = 50
)

type Tree struct {

	// a fancy term for leaves
	attractors []Attractor
	nodes      []Node

	IsDone bool
}

func NewTree(attrCount int, sW, sH float32) *Tree {

	boundary := NewRectangle(vec.Position{X: 0, Y: 0}, vec.Position{X: sW, Y: sH})
	quadTree := NewQuad(&boundary, 4)

	attractors := make([]Attractor, attrCount)

	for i := range attractors {
		x := marginWidth + rand.Float32()*(sW-2*marginWidth)
		y := marginHeightTop + rand.Float32()*(sH-marginHeightTop-marginHeightBottom)
		attractors[i] = *NewAttractor(x, y)
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

	// initiation process
	// iterate through attractors, if there are no matches, then we grab the last node and increase vertically
	growTrunk := true

	lastNode := nodes[len(nodes)-1]

	for growTrunk {
		for _, attractor := range attractors {
			if lastNode.position.Dist(&attractor.position) < attractionDistance {
				growTrunk = false
			} else {
				newPos := vec.Position{X: lastNode.position.X, Y: lastNode.position.Y - 1}
				newNode := NewNode(&lastNode, newPos, vec.NewVec2(0, -1))
				nodes = append(nodes, newNode)
				lastNode = newNode
			}
		}
	}

	for _, node := range nodes {
		quadTree.Insert(node)
	}

	return &Tree{
		nodes:      nodes,
		attractors: attractors,
		quadTree:   *quadTree,

		IsDone: false,
	}
}

// http://algorithmicbotany.org/papers/colonization.egwnp2007.large.pdf
func (t *Tree) Grow() {
	if len(t.attractors) == 0 {
		fmt.Println("DONE")
		t.IsDone = true
		return
	}

	for i := range t.nodes {
		t.nodes[i].avgDir = vec.Vec2{}
		t.nodes[i].count = 0
		fmt.Printf("Node %d reset\n", i)
	}

	aCount := len(t.attractors)
	fmt.Println("Processing attractors...")
	for i := len(t.attractors) - 1; i >= 0; i-- {
		a := t.attractors[i]
		fmt.Printf("Processing attractor %d\n", i)
		closestNodeIndex := t.findClosestNode(a)
		fmt.Printf("Closest node to attractor %d is %d\n", i, closestNodeIndex)

		if closestNodeIndex == -1 {
			fmt.Printf("No close node found for attractor %d, skipping\n", i)
			continue
		}

		cNode := &t.nodes[closestNodeIndex]
		dist := cNode.position.Dist(&a.position)
		fmt.Printf("Distance to closest node: %v\n", dist)

		if dist < killDistance {
			fmt.Printf("Pruning leaf %d\n", i)
			t.attractors = append(t.attractors[:i], t.attractors[i+1:]...)
		} else if dist < attractionDistance {
			direction := a.position.Sub(&cNode.position)
			cNode.avgDir = cNode.avgDir.Add(&direction)
			cNode.count += 1
			fmt.Printf("Attractor %d summing, new count is %v\n", i, cNode.count)
		}
	}

	fmt.Printf("Nodes left: %v\n", len(t.nodes))
	if len(t.attractors)-aCount == 0 {
		fmt.Println("NO MORE!")
	}

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
		}
	}

}

// Helper function to find the closest node to an attractor within attraction distance
// todo -- implement a quadtree. Right now searching is O(n^2). WTS O(n*logn)
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
