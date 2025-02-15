package DataStructures

type Vector3 struct {
	X []float64
	Y []float64
	Z []float64
}

type List struct {
	X [][]int
}

func NewVector3(size int) Vector3 { // Constructor
	return Vector3{
		X: make([]float64, size),
		Y: make([]float64, size),
		Z: make([]float64, size),
	};
}

func NewList(size int) List {
	return List{
		X: make([][]int, size),
	};
}