package main

import "fmt"

//   Bonus nom artistes   //
func (C *character) bonus() {
	fmt.Println("------------------------------------------------")
	fmt.Println("artiste de la partie 2 : Thanos le super-vilain de l'univers Marvel")
	fmt.Println("artiste de la partie 3 : Maximus Decimus Meridius dans Gladiator")
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez return pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "return":
		return
	default:
		fmt.Println("Je n'ai pas compris votre requête")
	}
}
