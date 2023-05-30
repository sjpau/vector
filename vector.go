package vector

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v *Vector2D) Dot(v2 *Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2D) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v *Vector2D) Angle(v2 *Vector2D) float64 {
	return math.Acos(v.Dot(v2) / (v.Magnitude() * v2.Magnitude()))
}

func (v *Vector2D) Normalize() {
	length := v.Magnitude()
	v.X /= length
	v.Y /= length
}

func (v *Vector2D) Add(v2 *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v *Vector2D) Sub(v2 *Vector2D) *Vector2D {
	return &Vector2D{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}
