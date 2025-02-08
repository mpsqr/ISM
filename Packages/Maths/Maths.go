package Maths
import ("ism/Packages/DataStructures")

// Return the squared distance between two particules
func SquaredDistance(x1, y1, z1, x2, y2, z2 float64) float64 {
	return ((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2));
}


func Vec3Sum(vec *DataStructures.Vector3, N int) float64 {
	var sum float64 = 0.0;
	for i := 0; i < N; i++ {
		sum += vec.X[i];
		sum += vec.Y[i];
		sum += vec.Z[i];
	}
	return sum;
}

func Sign(val float64) float64 {
	if val < 0.0 {
		return -1.0;
	}
	return 1.0;
}