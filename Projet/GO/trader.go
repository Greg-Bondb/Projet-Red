package main

import (
	"fmt"
	"time"
)

//   Pour marchander   //
func (C *character) trader() {
	fmt.Println("Marchand itinérant : Salut, tu viens pour marchander ? Dit moi si tu souhaite acheter ou vendre un objet.")
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous possedez", C.money, "pièces d'or")
	fmt.Println("------------------------------------------------")
	fmt.Println("buy. permet d'acheter un objet.")
	fmt.Println("sell. permet de vendre un objet.")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "buy":
		C.buy()
	case "sell":
		C.sell()
	case "return":
		fmt.Println("Vous revenez au menu.")
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Pour acheter des objets   //
func (C *character) buy() {
	fmt.Println("Marchand itinérant : jette un coup d'oeil à ma boutique, j'ai plein de choses à vendre.")
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous possedez", C.money, "pièces d'or")
	fmt.Println("------------------------------------------------")
	fmt.Println("1) potion de santé: Une potion qui rend 50 pv [3 po]")
	fmt.Println("2) potion de mana: Une potion qui rend 25 points de mana [3 po]")
	fmt.Println("3) fiole de poison: Une fiole de poison qui inflige 10 pv pendant 3s [6 po]")
	fmt.Println("4) livre de sort (boule de feu): un sort qui permet de projeter une boule de feu [25 po]")
	fmt.Println("5) Fourrure de Loup: une fourrure banale de gentil chien-chien[4 po]")
	fmt.Println("6) Peau de Troll: une belle peau fraichement dépecée d'un troll [7 po]")
	fmt.Println("7) Cuir de Sanglier: un cuir plutot résistant [3 po]")
	fmt.Println("8) Plume de Corbeau: une couleur magnifique pour une simple plume noire, non ? [2 po]")
	fmt.Println("9) augmentation d'inventaire: ajoute 10 places à votre inventaire [30 po]")
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "0":
		return
	case "1":
		C.add("potion de santé", 3)
		popogratuite++
	case "2":
		C.add("potion de mana", 3)
	case "3":
		C.add("fiole de poison", 6)
	case "4":
		for i := 0; i < len(C.inventory); i++ {
			if C.inventory[i] == "livre de sort (boule de feu)" {
				println("Vous avez déjà ce sort en votre possession.")
				return
			}
		}
		for i := 0; i < len(C.skill); i++ {
			if C.skill[i] == "boule de feu" {
				println("Vous avez déjà appris ce sort.")
				return
			}
		}
		C.add("livre de sort (boule de feu)", 25)
	case "5":
		C.add("fourrure de loup", 10)
	case "6":
		C.add("peau de troll", 7)
	case "7":
		C.add("cuir de sanglier", 3)
	case "8":
		C.add("plume de corbeau", 2)
	case "9":
		C.upgradeinventoryslot(30)
	case "return":
		fmt.Println("Vous revenez au menu.")
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
		return
	}
	fmt.Println("Merci pour ton achat, revient vite.")
}

//   Pour vendre des objets  //
func (C *character) sell() {
	if len(C.inventory) == 0 {
		fmt.Println("votre inventaire est vide")
		return
	}
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	fmt.Println("Vous possedez", C.money, "pièces d'or")
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i], "[2 po]")
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case 0:
		fmt.Println("Je n'ai pas compris votre requête")
		return
	case k:
		if k < 0 || k > len(C.inventory) {
			fmt.Println("Revient vite me voir.")
		} else {
			C.remove(C.inventory[k-1])
			fmt.Println("Vous recuperez 2 pièces d'or.")
			fmt.Println("J'aime faire des affaires avec toi, revient me voir à l'occasion.")
			C.money += 2
			time.Sleep(2 * time.Second)
		}
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}
