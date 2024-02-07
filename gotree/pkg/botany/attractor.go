package botany

import (
	"image/color"

	"github.com/friendlymatthew/nature/pkg/vec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Attractor struct {
	position vec.Position
	reached  bool
}

func NewAttractor(x, y float32) *Attractor {

	p := vec.NewPosition(x, y)

	return &Attractor{
		position: p,
		reached:  false,
	}
}

func (a *Attractor) HasReached() bool {
	return a.reached
}

func (a *Attractor) Draw(screen *ebiten.Image) {
	clr := color.RGBA{R: 0, G: 255, B: 0, A: 255}
	vector.DrawFilledCircle(screen, a.position.X, a.position.Y, 2.0, clr, false)
}
