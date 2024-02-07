package vec

import (
	"math"
	"testing"
)

func TestVec2(t *testing.T) {

	t.Run("Add", func(t *testing.T) {
		v1 := Vec2{Dx: 1, Dy: 2}
		v2 := Vec2{Dx: -3, Dy: 4}
		result := v1.Add(&v2)
		expected := Vec2{Dx: -2, Dy: 6}

		if result != expected {
			t.Errorf("Add was incorrect, got: %v, want: %v.", result, expected)
		}
	})

	t.Run("Dist", func(t *testing.T) {
		p1 := Position{X: 0, Y: 0}
		p2 := Position{X: 3, Y: 4}
		result := p1.Dist(&p2)
		expected := float32(5.0)

		if result != expected {
			t.Errorf("Dist was incorrect, got: %v, want: %v.", result, expected)
		}
	})

	t.Run("Mag", func(t *testing.T) {
		v := Vec2{Dx: 3, Dy: 4}
		result := v.Mag()
		expected := float32(5.0)

		if result != expected {
			t.Errorf("Magnitude was incorrect, got: %v, want: %v.", result, expected)
		}
	})

	t.Run("Norm", func(t *testing.T) {
		v := Vec2{Dx: 3, Dy: 4}
		result := v.Norm()
		expected := Vec2{Dx: 0.6, Dy: 0.8}

		if math.Abs(float64(result.Dx-expected.Dx)) > 1e-9 || math.Abs(float64(result.Dy-expected.Dy)) > 1e-9 {
			t.Errorf("Normalize was incorrect, got: %v, want: %v.", result, expected)
		}
	})
}
