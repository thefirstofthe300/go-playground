package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	f1, err := os.Open("file1.csv")
	if err != nil {
		log.Fatalf("Unable to open file1.csv: %s", err)
	}
	f2, err := os.Open("file2.csv")
	if err != nil {
		log.Fatalf("Unable to open file2.csv: %s", err)
	}
	dinos, err := readFile1(f1)
	if err != nil {
		log.Fatalf("Unable to load data from file1: %s", err)
	}
	dinos, err = readFile2(f2, dinos)
	if err != nil {
		log.Fatalf("Unable to load data from file2: %s", err)
	}
	sort.Sort(sort.Reverse(dinos))
	printDinos(dinos)
}

type Dinos []*Dino

func (d Dinos) Len() int {
	return len(d)
}

func (d Dinos) Less(i, j int) bool {
	if d[i].Speed() <= d[j].Speed() {
		return true
	}
	return false
}

func (d Dinos) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type Dino struct {
	Name         string
	StrideLength int
	LegLength    int
}

func (d Dino) Speed() float64 {
	return math.Log(float64(d.StrideLength)*float64(d.LegLength)) * 9.8
}

func readFile1(f *os.File) (Dinos, error) {
	var d Dinos
	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read csv file: %s", err)
	}
	for i, line := range data {
		if i != 0 {
			if line[2] == "Biped" {
				stride, err := strconv.Atoi(line[1])
				if err != nil {
					fmt.Errorf("unable to convert string to int: %s", err)
				}
				d = append(d, &Dino{Name: line[0], StrideLength: stride})
			}
		}
	}
	return d, nil
}

func readFile2(f *os.File, ds Dinos) (Dinos, error) {
	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read csv file: %s", err)
	}

	for _, line := range data {
		for _, d := range ds {
			if line[0] == d.Name {
				leg, err := strconv.Atoi(line[1])
				if err != nil {
					fmt.Errorf("unable to convert string to int: %s", err)
				}
				d.LegLength = leg
			}
		}
	}
	return ds, nil
}

func printDinos(d Dinos) {
	for _, dino := range d {
		fmt.Printf("%s\n", dino.Name)
	}
}
