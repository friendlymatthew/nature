package botany

import (
	"image/color"
	"math/rand"

	"github.com/friendlymatthew/nature/pkg/vec"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Attractor struct {
	position vec.Position
	reached  bool
}

func NewAttractor(sW, sH float32) *Attractor {
	marginWidth := float32(25)
	marginHeightTop := float32(25)
	marginHeightBottom := float32(50)

	x := marginWidth + rand.Float32()*(sW-2*marginWidth)
	y := marginHeightTop + rand.Float32()*(sH-marginHeightTop-marginHeightBottom)

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
