package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode/utf8"
)

type Produkt struct {
	name       string
	menge      int
	maßeinheit string
	kategorie  string
}

type Einkaufsliste map[string]Produkt

func ErstelleNeueEinkaufsListe() Einkaufsliste {
	var einkaufsListe Einkaufsliste
	einkaufsListe = make(map[string]Produkt)
	return einkaufsListe
}

func (e Einkaufsliste) FügeProduktHinzu(produkt Produkt) {
	alteWert, eintragSchonExistiert := e[produkt.name]
	if eintragSchonExistiert {
		produkt.menge += alteWert.menge
	}
	e[produkt.name] = produkt
}

func (e Einkaufsliste) PrintEinkaufsListe() {
	fmt.Println("GEMÜSE:\n")
	var längsteLängeName int
	var längsteLängeMenge int
	for _, produkt := range e {
		if längsteLängeName < utf8.RuneCountInString(produkt.name) {
			längsteLängeName = utf8.RuneCountInString(produkt.name)
		}
		if längsteLängeMenge < utf8.RuneCountInString(strconv.Itoa(produkt.menge)) {
			längsteLängeMenge = utf8.RuneCountInString(strconv.Itoa(produkt.menge))
		}
	}
	for _, produkt := range e {

		if e[produkt.name].kategorie == "Gemüse" {
			var längsteLängeNameAlsString string
			var längsteLängeMengeAlsString string
			for i := 0; i < längsteLängeName-utf8.RuneCountInString(produkt.name); i++ {
				längsteLängeNameAlsString += " "
			}
			for i := 0; i < längsteLängeMenge-utf8.RuneCountInString(strconv.Itoa(produkt.menge)); i++ {
				längsteLängeMengeAlsString += " "
			}
			fmt.Println(produkt.name, längsteLängeNameAlsString, produkt.menge, längsteLängeMengeAlsString, produkt.maßeinheit)
		}
	}
	fmt.Println("\nGEWÜRZE:\n")
	for _, produkt := range e {
		if e[produkt.name].kategorie == "Gewürze" {
			var längsteLängeNameAlsString string
			var längsteLängeMengeAlsString string
			for i := 0; i < längsteLängeName-utf8.RuneCountInString(produkt.name); i++ {
				längsteLängeNameAlsString += " "
			}
			for i := 0; i < längsteLängeMenge-utf8.RuneCountInString(strconv.Itoa(produkt.menge)); i++ {
				längsteLängeMengeAlsString += " "
			}
			fmt.Println(produkt.name, längsteLängeNameAlsString, produkt.menge, längsteLängeMengeAlsString, produkt.maßeinheit)
		}
	}
	fmt.Println("\nSONNSTIGES:\n")
	for _, produkt := range e {
		if e[produkt.name].kategorie == "Sonnstiges" {
			var längsteLängeNameAlsString string
			var längsteLängeMengeAlsString string
			for i := 0; i < längsteLängeName-utf8.RuneCountInString(produkt.name); i++ {
				längsteLängeNameAlsString += " "
			}
			for i := 0; i < längsteLängeMenge-utf8.RuneCountInString(strconv.Itoa(produkt.menge)); i++ {
				längsteLängeMengeAlsString += " "
			}
			fmt.Println(produkt.name, längsteLängeNameAlsString, produkt.menge, längsteLängeMengeAlsString, produkt.maßeinheit)
		}
	}
}

func (e Einkaufsliste) ErstellePlov() {
	e.FügeProduktHinzu(Produkt{"Möhren", 400, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Zwiebeln", 400, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Knoblauchzehen", 4, "Ziehen", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Lorbeerblätter", 2, "Stück", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Gemüsebrühe Pulver", 5, "Esslöffel", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Paprika Edelsüß", 1, "", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Paprika Rosenscharf", 1, "", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Kumin", 1, "", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Bratöl", 75, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Langkornreis", 250, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Rosinen", 50, "g", "Sonnstiges"})
}

func (e Einkaufsliste) ErstelleGelbeLinsenSuppe() {
	e.FügeProduktHinzu(Produkt{"Süßkartoffeln", 580, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Frühlingszwiebeln", 190, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Zwiebeln", 225, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Knoblauchzehen", 2, "Ziehen", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Currypulver", 1, "", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Gemüsebrühe Pulver", 1, "Esslöffel", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Petersilie", 3, "Esslöffel", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Gelbe Linsen", 250, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Tomatenmark", 60, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Kokosmilch", 400, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Öl", 24, "g", "Sonnstiges"})
}

func (e Einkaufsliste) ErstelleNudelSalatMitPesto() {
	e.FügeProduktHinzu(Produkt{"Cherrytomaten", 250, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Rucola", 125, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Kräutermischung", 50, "g", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Nudeln", 500, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Pesto", 190, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Tomaten in Öl", 180, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Öl", 24, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Cashewkerne", 50, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Knoblauchzehen", 1, "Ziehen", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Pesto", 190, "g", "Sonnstiges"})
}

func (e Einkaufsliste) ErstelleBologneseMitVeganemHack() {

	e.FügeProduktHinzu(Produkt{"Knoblauchzehen", 3, "Ziehen", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Zwiebeln", 350, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Möhren", 250, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Sellerie", 250, "g", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Champignons", 250, "", "Gemüse"})
	e.FügeProduktHinzu(Produkt{"Gemüsebrühe Pulver", 30, "g", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Lorbeerblätter", 2, "Stück", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Italienische Kräutermischung", 15, "g", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Sojasauce", 18, "g", "Gewürze"})
	e.FügeProduktHinzu(Produkt{"Hackfleisch", 360, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Dosen Tomaten", 3, "Stück", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Öl", 72, "g", "Sonnstiges"})
	e.FügeProduktHinzu(Produkt{"Spaghetti Nudeln", 500, "g", "Sonnstiges"})
}

func (e Einkaufsliste) PrintSalatOderHauptgerichtAuswahl() {
	fmt.Println("\nWähle die Kategorie aus:\n")
	fmt.Println("1. Salat")
	fmt.Println("2. Hauptgerichte")
	for {
		var imput byte
		fmt.Scanf("%d", &imput)
		if imput == 1 {
			e.PrintSalatAuswahl()
			break
		} else if imput == 2 {
			e.PrintHauptgerichtAuswahl()
			break
		}
	}
}

func (e Einkaufsliste) PrintSalatAuswahl() {
	fmt.Println("\nWähle Salat aus:\n")
	salatListe := []string{"Nudel Salat Mit Pesto"}
	sort.Strings(salatListe)
	fmt.Println("0. Zurück\n")
	for i := 0; i < len(salatListe); i++ {
		fmt.Printf("%d. %s\n", i+1, salatListe[i])
	}
	fmt.Printf("\n%d. Erstelle neue Einkaufsliste\n", len(salatListe)+1)
	for {
		var imput byte
		var salat string
		fmt.Scanf("%d", &imput)
		if imput == 0 {
			e.PrintSalatOderHauptgerichtAuswahl()
		} else if imput > byte(len(salatListe)+1) {
			e.PrintSalatAuswahl()
		} else if imput == byte(len(salatListe)+1) {
			e.PrintEinkaufsListe()
		} else {
			salat = salatListe[imput-1]
			switch {
			case salat == "Nudel Salat Mit Pesto":
				e.ErstelleNudelSalatMitPesto()
			}
			fmt.Printf("\n%s wurde hinzugefügt\n", salat)
			e.PrintSalatAuswahl()
		}

	}
}

func (e Einkaufsliste) PrintHauptgerichtAuswahl() {
	fmt.Println("\nWähle Hauptgericht aus:\n")
	hauptgerichtListe := []string{"Bolognese mit Veganem Hack", "Gelbe Linsen Suppe", "Plov"}
	sort.Strings(hauptgerichtListe)
	fmt.Println("0. Zurück\n")
	for i := 0; i < len(hauptgerichtListe); i++ {
		fmt.Printf("%d. %s\n", i+1, hauptgerichtListe[i])
	}
	fmt.Printf("\n%d. Erstelle neue Einkaufsliste\n", len(hauptgerichtListe)+1)
	for {
		var imput byte
		var hauptgericht string
		fmt.Scanf("%d", &imput)
		if imput == 0 {
			e.PrintSalatOderHauptgerichtAuswahl()
		} else if imput > byte(len(hauptgerichtListe)+1) {
			e.PrintHauptgerichtAuswahl()
		} else if imput == byte(len(hauptgerichtListe)+1) {
			e.PrintEinkaufsListe()
		} else {
			hauptgericht = hauptgerichtListe[imput-1]
			switch {
			case hauptgericht == "Bolognese mit Veganem Hack":
				e.ErstelleBologneseMitVeganemHack()
			case hauptgericht == "Gelbe Linsen Suppe":
				e.ErstelleGelbeLinsenSuppe()
			case hauptgericht == "Plov":
				e.ErstellePlov()
			}
			(fmt.Printf("\n%s wurde hinzugefügt\n", hauptgericht))
			e.PrintHauptgerichtAuswahl()
		}

	}
}

func main() {
	fmt.Println("Die Websites, von denen die Rezepte stammen:")
	fmt.Println("https://veggie-einhorn.de/")
	einkaufsListe := ErstelleNeueEinkaufsListe()
	einkaufsListe.PrintSalatOderHauptgerichtAuswahl()
}
