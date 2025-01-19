package main
import("fmt"
	"ism/Packages/DataStructures"
	"ism/Packages/Utilitary"
	"ism/Packages/Maths")

/*
   Nom : Matéo Pasquier
   Email : mateo.pasquier@ens.uvsq.fr
   Cours : ISM
   Compilation : go build -o ./bin/main main.go
   Lancement : ./bin/main
*/


const N int = 1000;

func main() {

	fmt.Println("skibidi");


	// Allocation mémoire
	var pos DataStructures.Vector3; // Positions, SOA
	Utilitary.ImportXYZ("Input/particule.xyz", &pos);

	fmt.Println(pos.Z[1]);
	fmt.Println(Maths.SquaredDistance(pos.X[0], pos.Y[0], pos.Z[0], pos.X[1], pos.Y[1], pos.Z[1]));

}
