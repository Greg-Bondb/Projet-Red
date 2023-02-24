package main

import "fmt"

//   Info personnage   //
func (C *character) displayinfo() {
	tab := ""
	fmt.Println("------------------------------------------------")
	fmt.Println("name :", C.name)
	fmt.Println("class :", C.class)
	fmt.Println("lvl :", C.lvl)
	fmt.Println("lvl up :", C.xpac, "/", C.needxp, "point d'xp nécessaire")
	fmt.Println("maxlife :", C.maxlife)
	fmt.Println("life :", C.life)
	fmt.Println("maxmana :", C.manamax)
	fmt.Println("mana :", C.mana)
	for i := 0; i < len(C.skill); i++ {
		if i > 0 {
			tab += ", "
		}
		tab += "["
		tab += C.skill[i]
		tab += "]"
	}
	fmt.Println("skill :", tab)
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez 'return' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		fmt.Println("Vous revenez au menu.")
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}
