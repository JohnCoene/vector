package vector

import (
	"fmt"
	"testing"
)

func TestOp(t *testing.T) {
	v1 := Vector{X: 1, Y: 0}
	v2 := Vector{X: 1, Y: 1}

	fmt.Printf("original vector x: %f, y: %f\n", v1.X, v1.Y)

	v1.Add(v2)
	fmt.Printf("Added vector x: %f, y: %f\n", v1.X, v1.Y)

	v1.Sub(v2)
	fmt.Printf("Subtracted vector x: %f, y: %f\n", v1.X, v1.Y)

	v1.Mult(5)
	fmt.Printf("Vector multiplied by 5 x: %f, y: %f\n", v1.X, v1.Y)

	mag := v1.Mag()
	fmt.Printf("Vector magnitude  %f\n", mag)

	v1.Norm()
	mag2 := v1.Mag()
	fmt.Printf("Vector normalised x: %f, y: %f, mag: %f\n", v1.X, v1.Y, mag2)

	v1.SetMag(20)
	mag3 := v1.Mag()
	fmt.Printf("Vector set mag to 20, mag: %f\n", mag3)

	v1.Limit(5)
	fmt.Printf("Vector magnitude limited to 5, mag: %f\n", v1.Mag())
	v1.Limit(10)
	fmt.Printf("Vector magnitude limited to 10, mag: %f\n", v1.Mag())
}
