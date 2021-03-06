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

func randomFloat(min, max float64) float64 {
	return min + r.Float64()*(max-min)
}

// RandomVector creates a random vector
func RandomVector(min, max float64) Vector {
	v := Vector{
		X: randomFloat(min, max),
		Y: randomFloat(min, max),
	}
	return v
}

// RandomVectors create n random vectors
func RandomVectors(n int, min, max float64) []Vector {
	var vector []Vector

	for i := 0; i < n; i++ {
		vector = append(vector, RandomVector(min, max))
	}

	return vector
}

// Add two vectors
func (v *Vector) Add(vect Vector) {
	v.X += vect.X
	v.Y += vect.Y
}

// Addf add a float to a vector
func (v *Vector) Addf(n float64) {
	v.X += n
	v.Y += n
}

// Sub subtract another vector
func (v *Vector) Sub(vect Vector) {
	v.X -= vect.X
	v.Y -= vect.Y
}

// Subf substract a float from vector
func (v *Vector) Subf(n float64) {
	v.X -= n
	v.Y -= n
}

// Mult multiply two vectors
func (v *Vector) Mult(vect Vector) {
	v.X *= vect.X
	v.Y *= vect.Y
}

// Multf multiply a vector by a float
func (v *Vector) Multf(n float64) {
	v.X *= n
	v.Y *= n
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
	v.Multf(x)
}

// Divf divide by a fload
func (v *Vector) Divf(n float64) {
	v.X /= n
	v.Y /= n
}

// Div divide by another vector
func (v *Vector) Div(vect Vector) {
	v.X /= vect.X
	v.Y /= vect.Y
}

// Limit defines a maximum magnitude for the vector
func (v *Vector) Limit(n float64) {
	mag := v.Mag()
	if mag > n {
		v.SetMag(n)
	}
}

// Sub subract one vector from another
func Sub(v1, v2 Vector) Vector {
	var results Vector
	results.X = v1.X - v2.X
	results.Y = v1.Y - v2.Y
	return results
}

// Add add one vector with another
func Add(v1, v2 Vector) Vector {
	var results Vector
	results.Add(v1)
	results.Add(v2)
	return results
}

// Mult multiply two vectors
func Mult(v1, v2 Vector) Vector {
	var results Vector
	results.X = v1.X * v2.X
	results.Y = v1.Y * v2.Y
	return results
}

// Addf add a float to a vector
func Addf(v Vector, n float64) Vector {
	results := Vector{X: v.X, Y: v.Y}
	results.X += n
	results.Y += n
	return results
}

// Subf subtract a float from a vector
func Subf(v Vector, n float64) Vector {
	results := Vector{X: v.X, Y: v.Y}
	results.X -= n
	results.Y -= n
	return results
}

// Multf multiply a vector by a float
func Multf(v Vector, n float64) Vector {
	results := Vector{X: v.X, Y: v.Y}
	results.X *= n
	results.Y *= n
	return results
}

// Divf divide by a float
func Divf(v Vector, n float64) Vector {
	results := Vector{X: v.X, Y: v.Y}
	results.X /= n
	results.Y /= n
	return results
}
