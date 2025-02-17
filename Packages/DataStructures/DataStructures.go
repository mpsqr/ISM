package DataStructures

type Vector3 struct {
	X []float64
	Y []float64
	Z []float64
}

type List struct {
	X [][]int
}

type Array struct {
	X[]string
}


type FloatArr struct {
	X[]float64
}

func NewFloatArr(size int) FloatArr {
	return FloatArr{
		X: make([]float64, size),
	}
}


func NewArray(size int) Array {
	return Array{
		X: make([]string, size),
	};
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