package vector

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (self *Vector2D) Dot(v *Vector2D) float64 {
	return self.X*v.X + self.Y*v.Y
}

func (self *Vector2D) Magnitude() float64 {
	return math.Sqrt(self.X*self.X + self.Y*self.Y)
}

func (self *Vector2D) Angle(v *Vector2D) float64 {
	return math.Cos(self.Dot(v) / (self.Magnitude() * v.Magnitude()))
}
