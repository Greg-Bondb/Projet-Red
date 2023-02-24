package main

import (
	"fmt"
	"os"
	"time"
)

//   Fin du jeu   //
func (C *character) end(fight int) {
	if fight == 3 && C.name == "Lucas" {
		slow("Vous avez vaincu le Roi Gobelin, mais la puissance de votre attaque a provoquée une déchirure de l'espace temps.\n")
		time.Sleep(3 * time.Second)
		slow("Vous vous envolez alors en regardant le monde se déchirer sous vos pas...\n")
		time.Sleep(3 * time.Second)
		slow("Heuresement vous arrivez à voler une part de pizza avant la fin et vous tappez votre meilleur croc dedans.\n")
	} else if fight == 3 {
		slow("Vous avez vaincu le Roi Gobelin, mais l'intensité de votre combat vous à fait vous évanouir sur place.\n")
		time.Sleep(3 * time.Second)
		slow("En fermant les yeux, vous appercevez les larbins du roi qui foncent vers vous.\n")
		time.Sleep(3 * time.Second)
		slow("Vous sentez votre fin approcher mais une mystérieuse personne encapuchonnée viens à votre secours...\n")
		time.Sleep(3 * time.Second)
		slow("Vous vous reveillez dans une foret au coté de cette personne.\n")
		time.Sleep(3 * time.Second)
		slow("Vous vous demandez pourquoi vous à t-elle sauvé.\n")
		time.Sleep(3 * time.Second)
		slow("Soudain, la personne retire sa capuche affirme que vous devez vous reveiller pour continuer de bosser sur Ytrack.\n")
	}
	if fight == 3 {
		time.Sleep(3 * time.Second)
		slow("La suite dans la partie 2...\n")
		slow("------------------------------------------------\n")
		time.Sleep(2 * time.Second)
		slow("Vous avez finit le jeu en étant lvl : ")
		fmt.Print(C.lvl, "\n")
		slow("Vous êtes mort : ")
		fmt.Print(mort, "\n")
		slow("------------------------------------------------\n")
		time.Sleep(2 * time.Second)
		slow("Merci d'avoir joué a notre jeu, nous espérons que ce petit RPG vous aura bien diverti.\n")
		slow("------------------------------------------------\n")
		time.Sleep(2 * time.Second)
		slow("Crédit :\n")
		time.Sleep(1 * time.Second)
		slow("---Balatre Grégory---\n")
		time.Sleep(1 * time.Second)
		slow("---Zelphati Nolan---\n")
		time.Sleep(3 * time.Second)
		p := Convert2Ascii(ScaleImage(Init()))
		fmt.Print(string(p))
		time.Sleep(6 * time.Second)
		os.Exit(0)
	}
}
