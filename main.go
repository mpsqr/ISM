package main
import("fmt"
	"time"
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

	if len(os.Args) < 4 {
		fmt.Println("Usage : <Input file> <Output FIle> <iterations>");
		return;
	}

	iters, err := strconv.Atoi(os.Args[3]);
	if err != nil {
		fmt.Println("Error: Invalid number format")
		return
	}

	var InputFile string = os.Args[1];
	var outputFile string = os.Args[2];


	// Allocation
	var pos = DataStructures.NewVector3(N); // Positions, SOA
	var forces = DataStructures.NewVector3(N); // Forces
	var forcesPeriodic = DataStructures.NewVector3(N); // Periodic forces
	var forcesPeriodicLists = DataStructures.NewVector3(N); // Periodic forces with Verlet Lists
	var moment = DataStructures.NewVector3(N); // Moment
	var VerletLists = DataStructures.NewList(N);



	// Initialisation
	Utilitary.ImportXYZ(InputFile, &pos);
	Kernels.BuildVerletLists(&pos, &VerletLists, N);



	fmt.Println("ULJ: ", Kernels.ComputeForces(&pos, &forces, N));
	fmt.Println("Somme des forces du système : ", Kernels.ComputeSumForces(&forces, N));

	fmt.Println("ULJ periodic: ", Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodic, N));

	fmt.Println("ULJ periodic with Verlet Lists: ", Kernels.ComputeForcesPeriodicLists(&pos, &forcesPeriodicLists, &VerletLists, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodicLists, N));


	fmt.Println("\n\n---------------Début de la simulation---------------\n\n");

	var cineticEnergy float64 = 0.0;
	var cineticTemperature float64 = 0.0;

	Kernels.GenerateMoment(&moment, N);

	var start = time.Now();

	for i := 0; i < iters; i++ {
		fmt.Println("\n-----Itération", i, "-----\n");

		//Kernels.VelocityVerlet(&pos, &forcesPeriodic, &moment, N);
		Kernels.VelocityVerletLists(&pos, &forcesPeriodic, &moment, &VerletLists, N);




		if i % 20 == 0 {
			Kernels.BerendsenCorrection(&moment, N);

			VerletLists = DataStructures.NewList(N);
			Kernels.BuildVerletLists(&pos, &VerletLists, N);
		}




		Utilitary.ExportXYZ(outputFile, &pos, i, N);
	}



	cineticEnergy = Kernels.KineticEnergy(&moment, N);
	cineticTemperature = Kernels.KineticTemperature(cineticEnergy, N);
	fmt.Println("Énergie cinétique : ", cineticEnergy);
	fmt.Println("Température cinétique : ", cineticTemperature);

	fmt.Println(time.Since(start));
}
