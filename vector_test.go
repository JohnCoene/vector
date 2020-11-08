package vector

import (
	"testing"
)

func TestRandomVector(t *testing.T) {
	RandomVector(1, 100)
	RandomVectors(5, 1, 100)
}
