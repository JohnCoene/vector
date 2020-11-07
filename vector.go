package vector

import (
	"math"
	"math/rand"
	"time"
)

// Vector defines a vector with x, y, and and id
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

var seed = rand.NewSource(time.Now().UnixNano())
var r = rand.New(seed)

// RandomVector creates a random vector
func RandomVector() Vector {
	v := Vector{
		X: r.Float64(),
		Y: r.Float64(),
	}
	return v
}

// RandomVectors create n random vectors
func RandomVectors(n int) []Vector {
	var vector []Vector

	for i := 0; i < n; i++ {
		vector = append(vector, RandomVector())
	}

	return vector
}

func (v *Vector) add(vect Vector) {
	v.X += vect.X
	v.Y += vect.Y
}

func (v *Vector) sub(vect Vector) {
	v.X -= vect.X
	v.Y -= vect.Y
}

func (v *Vector) mult(x float64) {
	v.X *= x
	v.Y *= x
}

func (v *Vector) mag() float64 {
	h := math.Pow(v.X, 2) + math.Pow(v.Y, 2)
	return math.Sqrt(h)
}

func (v *Vector) norm() {
	h := v.mag()
	v.X /= h
	v.Y /= h
}

func (v *Vector) setMag(x float64) {
	v.norm()
	v.mult(x)
}
