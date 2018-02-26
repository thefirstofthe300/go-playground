package dino

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Dinos is a map used to make ingesting data easy.
type Dinos map[string]*Dino

// Speeds is an array used to easily sort the dinos by speed.
type Speeds []*Dino

// NewSpeeds creates a new DinoSpeeds array from the given Dinos map
func NewSpeeds(d Dinos) Speeds {
	var ds Speeds

	for _, v := range d {
		ds = append(ds, v)
	}
	return ds
}

func (d Speeds) Len() int {
	return len(d)
}

func (d Speeds) Less(i, j int) bool {
	if d[i].Speed() <= d[j].Speed() {
		return true
	}
	return false
}

func (d Speeds) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Dino is the data struct used to store metadata about a dino.
type Dino struct {
	Name         string
	StrideLength int
	LegLength    int
}

// Speed returns the speed of the dino. Used by the sort to determine the fastest
// dino.
func (d *Dino) Speed() float64 {
	return math.Log(float64(d.StrideLength)*float64(d.LegLength)) * 9.8
}

// ReadFile reads a file.
func ReadFile(f *os.File, d Dinos, wantBipeds bool) (Dinos, error) {
	log.Printf("reading file %s", f.Name())
	r := csv.NewReader(f)
	input, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	var isFile1 bool

	for i, record := range input {
		if i == 0 {
			if record[1] == "StrideLength" {
				isFile1 = true
			} else {
				isFile1 = false
			}
			continue
		}
		if isFile1 {
			if record[2] == "Biped" && wantBipeds {
				if val, ok := d[record[0]]; !ok {
					stridelen, err := strconv.Atoi(record[1])
					if err != nil {
						return nil, fmt.Errorf("unable to convert string to int: %s", err)
					}
					newDino := &Dino{Name: record[0], StrideLength: stridelen}
					d[record[0]] = newDino
					log.Printf("added StrideLength %d to Dino", stridelen)
				} else {
					stridelen, err := strconv.Atoi(record[1])
					if err != nil {
						return nil, fmt.Errorf("unable to convert string to int: %s", err)
					}
					val.StrideLength = stridelen
					log.Printf("added StrideLength %d to Dino", stridelen)
				}
			} else {
				if val, ok := d[record[0]]; !ok {
					stridelen, err := strconv.Atoi(record[1])
					if err != nil {
						return nil, fmt.Errorf("unable to convert string to int: %s", err)
					}
					newDino := &Dino{Name: record[0], StrideLength: stridelen}
					d[record[0]] = newDino
					log.Printf("added StrideLength %d to Dino", stridelen)
				} else {
					stridelen, err := strconv.Atoi(record[1])
					if err != nil {
						return nil, fmt.Errorf("unable to convert string to int: %s", err)
					}
					val.StrideLength = stridelen
					log.Printf("added StrideLength %d to Dino", stridelen)
				}
			}
		} else {
			if val, ok := d[record[0]]; ok {
				leglen, err := strconv.Atoi(record[1])
				if err != nil {
					return nil, fmt.Errorf("unable to convert string to int: %s", err)
				}
				val.LegLength = leglen
				log.Printf("added LegLength %d to Dino", leglen)
			}
		}
	}
	return d, nil
}

// PrintDinos prints the dinos that exist in the given Speeds array.
func PrintDinos(d Speeds) {
	for _, dino := range d {
		speed := dino.Speed()
		fmt.Printf("%s %v %v %v\n", dino.Name, dino.LegLength, dino.StrideLength, speed)
	}
}
