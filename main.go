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
	var moment = DataStructures.NewVector3(N); // Moment



	// Initialisation
	Utilitary.ImportXYZ(InputFile, &pos);


	fmt.Println("ULJ: ", Kernels.ComputeForces(&pos, &forces, N));
	fmt.Println("Somme des forces du système : ", Kernels.ComputeSumForces(&forces, N));

	fmt.Println("ULJ periodic: ", Kernels.ComputeForcesPeriodic(&pos, &forcesPeriodic, N));
	fmt.Println("Somme des forces du système périodique : ", Kernels.ComputeSumForces(&forcesPeriodic, N));

	/*
	Kernels.GenerateMoment(&moment, N);
	Kernels.CenterOfMassCorrection(&moment, N);
	*/


	fmt.Println("\n\n---------------Début de la simulation---------------\n\n");

	var cineticEnergy float64 = 0.0;
	var cineticTemperature float64 = 0.0;

	Kernels.GenerateMoment(&moment, N);
	//Utilitary.ExportXYZ("Results/pos" + ".pdb", &pos, 0, N);

	for i := 0; i < iters; i++ {
		fmt.Println("\n-----Itération", i, "-----\n");

		Kernels.VelocityVerlet(&pos, &forcesPeriodic, &moment, N);


		cineticEnergy = Kernels.KineticEnergy(&moment, N);
		cineticTemperature = Kernels.KineticTemperature(cineticEnergy, N);

		//Kernels.CalibrateMoment(&moment, N);

		if i % 20 == 0 {
			Kernels.BerendsenCorrection(&moment, N);
		}



		fmt.Println("Énergie cinétique : ", cineticEnergy);
		fmt.Println("Température cinétique : ", cineticTemperature);


		Utilitary.ExportXYZ(outputFile + ".pdb", &pos, i, N);
		//Utilitary.ExportXYZ("Results/mom" + ".pdb", &moment, i, N);
	}

}
