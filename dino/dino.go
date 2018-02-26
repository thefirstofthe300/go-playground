package main

import (
	"github.com/thefirstofthe300/go-playground/dino/dino"
	"log"
	"os"
	"sort"
)

func main() {
	dinos := make(dino.Dinos)
	f1, err := os.Open("file1.csv")
	if err != nil {
		log.Fatalf("Unable to open file1.csv: %s", err)
	}
	f2, err := os.Open("file2.csv")
	if err != nil {
		log.Fatalf("Unable to open file2.csv: %s", err)
	}
	dinos, err = dino.ReadFile(f1, dinos, true)
	if err != nil {
		log.Fatalf("Unable to load data from file1: %s", err)
	}
	dinos, err = dino.ReadFile(f2, dinos, true)
	if err != nil {
		log.Fatalf("Unable to load data from file2: %s", err)
	}
	ds := dino.NewSpeeds(dinos)
	sort.Sort(sort.Reverse(ds))
	dino.PrintDinos(ds)
}
