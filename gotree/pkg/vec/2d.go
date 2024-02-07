package vec

import "math"

// Position in the flat earth simulation
type Position struct {
	X float32
	Y float32
}

func NewPosition(x, y float32) Position {
	return Position{
		X: x,
		Y: y,
	}
}

// Euclidean Distance - assuming deltaX and deltaY are float32
func (p1 *Position) Dist(p2 *Position) float32 {
	deltaX := p1.X - p2.X
	deltaY := p1.Y - p2.Y
	return float32(math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY)))
}

// Sub performs vector subtraction
func (p1 *Position) Sub(p2 *Position) Vec2 {
	return Vec2{
		Dx: p1.X - p2.X,
		Dy: p1.Y - p2.Y,
	}
}

// Add
func (p *Position) Add(v *Vec2) Position {
	return Position{
		X: p.X + v.Dx,
		Y: p.Y + v.Dy,
	}
}

// Vec2 represents a 2D vector in the simulation
type Vec2 struct {
	Dx, Dy float32
}

func NewVec2(dx, dy float32) Vec2 {
	return Vec2{Dx: dx, Dy: dy}
}

// Euclidean Distance - assuming deltaX and deltaY are float32
func (v1 *Vec2) Dist(v2 *Vec2) float32 {
	deltaX := v1.Dx - v2.Dx
	deltaY := v1.Dy - v2.Dy
	return float32(math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY)))
}

// Vector Magnitude
func (v *Vec2) Mag() float32 {
	return float32(math.Sqrt(float64(v.Dx*v.Dx + v.Dy*v.Dy)))
}

// Add performs vector addition
func (v1 *Vec2) Add(v2 *Vec2) Vec2 {
	return Vec2{
		Dx: v1.Dx + v2.Dx,
		Dy: v1.Dy + v2.Dy,
	}
}

// Sub performs vector subtraction
func (v1 *Vec2) Sub(v2 *Vec2) Vec2 {
	return Vec2{
		Dx: v1.Dx - v2.Dx,
		Dy: v1.Dy - v2.Dy,
	}
}

// Norm normalizes the vector
func (v Vec2) Norm() Vec2 {
	mag := v.Mag()

	if mag > 0 {
		return Vec2{
			Dx: v.Dx / mag,
			Dy: v.Dy / mag,
		}
	}

	return Vec2{} // Zero values for float32
}

// Div performs scalar division on the vector
func (v *Vec2) Div(scalar float32) Vec2 {
	if scalar != 0 {
		return Vec2{Dx: v.Dx / scalar, Dy: v.Dy / scalar}
	}
	return Vec2{}
}
