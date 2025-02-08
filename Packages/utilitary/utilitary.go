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

func resetVec3(vec *DataStructures.Vector3, N int) {
	var zeroSlice = make([]float64, N);

	copy(vec.X, zeroSlice);
	copy(vec.Y, zeroSlice);
	copy(vec.Z, zeroSlice);
}