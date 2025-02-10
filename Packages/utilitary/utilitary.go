package Utilitary
import ("fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"ism/Packages/DataStructures")



func ImportXYZ(path string, pos *DataStructures.Vector3) {
	
	// Lecture du fichier
	file, err := os.Open(path);
	if err != nil {
		log.Fatal(err);
		return;
	}

	var i int = 0;
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		line := scanner.Text();

		parts := strings.Fields(line);
		if len(parts) == 4 { // Ligne valable

			x, errx := strconv.ParseFloat(parts[1], 64);
			y, erry := strconv.ParseFloat(parts[2], 64);
			z, errz := strconv.ParseFloat(parts[3], 64);

			if (errx != nil) || (erry != nil) || (errz != nil) {
				fmt.Printf("Error parsing line: %s\n", line);
				continue;
			}

			pos.X[i] = x;
			pos.Y[i] = y;
			pos.Z[i] = z;
			i++;
		}
	}

	file.Close();

}
/*
func ExportXYZ(path string, pos *DataStructures.Vector3) {

	file, err := os.Create(path);

	if err != nil {
		fmt.Println("Error creating the file.");
		return;
	}


	// Write comment
	_, err = fmt.Fprintf(file, "0 1\n");
	if err != nil {
		fmt.Println("Error writing to file.");
		return;
	}


	// Write atom data
	for i := 0; i < len(pos.X); i++ {
		_, err = fmt.Fprintf(file, "%d %.6f %.6f %.6f\n", 2, pos.X[i], pos.Y[i], pos.Z[i]);
		if err != nil {
			fmt.Println("Error writing to file.");
			return;
		}
	}

	file.Close();
}
*/

func ExportXYZ(path string, pos *DataStructures.Vector3, iter int, N int) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644);
	if err != nil {
		return err
	}
	defer file.Close();

	// Écrire l'en-tête CRYST1 et MODEL
	_, err = fmt.Fprintf(file, "CRYST1%9.3f%9.3f%9.3f  90.00  90.00  90.00 P             1\n", 42.0, 42.0, 42.0);
	if err != nil {
		return err;
	}
	_, err = fmt.Fprintf(file, "MODEL     %d\n", iter);
	if err != nil {
		return err;
	}

	// Écrire les coordonnées des particules
	for i := 0; i < N; i++ {
		_, err = fmt.Fprintf(file, "ATOM  %5d  C           0    %8.3f%8.3f%8.3f                 MRES\n", i, pos.X[i], pos.Y[i], pos.Z[i]);
		if err != nil {
			return err;
		}
	}

	// Écrire les lignes de fin
	_, err = fmt.Fprintf(file, "TER\nENDMDL\n");
	if err != nil {
		return err;
	}

	return nil;
}

func CopyVec3(to *DataStructures.Vector3, from *DataStructures.Vector3) {
	copy(to.X, from.X);
	copy(to.Y, from.Y);
	copy(to.Z, from.Z);
}

func IniVec3(vec *DataStructures.Vector3, val float64, N int) {
	for i := 0; i < N; i++ {
		vec.X[i] = val;
		vec.Y[i] = val;
		vec.Z[i] = val;
	}
}
/*
func ResetVec3(vec *DataStructures.Vector3, N int) {
	var zeroSlice = make([]float64, N);

	copy(vec.X, zeroSlice);
	copy(vec.Y, zeroSlice);
	copy(vec.Z, zeroSlice);
}
*/