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

// Add add another vector
func (v *Vector) Add(vect Vector) {
	v.X += vect.X
	v.Y += vect.Y
}

// Sub subtract another vector
func (v *Vector) Sub(vect Vector) {
	v.X -= vect.X
	v.Y -= vect.Y
}

// Mult multiply a vector by a float
func (v *Vector) Mult(x float64) {
	v.X *= x
	v.Y *= x
}

// Mag get the magnitude of a vector
func (v *Vector) Mag() float64 {
	h := math.Pow(v.X, 2) + math.Pow(v.Y, 2)
	return math.Sqrt(h)
}

// Norm normlaise a vector between 0 and 1
func (v *Vector) Norm() {
	h := v.Mag()
	v.X /= h
	v.Y /= h
}

// SetMag set the magnitude of a vector
func (v *Vector) SetMag(x float64) {
	v.Norm()
	v.Mult(x)
}

// Div divide by a fload
func (v *Vector) Div(x float64) {
	v.X /= x
	v.Y /= x
}
