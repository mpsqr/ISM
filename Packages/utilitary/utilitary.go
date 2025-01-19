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

	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		line := scanner.Text();

		parts := strings.Fields(line);
		if len(parts) != 4 { // Ligne non valable
			continue;
		}

		x, errx := strconv.ParseFloat(parts[1], 64);
		y, erry := strconv.ParseFloat(parts[2], 64);
		z, errz := strconv.ParseFloat(parts[3], 64);

		if (errx != nil) || (erry != nil) || (errz != nil) {
			fmt.Printf("Error parsing line: %s\n", line);
			continue;
		}

		pos.X = append(pos.X, x);
		pos.Y = append(pos.Y, y);
		pos.Z = append(pos.Z, z);
	}

	file.Close();

}
