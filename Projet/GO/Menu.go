package main

import (
	"fmt"
	"os"
	"time"
)

//   Menu   //
func (C1 *character) menu() {
	fmt.Println("------------------------------------------------")
	fmt.Println("info. accéder à vos informations")
	fmt.Println("inv. accéder à votre inventaire")
	fmt.Println("stuff. voir les équipements équipés")
	fmt.Println("trade. parler au marchand")
	fmt.Println("forge. permet d'accéder à la forge")
	fmt.Println("fight. permet de s'excercer au combat")
	fmt.Println("bonus. affiche le nom des deux artistes cachés")
	fmt.Println("left. quitter le jeu")
	fmt.Println("------------------------------------------------")
	var interaction string
	fmt.Scanf("%s\n", &interaction)

	switch interaction {
	case "info":
		Clear()
		C1.displayinfo()
	case "inv":
		Clear()
		C1.accesinventory()
	case "stuff":
		Clear()
		C1.stuff()
	case "trade":
		Clear()
		C1.trader()
	case "forge":
		Clear()
		C1.forge()
	case "fight":
		Clear()
		C1.TrainingFight(M)
	case "bonus":
		Clear()
		C1.bonus()
	case "left":
		Clear()
		os.Exit(0)
	case "cheat":
		Clear()
		C1.cheat()
	default:
		fmt.Println("Je n'ai pas compris votre requête")
		time.Sleep(2 * time.Second)
	}
}
