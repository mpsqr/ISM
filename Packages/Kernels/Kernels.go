package Kernels
import (/*"fmt"*/
	"ism/Packages/DataStructures"
	"ism/Packages/Maths"
	//"ism/Packages/Utilitary"
	"math/rand"
	/*"time"*/)


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

const T0 float64 = 300.0;

const deltaT float64 = 1.0;
const CONVERSION_FORCE float64 = 0.0001 * 4.186;
const mi float64 = 18.0; // Mass
const CONSTANTE_R float64 = 0.00199;


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




func KineticEnergy(p *DataStructures.Vector3, N int) float64 {
	var sum float64 = 0.0;
	for i := 0; i < N; i++ {
		sum += ((p.X[i] * p.X[i]) + (p.Y[i] * p.Y[i]) + (p.Z[i] * p.Z[i])) / mi;
	}

	//return 1.0 / (2 * CONVERSION_FORCE) * (sum / mi);
	return sum / (2 * CONVERSION_FORCE);
}

func KineticTemperature(cineticEnergy float64, N int) float64 {
	return 1.0 / (float64( ((3 * N) -3) ) * CONSTANTE_R) * cineticEnergy;
}


func CalibrateMoment(p *DataStructures.Vector3, N int) {
	var kineticEnergy float64 = KineticEnergy(p, N);

	var rapport float64 = (float64((3 * N) - 3) * CONSTANTE_R * T0) / kineticEnergy;

	for i := 0; i < N; i++ {
		p.X[i] *= rapport;
		p.Y[i] *= rapport;
		p.Z[i] *= rapport;
	}
}


func GenerateMoment(p *DataStructures.Vector3, N int) {
	rand.Seed(0);

	for i := 0; i < N; i++ {
		var cx float64 = rand.Float64();
		var cy float64 = rand.Float64();
		var cz float64 = rand.Float64();

		p.X[i] = Maths.Sign((rand.Float64() * 2) - 1) * cx;
		p.Y[i] = Maths.Sign((rand.Float64() * 2) - 1) * cy;
		p.Z[i] = Maths.Sign((rand.Float64() * 2) - 1) * cz;
	}

	CalibrateMoment(p, N);
}

func CenterOfMassCorrection(p *DataStructures.Vector3, N int) {
	var Px float64 = 0.0;
	var Py float64 = 0.0;
	var Pz float64 = 0.0;

	for i := 0; i < N; i++ {
		Px += p.X[i];
		Py += p.Y[i];
		Pz += p.Z[i];
	}

	Px /= float64(N);
	Py /= float64(N);
	Pz /= float64(N);

	for i := 0; i < N; i++ {
		p.X[i] -= Px;
		p.Y[i] -= Py;
		p.Z[i] -= Pz;
	}

	CalibrateMoment(p, N);
}

func VelocityVerlet(pos *DataStructures.Vector3, forces *DataStructures.Vector3, p *DataStructures.Vector3, N int) {
	for i := 0; i < N; i++ {
		p.X[i] += 0.5 * forces.X[i] * deltaT / mi;
		p.Y[i] += 0.5 * forces.Y[i] * deltaT / mi;
		p.Z[i] += 0.5 * forces.Z[i] * deltaT / mi;

		pos.X[i] += p.X[i] * deltaT / mi;
		pos.Y[i] += p.Y[i] * deltaT / mi;
		pos.Z[i] += p.Z[i] * deltaT / mi;
	}

	ComputeForcesPeriodic(pos, forces, N);

	for i := 0; i < N; i++ {
		p.X[i] += 0.5 * forces.X[i] * deltaT / mi;
		p.Y[i] += 0.5 * forces.Y[i] * deltaT / mi;
		p.Z[i] += 0.5 * forces.Z[i] * deltaT / mi;
	}

}

/*
func VelocityVerlet(pos *DataStructures.Vector3, forces *DataStructures.Vector3, p *DataStructures.Vector3, N int) {

	var newForces = DataStructures.NewVector3(N);
	var d2 float64 = deltaT*deltaT;

	for i := 0; i < N; i++ { // Update positions

		pos.X[i] += (vel.X[i] * deltaT) + ((0.5 * (forces.X[i] / mi)) * d2);
		pos.Y[i] += (vel.Y[i] * deltaT) + ((0.5 * (forces.Y[i] / mi)) * d2);
		pos.Z[i] += (vel.Z[i] * deltaT) + ((0.5 * (forces.Z[i] / mi)) * d2);
	}

	ComputeForcesPeriodic(pos, &newForces, N);

	for i := 0; i < N; i++ { // Update velocities
		vel.X[i] += (0.5 * (forces.X[i] + newForces.X[i]) / mi) * deltaT;
		vel.Y[i] += (0.5 * (forces.Y[i] + newForces.Y[i]) / mi) * deltaT;
		vel.Z[i] += (0.5 * (forces.Z[i] + newForces.Z[i]) / mi) * deltaT;

		forces.X[i] = newForces.X[i];
		forces.Y[i] = newForces.Y[i];
		forces.Z[i] = newForces.Z[i];

		
		p.X[i] = vel.X[i] / mi;
		p.Y[i] = vel.Y[i] / mi;
		p.Z[i] = vel.Z[i] / mi;
		
	}

	//Utilitary.CopyVec3(forces, &newForces);
}
*/