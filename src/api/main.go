package main

import (
	"fmt"
	"github.com/GrimlockMaster/santanauts/src/api/mail"
	"github.com/GrimlockMaster/santanauts/src/api/santanaut"
	"math/rand"
	"strings"
)

func main() {
	entrants := loadEntrants()
	hat := loadHat(entrants)
	drawFromHat(entrants, hat)

	fmt.Printf("Result: %v", entrants)

	sendMails(entrants)
}

func loadEntrants() map[string]*santanaut.Santanaut {
	target := make(map[string]*santanaut.Santanaut)

	target["paula"] = santanaut.New("paula", "Paula", "paulaicsr@gmail.com", 2, []string{"seba", "juli"})
	target["facu"] = santanaut.New("facu", "Facu", "ocampo.facundo@gmail.com", 2, []string{"estela"})
	target["maxi"] = santanaut.New("maxi", "Maxi", "maxxts@gmail.com", 1, []string{"juli"})
	target["inti"] = santanaut.New("inti", "Inti", "nopuedeserenserio@gmail.com", 2, []string{})
	target["juli"] = santanaut.New("juli", "Juli", "jucacheiro@gmail.com", 2, []string{"maxi", "paula"})
	target["milton"] = santanaut.New("milton", "Milton", "milton.gabriel.castro@gmail.com", 1, []string{})
	target["rodri"] = santanaut.New("rodri", "Rodri", "artofrodrigovega@gmail.com", 1, []string{"magui"})
	target["seba"] = santanaut.New("seba", "Seba", "seba.koni@gmail.com", 3, []string{"paula"})
	target["mati"] = santanaut.New("mati", "Mati", "delenor_stormrage@hotmail.com", 1, []string{})
	target["magui"] = santanaut.New("magui", "Magui", "magagalidomenech94@gmail.com", 1, []string{"rodri"})
	target["estela"] = santanaut.New("estela", "Estela", "estelita.dietrich@gmail.com", 2, []string{"facu"})
	target["mili"] = santanaut.New("mili", "Mili", "nanablack19@gmail.com", 1, []string{})
	target["jorge"] = santanaut.New("jorge", "Jorge", "hjorgec@hotmail.com", 1, []string{})

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
			naut.Targets = append(naut.Targets, hat[index].Id)

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

func sendMails(entrants map[string]*santanaut.Santanaut) {
	subject := "Asignacion de Santanaut! (puede haber mas de un correo, esto es correcto)"
	mail.Init()

	for id, naut := range entrants {
		gifters := make([]string, 0)
		receivers := make([]string, 0)

		for _, nautGifter := range entrants {
			if santanaut.Contains(nautGifter.Targets, id) {
				gifters = append(gifters, nautGifter.Name)
				receivers = append(receivers, nautGifter.Email)
			}
		}

		r := mail.NewRequest(receivers, subject)
		var template string
		var params map[string]string

		if naut.Entries == 1 {
			template = "templates/templateSingle.html"
			params = map[string]string{"target": naut.Name, "sender": gifters[0]}
		} else {
			template = "templates/templateMultiple.html"
			params = map[string]string{"target": naut.Name, "senders": strings.Join(gifters, ", ")}
		}

		r.Send(template, params)
	}
}
