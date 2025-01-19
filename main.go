package main
import("fmt"
	"ism/Packages/DataStructures"
	"ism/Packages/utilitary")

/*
   Nom : Matéo Pasquier
   Email : mateo.pasquier@ens.uvsq.fr
   Cours : ISM
   Compilation : go build main.go
   Lancement : ./main
*/


const N int = 1000;

func main() {

	fmt.Println("skibidi");


	// Allocation mémoire
	var pos data_structures.Vector3; // Positions, SOA
	utilitary.ImportXYZ("Input/particule.xyz", &pos);

	fmt.Println(pos.Z[1]);

}
