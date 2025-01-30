package Kernels
import (/*"fmt"*/
	"ism/Packages/DataStructures"
	"ism/Packages/Maths")


const (
    X = iota  // X = 0
    Y         // Y = 1
    Z         // Z = 2
)

const R float64 = 3.0;
const RSquared float64 = R * R;
const Epsilon float64 = 0.2;
const EpsilonLJ float64 = -48.0 * Epsilon;
const FloatCompensation = 1e-11; // Afin d'éviter les pertes de précision lors de calculs avec des valeurs proches de 0


const NSym int = 27;
const RCut float64 = 10.0;
const RCutSq float64 = RCut * RCut;
const L float64 = 42.0;


const dt float64 = 1.0;
const CONVERSION_FORCE float64 = 0.0001 * 4.186;
const mi float64 = 18.0; // Mass
const CONSTANTE_R = 0.00199;

var translate = [27][3]float64{
	 {0.0,   0.0,   0.0},
	 {0.0,   0.0,    L },
	 {0.0,   0.0,   -L },
	 {0.0,    L ,   0.0},
	 {0.0,    L ,    L },
	 {0.0,    L ,   -L },
	 {0.0,   -L ,   0.0},
	 {0.0,   -L ,    L },
	 {0.0,   -L ,   -L },
	 { L ,   0.0,   0.0},
	 { L ,   0.0,    L },
	 { L ,   0.0,   -L },
	 { L ,    L ,   0.0},
	 { L ,    L ,    L },
	 { L ,    L ,   -L },
	 { L ,   -L ,   0.0},
	 { L ,   -L ,    L },
	 { L ,   -L ,   -L },
	 {-L ,   0.0,   0.0},
	 {-L ,   0.0,    L },
	 {-L ,   0.0,   -L },
	 {-L ,    L ,   0.0},
	 {-L ,    L ,    L },
	 {-L ,    L ,   -L },
	 {-L ,   -L ,   0.0},
	 {-L ,   -L ,    L },
	 {-L ,   -L ,   -L },
}


// Calcule les forces entre les particules
func ComputeForces(pos *DataStructures.Vector3, forces *DataStructures.Vector3, N int) float64 {
	
	var energy float64 = 0.0;
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			// Optimisation des calculs des puissances
			var r2 float64 = RSquared / (Maths.SquaredDistance(pos.X[i], pos.Y[i], pos.Z[i], pos.X[j], pos.Y[j], pos.Z[j]) + FloatCompensation);
			var r4 float64 = r2 * r2;
			var r6 float64 = r4 * r2;
			var r8 float64 = r4 * r4;
			var r12 float64 = r8 * r4;
			var r14 float64 = r12 * r2;

			var localForce = EpsilonLJ * (r14 - r8);

			var forceX = localForce * (pos.X[i] - pos.X[j]);
			var forceY = localForce * (pos.Y[i] - pos.Y[j]);
			var forceZ = localForce * (pos.Z[i] - pos.Z[j]);
			
			// Mise à jour des forces
			forces.X[i] += forceX;
			forces.Y[i] += forceY;
			forces.Z[i] += forceZ;
		
			forces.X[j] -= forceX;
			forces.Y[j] -= forceY;
			forces.Z[j] -= forceZ;

			energy += r12 - (r6 + r6);
		}
	}

	return (energy * Epsilon) * 4.0;
}

func ComputeForcesPeriodic(pos *DataStructures.Vector3, forces *DataStructures.Vector3, N int) float64 {
	
	var energy float64 = 0.0;

	for n := 0; n < NSym; n++ {
		for i := 0; i < N; i++ {
			for j := i+1; j < N; j++ {

				var xj float64 = pos.X[j] + translate[n][X];
				var yj float64 = pos.Y[j] + translate[n][Y];
				var zj float64 = pos.Z[j] + translate[n][Z];

					
				var dist float64 = Maths.SquaredDistance(pos.X[i], pos.Y[i], pos.Z[i], xj, yj, zj);

				if dist < RCutSq {
					// Optimisation des calculs des puissances
					var r2 float64 = RSquared / (dist + FloatCompensation);
					var r4 float64 = r2 * r2;
					var r6 float64 = r4 * r2;
					var r8 float64 = r4 * r4;
					var r12 float64 = r8 * r4;
					var r14 float64 = r12 * r2;

					var localForce = EpsilonLJ * (r14 - r8);

					var forceX = localForce * (pos.X[i] - xj);
					var forceY = localForce * (pos.Y[i] - yj);
					var forceZ = localForce * (pos.Z[i] - zj);
						
					// Mise à jour des forces
					forces.X[i] += forceX;
					forces.Y[i] += forceY;
					forces.Z[i] += forceZ;
						
					forces.X[j] -= forceX;
					forces.Y[j] -= forceY;
					forces.Z[j] -= forceZ;
						



					energy += r12 - (r6 + r6);
				}
			}
		}
	}

	return (energy * Epsilon) * 4.0;
}

func ComputeSumForces(forces *DataStructures.Vector3, N int) float64 {
	return Maths.Vec3Sum(forces, N);
}



func VelocityVerlet(pos *DataStructures.Vector3, vel *DataStructures.Vector3, forces *DataStructures.Vector3, N int) {
	
}