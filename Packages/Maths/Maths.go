package Maths

// Return the squared distance between two particules
func SquaredDistance(x1, y1, z1, x2, y2, z2 float64) float64 {
	return ((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2));
}
