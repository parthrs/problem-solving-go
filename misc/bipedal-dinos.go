package misc

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

/*
You will be supplied with two data files in CSV format.
The first file contains # statistics about various dinosaurs.
The second file contains additional data.

Given the following formula:
speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g),
Where g = 9.8 m/s^2 (gravitational constant)

Write a program to read in the data files from disk, it must then print the names
of only the bipedal dinosaurs from fastest to slowest. Do not print any other information.
*/

type Dino struct {
	StrideLen float64
	LegLen    float64
	Speed     float64
	Name      string
	Diet      string
	Stance    string
}

func NewDino(name string) *Dino {
	return &Dino{
		Name: name,
	}
}

// Returns speed in m/s (because of the square root)
func (d *Dino) CalcSpeed() {
	d.Speed = ((d.StrideLen / d.LegLen) - 1) * math.Sqrt(d.LegLen*9.8)
}

// Prints bipedal dinos read from dataset1.csv and dataset2.csv
// on the basis of their speed. The print order is set by
// ascSort, which defaults to false meaning descinding sort
func PrintBipedalDinos(ascSort bool) {
	// Shall I use a struct or simply a dict?
	// Lets use struct - since we have to combine
	// fields - leg length, stride length -
	// to get speed. And it will be tricky to
	// to manipulate that in a simple dict only.
	// We'll need a dict, but it will hold name
	// to the struct type.
	dinoMap := map[string]*Dino{}

	f, err := os.Open("misc/dataset1.csv")
	if err != nil {
		log.Fatalf("Failed to open dataset1.csv")
	}

	csvReader := csv.NewReader(f)
	csvReader.Read() // Skip first line containing col headers
	// to read all in memory
	// dataSlice, err := csvReader.ReadAll(f)
	for {
		dataSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Unable to process line, skipping..")
			continue
		}
		legLength, err := strconv.ParseFloat(dataSlice[1], 64)
		if err != nil {
			log.Printf("Unable to convert leg length to float (%s), skipping..", dataSlice[1])
			continue
		}
		dino := NewDino(dataSlice[0])
		dino.LegLen, dino.Diet = legLength, dataSlice[2]
		dinoMap[dataSlice[0]] = dino
	}

	f, err = os.Open("misc/dataset2.csv")
	if err != nil {
		log.Fatalf("Failed to open dataset2.csv")
	}

	csvReader = csv.NewReader(f)
	csvReader.Read()
	for {
		dataSlice, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Unable to process line, skipping..")
			continue
		}
		strideLength, err := strconv.ParseFloat(dataSlice[1], 64)
		if err != nil {
			log.Printf("Unable to convert leg length to float (%s), skipping..", dataSlice[1])
			continue
		}
		var dino *Dino
		var found bool
		if dino, found = dinoMap[dataSlice[0]]; !found {
			dino = NewDino(dataSlice[0])
			dino.StrideLen, dino.Stance = strideLength, dataSlice[2]
			dinoMap[dataSlice[0]] = dino
		} else {
			dino.StrideLen, dino.Stance = strideLength, dataSlice[2]
			dino.CalcSpeed()
		}
	}

	dinoSlice := []string{}
	for k, v := range dinoMap {
		if v.Stance == "bipedal" && v.LegLen != 0.0 && v.StrideLen != 0.0 {
			dinoSlice = append(dinoSlice, k)
		}
	}

	sort.Slice(dinoSlice, func(i int, j int) bool {
		if ascSort {
			return dinoMap[dinoSlice[i]].Speed < dinoMap[dinoSlice[j]].Speed
		} else {
			return dinoMap[dinoSlice[i]].Speed > dinoMap[dinoSlice[j]].Speed
		}
	})

	for _, dino := range dinoSlice {
		fmt.Println(dinoMap[dino])
	}
}
