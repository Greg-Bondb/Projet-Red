package main

import "fmt"

//   Pour supprimer un objet   //
func (C *character) delete() {
	nb := 1
	var k int
	fmt.Println("------------------------------------------------")
	for i := 0; i < len(C.inventory); i++ {
		fmt.Println(nb, ")", C.inventory[i])
		nb++
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("écrivez le numérot qui correspond à l'objet que vous voulez supprimer")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	fmt.Scanf("%d\n", &k)
	switch k {
	case 0:
		return
	case k:
		C.remove(C.inventory[k-1])
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}

//   Ajouter un objet   //
func (C *character) add(object string, price int) {
	if popogratuite == 0 && object == "potion de santé" {
		price = 0
		println("En tant que bon vendeur, je vous offre votre premiere potion de santé")
		popogratuite++
	}
	if object == "livre de sort (pichenette)" {
		println("Le", object, "a été ajouté à votre inventaire.")
		C.inventory = append(C.inventory, object)
		return
	}
	if object == "livre de sort (tatane celeste)" {
		println("Le", object, "a été ajouté à votre inventaire.")
		C.inventory = append(C.inventory, object)
		return
	}
	if object == "empty" {
		return
	}
	if len(C.inventory) >= C.slotinv {
		fmt.Println("votre inventaire est plein")
		return
	}
	if C.money > price {
		C.money -= price
		C.inventory = append(C.inventory, object)
		fmt.Println(object, "a été ajouté à votre inventaire")
	} else {
		fmt.Println("Vous n'avez pas assez d'argent")
	}
}

//   Supprimer un objet   //
func (C *character) remove(s string) {
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == s {
			break
		}
	}
	if i < len(C.inventory)-1 {
		fmt.Println("Vous enlevez", C.inventory[i], "de votre inventaire.")
		C.inventory = append(C.inventory[:i], C.inventory[i+1:]...)
	} else if i == len(C.inventory)-1 {
		fmt.Println("Vous enlevez", C.inventory[i], "de votre inventaire.")
		C.inventory = C.inventory[:i]
	}
}
