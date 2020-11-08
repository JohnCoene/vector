package vector

import (
	"fmt"
	"testing"
)

func TestRandomVector(t *testing.T) {
	p := RandomVector(1, 100)
	fmt.Printf("Random vector x: %f, y: %f", p.X, p.Y)
	RandomVectors(5, 1, 100)
}
