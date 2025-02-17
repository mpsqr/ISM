package Kernels
import ("ism/Packages/DataStructures"
		"ism/Packages/Maths"
		"math"
		"math/rand")


const (
    X = iota  // X = 0
    Y         // Y = 1
    Z         // Z = 2
)

const (
	A = iota
	B
)

const R 				float64 = 2.850;
const RSquared 			float64 = R * R;
const Epsilon 			float64 = 0.53;
const EpsilonLJ 		float64 = -48.0 * Epsilon;
const FloatCompensation float64 = 1e-12; // Afin d'éviter les pertes de précision lors de calculs avec des valeurs proches de 0


const EpsilonAB			float64 = 1.0;
const sigmaAB			float64 = 3.50;

const NSym 				int 	= 27;
const RCut 				float64 = 10.0;
const RCutSq 			float64 = RCut * RCut;
const RMax 				float64 = RCut * 1.2;
const RMaxSq 			float64 = RMax * RMax;
const L 				float64 = 32.0;

const T0 				float64 = 300.0;

const deltaT 			float64 = 1.0;
const CONVERSION_FORCE  float64 = 0.0001 * 4.186;
const mi 				float64 = 18.0; // Mass
const CONSTANTE_R 		float64 = 0.00199;

const GAMMA 			float64 = 0.01;


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

func ResetVec3(vec *DataStructures.Vector3, N int) {
	var zeroSlice = make([]float64, N);

	copy(vec.X, zeroSlice);
	copy(vec.Y, zeroSlice);
	copy(vec.Z, zeroSlice);
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
	ResetVec3(forces, N);

	for n := 0; n < NSym; n++ {
		for i := 0; i < N; i++ {
			for j := i+1; j < N; j++ {

				var xj float64 = pos.X[j] + translate[n][X];
				var yj float64 = pos.Y[j] + translate[n][Y];
				var zj float64 = pos.Z[j] + translate[n][Z];

				
				var dist float64 = Maths.SquaredDistance(pos.X[i], pos.Y[i], pos.Z[i], xj, yj, zj);

				if dist < RCutSq /*|| dist > 1e-6*/ {
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


func ComputeForcesPeriodicLists(pos *DataStructures.Vector3, forces *DataStructures.Vector3, list *DataStructures.List, forcesA *float64, indB int, N int) (float64, float64) {
	
	var energy float64 = 0.0;
	var energyB float64 = 0.0;
	ResetVec3(forces, N);

	for n := 0; n < NSym; n++ {
		for i := 0; i < N; i++ {
			for j := 0; j < len(list.X[i]); j++ {

				var ind int = list.X[i][j];

				var xj float64 = pos.X[ind] + translate[n][X];
				var yj float64 = pos.Y[ind] + translate[n][Y];
				var zj float64 = pos.Z[ind] + translate[n][Z];

				if ind < i {
					continue;
				}

					
				var dist float64 = Maths.SquaredDistance(pos.X[i], pos.Y[i], pos.Z[i], xj, yj, zj);

				if dist < RCutSq /*|| dist > 1e-6*/ {
					// Optimisation des calculs des puissances

					if i == indB || ind == indB {
						var r1 float64 = sigmaAB / (dist + FloatCompensation);
						var r2 = r1 * r1;
						var r4 = r2 * r2;
						var r5 = r4 * r1;
						var r8 = r4 * r4;
						var r12 = r8 * r4;
						var r14 = r12 * r2;
						var r15 = r14 * r1;

						var localForce = (-15.0 * EpsilonAB) * (r15 - r5);

						var forceX = localForce * (pos.X[i] - xj);
						var forceY = localForce * (pos.Y[i] - yj);
						var forceZ = localForce * (pos.Z[i] - zj);
								
						// Mise à jour des forces
						forces.X[i] += forceX;
						forces.Y[i] += forceY;
						forces.Z[i] += forceZ;
								
						forces.X[ind] -= forceX;
						forces.Y[ind] -= forceY;
						forces.Z[ind] -= forceZ;
								



						energyB += r15 - (r5 + r5 + r5);


						continue;
					}

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
							
					forces.X[ind] -= forceX;
					forces.Y[ind] -= forceY;
					forces.Z[ind] -= forceZ;
							


					*forcesA += r12 - (r6 + r6);
					energy += r12 - (r6 + r6);
				}
			}
		}
	}

	return (energy * Epsilon) * 4.0, (energyB * EpsilonAB);
}



func ComputeSumForces(forces *DataStructures.Vector3, N int) float64 {
	return Maths.Vec3Sum(forces, N);
}




func KineticEnergy(p *DataStructures.Vector3, N int) float64 {
	var sum float64 = 0.0;
	for i := 0; i < N; i++ {
		sum += ((p.X[i] * p.X[i]) + (p.Y[i] * p.Y[i]) + (p.Z[i] * p.Z[i])) / mi;
	}

	return sum / (2 * CONVERSION_FORCE);
}

func KineticTemperature(cineticEnergy float64, N int) float64 {
	return 1.0 / (float64( ((3 * N) -3) ) * CONSTANTE_R) * cineticEnergy;
}


func CalibrateMoment(p *DataStructures.Vector3, N int) {
	var kineticEnergy float64 = KineticEnergy(p, N);

	var rapport float64 = (float64((3 * N) - 3) * CONSTANTE_R * T0) / kineticEnergy;

	rapport = math.Sqrt(rapport);

	for i := 0; i < N; i++ {
		p.X[i] *= rapport;
		p.Y[i] *= rapport;
		p.Z[i] *= rapport;
	}
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
	CenterOfMassCorrection(p, N);
	CalibrateMoment(p, N);
}


func BerendsenCorrection(p *DataStructures.Vector3, N int) {
	var cineticEnergy float64 = KineticEnergy(p, N);
	//var ratio float64 = T0 / KineticTemperature(cineticEnergy, N);
	var ratio float64 = GAMMA * (T0 / KineticTemperature(cineticEnergy, N) - 1.0);
	/*
	for i := 0; i < N; i++ {
		p.X[i] = p.X[i] + (GAMMA * (ratio - 1.0) * p.X[i]);
		p.Y[i] = p.Y[i] + (GAMMA * (ratio - 1.0) * p.Y[i]);
		p.Z[i] = p.Z[i] + (GAMMA * (ratio - 1.0) * p.Z[i]);
	}
	*/
	for i := 0; i < N; i++ {
		p.X[i] = p.X[i] * ratio + p.X[i];
		p.Y[i] = p.Y[i] * ratio + p.Y[i];;
		p.Z[i] = p.Z[i] * ratio + p.Z[i];;
	}

}

func VelocityVerlet(pos *DataStructures.Vector3, forces *DataStructures.Vector3, p *DataStructures.Vector3, N int) {

	var hdt float64 = deltaT * 0.5;

	for i := 0; i < N; i++ {
		p.X[i] -= hdt * (forces.X[i] * CONVERSION_FORCE);
		p.Y[i] -= hdt * (forces.Y[i] * CONVERSION_FORCE);
		p.Z[i] -= hdt * (forces.Z[i] * CONVERSION_FORCE);

		pos.X[i] += p.X[i] * deltaT / mi;
		pos.Y[i] += p.Y[i] * deltaT / mi;
		pos.Z[i] += p.Z[i] * deltaT / mi;
	}
	
	for i := 0; i < N; i++ {
		pos.X[i] -= math.Round((pos.X[i] / L) + 0.5) * L;
		pos.Y[i] -= math.Round((pos.Y[i] / L) + 0.5) * L;
		pos.Z[i] -= math.Round((pos.Z[i] / L) + 0.5) * L;
	}

	ComputeForcesPeriodic(pos, forces, N);

	for i := 0; i < N; i++ {
		p.X[i] -= hdt * (forces.X[i] * CONVERSION_FORCE);
		p.Y[i] -= hdt * (forces.Y[i] * CONVERSION_FORCE);
		p.Z[i] -= hdt * (forces.Z[i] * CONVERSION_FORCE);
	}

}

func VelocityVerletLists(pos *DataStructures.Vector3, forces *DataStructures.Vector3, p *DataStructures.Vector3, list *DataStructures.List, forcesA *float64, indB int, N int) {

	var hdt float64 = deltaT * 0.5;

	for i := 0; i < N; i++ {
		p.X[i] -= hdt * (forces.X[i] * CONVERSION_FORCE);
		p.Y[i] -= hdt * (forces.Y[i] * CONVERSION_FORCE);
		p.Z[i] -= hdt * (forces.Z[i] * CONVERSION_FORCE);

		pos.X[i] += p.X[i] * deltaT / mi;
		pos.Y[i] += p.Y[i] * deltaT / mi;
		pos.Z[i] += p.Z[i] * deltaT / mi;
	}
	
	for i := 0; i < N; i++ {
		pos.X[i] -= math.Round((pos.X[i] / L) + 0.5) * L;
		pos.Y[i] -= math.Round((pos.Y[i] / L) + 0.5) * L;
		pos.Z[i] -= math.Round((pos.Z[i] / L) + 0.5) * L;
	}


	ComputeForcesPeriodicLists(pos, forces, list, forcesA, indB, N);

	for i := 0; i < N; i++ {
		p.X[i] -= hdt * (forces.X[i] * CONVERSION_FORCE);
		p.Y[i] -= hdt * (forces.Y[i] * CONVERSION_FORCE);
		p.Z[i] -= hdt * (forces.Z[i] * CONVERSION_FORCE);
	}

}


func BuildVerletLists(pos *DataStructures.Vector3, list *DataStructures.List, N int) {
	for i := 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			for n := 0; n < NSym; n++ {

				var xj float64 = pos.X[j] + translate[n][X];
				var yj float64 = pos.Y[j] + translate[n][Y];
				var zj float64 = pos.Z[j] + translate[n][Z];

				
				var dist float64 = Maths.SquaredDistance(pos.X[i], pos.Y[i], pos.Z[i], xj, yj, zj);

				if dist < RMaxSq {
					list.X[i] = append(list.X[i], j);
					list.X[j] = append(list.X[j], i); // Optimization
					break;
				}
			}
		}
	}
}