package main
import("fmt"
	"os"
	"strconv"
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

func main() {

	var iters int = 0;

	if len(os.Args) < 2 {
		fmt.Println("Usage : <iterations>");
		return;
	}

	iters, err := strconv.Atoi(os.Args[1]);
	if err != nil {
		fmt.Println("Error: Invalid number format")
		return
	}


	// Allocation
	var pos = DataStructures.NewVector3(N); // Positions, SOA
	var forces = DataStructures.NewVector3(N); // Forces
	var forcesPeriodic = DataStructures.NewVector3(N); // Periodic forces
	var angularMomentum = DataStructures.NewVector3(N); // Moment



	// Initialisation
	Utilitary.ImportXYZ("Input/particule.xyz", &pos);


	fmt.Println("ULJ: ", Kernels.ComputeForces(&pos, &forces, N));
	fmt.Println("Somme des forces du système : ", Kernels.ComputeSumForces(&forces, N));

	fmt.Println("ULJ periodic: ", Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodic, N));

	/*
	Kernels.GenerateMoment(&angularMomentum, N);
	Kernels.CenterOfMassCorrection(&angularMomentum, N);
	*/


	fmt.Println("\n\n---------------Début de la simulation---------------\n\n");

	var cineticEnergy float64 = 0.0;
	var cineticTemperature float64 = 0.0;
	var U float64 = 0.0;

	Kernels.GenerateMoment(&angularMomentum, N);

	for i := 0; i < iters; i++ {
		fmt.Println("\n-----Itération", i, "-----\n");

		Kernels.VelocityVerlet(&pos, &forcesPeriodic, &angularMomentum, N);

		Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N);
		U = Kernels.ComputeSumForces(&forcesPeriodic, N);
		cineticEnergy = Kernels.KineticEnergy(&angularMomentum, N);
		cineticTemperature = Kernels.KineticTemperature(cineticEnergy, N);

		fmt.Println("Énergie cinétique : ", cineticEnergy);
		fmt.Println("Somme des forces périodiques : ", U);
		fmt.Println("Température cinétique : ", cineticTemperature);
		fmt.Println("Énergie totale : ", U + cineticEnergy);

		Utilitary.ExportXYZ("Results/res" + ".pdb", &pos, i, N);
	}

}
