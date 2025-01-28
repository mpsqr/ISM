package main
import("fmt"
	"ism/Packages/DataStructures"
	"ism/Packages/Utilitary"
	"ism/Packages/Kernels")

/*
   Nom : Matéo Pasquier
   Email : mateo.pasquier@ens.uvsq.fr
   Cours : ISM
   Compilation : go build -o ./bin/main main.go
   Lancement : ./bin/main
*/


const N int = 1000;
const NLocal int = 100;

const NSym int = 27;

func main() {

	//fmt.Println("skibidi");


	// Allocation
	var pos DataStructures.Vector3; // Positions, SOA
	var forces DataStructures.Vector3; // Forces
	var forcesPeriodic DataStructures.Vector3; // Periodic forces


	// Initialisation
	Utilitary.ImportXYZ("Input/particule.xyz", &pos);
	Utilitary.IniVec3(&forces, 0.0, N);
	Utilitary.IniVec3(&forcesPeriodic, 0.0, N);


	fmt.Println("ULJ: ", Kernels.ComputeForces(&pos, &forces, N));
	fmt.Println("Somme des forces du système : ", Kernels.ComputeSumForces(&forces, N));

	fmt.Println("ULJ periodic: ", Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodic, N));
}
