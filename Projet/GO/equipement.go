package main

import "fmt"

//   Ajouter un équipement   //
func (C *character) addstuff(equip string, nb int) {
	if nb == 1 && C.equip.head != equip {
		C.add(C.equip.head, 0)
		C.equip.head = equip
		C.remove(equip)
		fmt.Println("Vous vous équipez du", equip)
		C.verifstuff()
		return
	} else if nb == 2 && C.equip.body != equip {
		C.add(C.equip.body, 0)
		C.equip.body = equip
		C.remove(equip)
		fmt.Println("Vous vous équipez de la", equip)
		C.verifstuff()
		return
	} else if nb == 3 && C.equip.foot != equip {
		C.add(C.equip.foot, 0)
		C.equip.foot = equip
		C.remove(equip)
		fmt.Println("Vous vous équipez des", equip)
		C.verifstuff()
		return
	} else if nb == 4 && equip == "épée rouillée [5 damage]" && C.equip.hand != equip {
		C.add(C.equip.hand, 0)
		C.equip.hand = equip
		C.remove(equip)
		fmt.Println("Vous vous équipez de l'", equip)
		C.skill = append(C.skill, "épée rouillée")
		return
	} else if nb == 4 && equip == "épée empoisonnée [5 damage]" && C.equip.hand != equip {
		C.add(C.equip.hand, 0)
		C.equip.hand = equip
		C.remove(equip)
		fmt.Println("Vous vous équipez de l'", equip)
		C.skill = append(C.skill, "épée empoisonnée")
		return
	}
	println("Vous etes déja équipé de cet objet")
}

//   Ajouter des bonus en fonction du stuff équipé   //
func (C *character) verifstuff() {
	if C.class == "Humain" {
		C.maxlife = 100
	} else if C.class == "Elfe" {
		C.maxlife = 80
	} else {
		C.maxlife = 120
	}
	if C.equip.head == "chapeau de l'aventurier [15 pv]" {
		C.maxlife += 15
	}
	if C.equip.head == "casqu'ette [5 pv]" {
		C.maxlife += 5
	}
	if C.equip.body == "tunique de l'aventurier [25 pv]" {
		C.maxlife += 25
	}
	if C.equip.body == "armure en carton [10 pv]" {
		C.maxlife += 10
	}
	if C.equip.foot == "bottes de l'aventurier [15 pv]" {
		C.maxlife += 15
	}
	if C.equip.foot == "jambière en lin [5 pv]" {
		C.maxlife += 5
	}
	if C.equip.hand == "épée rouillée [3 damage]" {
		C.damage += 3
	}
	if C.equip.hand == "épée empoisonnée [5 damage]" {
		C.maxlife += 5
	}
}

//   Supprimer un équipement   //
func (C *character) removestuff(number int) {
	if number < len(C.inventory)-1 {
		C.inventory = append(C.inventory[:number], C.inventory[number+1:]...)
	} else if number == len(C.inventory)-1 {
		C.inventory = C.inventory[:number]
	}
}

//   Pour voir ses équipements   //
func (C *character) stuff() {
	fmt.Println("------------------------------------------------")
	fmt.Println("head stuff :", C.equip.head)
	fmt.Println("body stuff :", C.equip.body)
	fmt.Println("foot stuff :", C.equip.foot)
	fmt.Println("hand stuff :", C.equip.hand)
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
