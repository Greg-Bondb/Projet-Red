package main

import (
	"fmt"
	"time"
)

//   Initialisation dans le jeu   //
func (C *character) charcreation() {
	a := 0
	fmt.Println("Bonjour, bienvenue dans l'Aincrad, pour lancer la simulation du RPG écrivez 'start' dans le terminal de commande.")
	time.Sleep(1 * time.Second)
	for {
		if a <= 3 {
			var interaction string
			fmt.Scanf("%s\n", &interaction)

			switch interaction {
			case "start":
				C.start()
				return
			}
			if a == 0 {
				fmt.Println("D'accord... Si vous changez d'avis, je suis là hein, de toute façons je peux pas partir.")
				time.Sleep(2 * time.Second)
				fmt.Println("Enfin bref, faites moi signe, pendant ce temps je vais regarder ce mur.")
				a++
			} else if a == 1 {
				fmt.Println("Dites juste 'start', c'est pas compliqué non ?.")
				a++
			} else if a == 2 {
				fmt.Println("Vous vous souvenez quand vous lanciez un jeu, il fallait écrire start, c'était génial hein ? Eh bien refaite le.")
				a++
			} else {
				fmt.Println("'start', 'start', 'start', 'start', 'start'.")
			}
		}
	}
}
