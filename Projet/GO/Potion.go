package main

import (
	"fmt"
	"time"
)

//   Utiliser potion santé  //
func (C *character) takeapot() {
	if C.life == C.maxlife {
		fmt.Println("vous etes déjà full pv")
		return
	}
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "potion de santé" {
			break
		}
	}
	if C.life >= C.maxlife-50 {
		fmt.Println("Vous utilisez une potion de santé, vous regagnez 50 pv")
		C.life = C.maxlife
		C.removestuff(i)
	} else {
		fmt.Println("Vous utilisez une potion de santé, vous regagnez 50 pv")
		C.life += 50
		C.removestuff(i)
	}
}

//   Utiliser potion mana   //
func (C *character) PotiondeMana() {
	if C.mana == C.manamax {
		fmt.Println("vous etes déjà full mana")
		return
	}
	i := 0
	for i = 0; i < len(C.inventory); i++ {
		if C.inventory[i] == "potion de mana" {
			break
		}
	}
	if C.mana >= C.manamax-25 {
		fmt.Println("Vous utilisez une potion de mana, vous regagnez 25 pts de mana")
		C.mana = C.manamax
		C.removestuff(i)
	} else {
		fmt.Println("Vous utilisez une potion de mana, vous regagnez 25 pts de mana")
		C.mana += 25
		C.removestuff(i)
	}
}

//   Poison sur perso  //
func (C *character) poison() {
	seconds := 3
	fmt.Println("Vous buvez une fiole de poison, vous obtenez des effets indésirables pendant 3s")
	for s := seconds; s >= 1; s-- {
		if C.life-10 < 0 {
			C.life = 0
			time.Sleep(2 * time.Second)
			fmt.Println("-10 pv")
			fmt.Println("life :", C.life)
			print("\n")
			time.Sleep(1 * time.Second)
			C.dead()
			time.Sleep(2 * time.Second)
			print("\n")
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("-10 pv")
		C.life -= 10
		fmt.Println("life :", C.life)
	}
}
