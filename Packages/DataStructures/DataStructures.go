package DataStructures

type Vector3 struct {
	X []float64
	Y []float64
	Z []float64
}

func NewVector3(size int) Vector3 { // Constructor
	return Vector3{
		X: make([]float64, size),
		Y: make([]float64, size),
		Z: make([]float64, size),
	}
}