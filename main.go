package main
import("fmt"
		"os"
		"log"
		"ism/Packages")

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



	// Read the file
	file, err := os.Open("Input/particule.xyz");
	if err != nil {
		log.Fatal(err);
	}


	// Memory allocation
	var pos data_structures.Vector3;
	pos.X = append(pos.X, 69.0);

	fmt.Println(pos.X[0]);


	file.Close();
}