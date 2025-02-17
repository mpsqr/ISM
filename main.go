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
   Lancement : ./bin/main InputFile.xyz OutputFile.pdb iterations
*/


const N int = 1000;

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

	var types = DataStructures.NewArray(N);



	// Initialisation
	Utilitary.ImportXYZ(InputFile, &pos, &types);
	Kernels.BuildVerletLists(&pos, &VerletLists, N);

	var indB = Utilitary.FindInArray(&types, "B", N); // Fonction qui permet de connaître la position de l'unique particule B dans la liste, pour ne pas le recalculer à chaque fois


	fmt.Println("ULJ: ", Kernels.ComputeForces(&pos, &forces, N));
	fmt.Println("Somme des forces du système : ", Kernels.ComputeSumForces(&forces, N));

	fmt.Println("ULJ periodic: ", Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodic, N));


	var forcesA float64 = 0.0;
	var forcesB float64 = 0.0;
	forcesA, forcesB = Kernels.ComputeForcesPeriodicLists(&pos, &forcesPeriodicLists, &VerletLists, &forcesA, indB, N);
	fmt.Println("Moyenne des forces A : ", forcesA / float64(N-1));
	fmt.Println("Force B : ", forcesB);


	fmt.Println("\n\n---------------Début de la simulation---------------\n\n");

	var cineticEnergy float64 = 0.0;
	var cineticTemperature float64 = 0.0;

	Kernels.GenerateMoment(&moment, N);

	var start = time.Now();

	for i := 0; i < iters; i++ {
		fmt.Println("\n-----Itération", i, "-----\n");

		Kernels.VelocityVerletLists(&pos, &forcesPeriodic, &moment, &VerletLists, &forcesA, indB, N);


		if i % 150 == 0 {
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
