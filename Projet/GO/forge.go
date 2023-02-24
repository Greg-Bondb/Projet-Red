package main

import (
	"fmt"
	"time"
)

//   Forger un équipement   //
func (C *character) forge() {
	if Combat == 2 && a == 0 {
		fmt.Println("forgeron : comme tu est un bon client, je te confie cette épée que j'ai trouvé ce matin.")
		time.Sleep(2 * time.Second)
		C.add("épée rouillée [3 damage]", 0)
		time.Sleep(2 * time.Second)
		fmt.Println("N'hésite pas à revenir me voir")
		time.Sleep(2 * time.Second)
		a++
		return
	} else {
		fmt.Println("Forgeron : bienvenu dans ma forge, je peux vous craft ce que vous voulez contre 5 pièces d'or, enfin que si vous m'apportez les mattériaux.")
		fmt.Println("------------------------------------------------")
		fmt.Println("Vous possedez", C.money, "pièces d'or")
		fmt.Println("------------------------------------------------")
		fmt.Println("1) chapeau de l'aventurier [15 pv] (1 plume de corbeau, 1 cuir de sanglier)")
		fmt.Println("2) tunique de l’aventurier [25 pv] (2 fourrure de loup, 1 peau de troll)")
		fmt.Println("3) bottes de l’aventurier [15 pv] (1 fourrure de loup, 1 cuir de sanglier)")
		fmt.Println("4) épée empoisonnée [5 damage] (1 épee rouillée, 1 fiole de poison)")
		fmt.Println("utilisez 'return' pour revenir en arrière")
		fmt.Println("------------------------------------------------")
		var interaction string
		fmt.Scanf("%s\n", &interaction)

		switch interaction {
		case "1":
			C.craft("plume de corbeau", "cuir de sanglier", "chapeau de l'aventurier [15 pv]")
		case "2":
			C.craft("fourrure de loup", "peau de troll", "tunique de l'aventurier [25 pv]")
		case "3":
			C.craft("fourrure de loup", "cuir de sanglier", "bottes de l'aventurier [15 pv]")
		case "4":
			C.craft("épée rouillée", "fiole de poison", "épée empoisonnée [5 damage]")
		case "return":
			fmt.Println("Vous revenez au menu.")
			return
		default:
			fmt.Println("Je n'ai pas compris votre requête")
		}
	}
}

//   Vérifier les objets avant de forger l'équipement   //
func (C *character) craft(q string, r string, p string) {
	verf1 := 0
	verf2 := 0
	for i := 0; i < len(C.inventory); i++ {
		if C.inventory[i] == q {
			verf1++
		}
		if C.inventory[i] == r {
			verf2++
		}
	}
	if p == "tunique de l'aventurier [25 pv]" && verf1 > 1 && verf2 > 0 {
		C.remove(q)
		C.remove(q)
		C.remove(r)
		C.add(p, 5)
		return
	} else if p != "tunique de l'aventurier [25 pv]" && verf1 > 0 && verf2 > 0 {
		C.remove(q)
		C.remove(r)
		C.add(p, 5)
		return
	}
	fmt.Println("vous n'avez pas les matériaux")
}

