package main
import("fmt"
		"ism/DS"
		"ism/utilitary")

/*
   Name: Matéo Pasquier
   Email: mateo.pasquier@ens.uvsq.fr
   Course Section: ISM
   Instructions to build the program: go build main.go
   Instructions to run the program: ./main
*/


const N int = 1000;

func main() {

	fmt.Println("skibidi");




	// Memory allocation
	var pos data_structures.Vector3;

	utilitary.ImportXYZ("Input/particule.xyz", &pos);

	fmt.Println(pos.Z[1]);

}