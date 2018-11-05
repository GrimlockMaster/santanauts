package main

import (
	"fmt"
	"github.com/GrimlockMaster/santanauts/src/api/santanaut"
	"math/rand"
)

func main() {
	entrants := loadEntrants()
	hat := loadHat(entrants)
	drawFromHat(entrants, hat)

	fmt.Printf("Result: %v", entrants)
}

func loadEntrants() map[string]*santanaut.Santanaut {
	target := make(map[string]*santanaut.Santanaut)
	target["seba"] = santanaut.New("seba", "Sebastian", "seba.koni@gmail.com", 1, []string{"paula"})
	target["paula"] = santanaut.New("paula", "Paula", "paulaicsr@gmail.com", 1, []string{"seba"})
	target["juli"] = santanaut.New("juli", "Juli", "juli@gmail.com", 1, []string{"paula"})
	target["inti"] = santanaut.New("inti", "Inti", "inti@gmail.com", 1, []string{})

	return target
}

func loadHat(entrants map[string]*santanaut.Santanaut) []*santanaut.Santanaut {
	hat := make([]*santanaut.Santanaut, 0)

	for _, naut := range entrants {
		for i := 0; i < naut.Entries; i++ {
			hat = append(hat, naut)
		}
	}

	return hat
}

func drawFromHat(entrants map[string]*santanaut.Santanaut, hat []*santanaut.Santanaut) {
	fmt.Sprintf("Result: %v", entrants)
	for _, naut := range entrants {
		for i := 0; i < naut.Entries; i++ {
			index := findTarget(naut, hat)
			naut.Targets = append(naut.Targets, hat[index].Name)

			hat = append(hat[:index], hat[index+1:]...)
		}
	}
}

func findTarget(naut *santanaut.Santanaut, hat []*santanaut.Santanaut) int {
	found := false
	tries := 0

	for !found {
		try := rand.Intn(len(hat))

		if naut.IsValidTarget(*hat[try]) {
			return try
		}

		if tries > 100 {
			panic("too many tries!!!")
		}

		tries = tries + 1
	}

	return -1
}
