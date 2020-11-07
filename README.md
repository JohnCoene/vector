# vector

A package for basic 2d vector operations.

```go
v1 := Vector{X: 1, Y: 0}
v2 := Vector{X: 1, Y: 1}

fmt.Printf("original vector x: %f, y: %f\n", v1.X, v1.Y)

v1.add(v2)
fmt.Printf("Added vector x: %f, y: %f\n", v1.X, v1.Y)

v1.sub(v2)
fmt.Printf("Subtracted vector x: %f, y: %f\n", v1.X, v1.Y)

v1.mult(5)
fmt.Printf("Vector multiplied by 5 x: %f, y: %f\n", v1.X, v1.Y)

mag := v1.mag()
fmt.Printf("Vector magnitude  %f\n", mag)

v1.norm()
mag2 := v1.mag()
fmt.Printf("Vector normalised x: %f, y: %f, mag: %f\n", v1.X, v1.Y, mag2)

v1.setMag(20)
mag3 := v1.mag()
fmt.Printf("Vector set mag to 20, mag: %f\n", mag3)
```

