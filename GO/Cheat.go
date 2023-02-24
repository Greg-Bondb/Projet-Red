package main

import (
	"fmt"
	"time"
)

//   ???   //
func (C *character) cheat() {
	fmt.Println("Dans un élan de générosité, les Dieux entendent votre prière et vous octroie la puissance de Le Lucas.")
	time.Sleep(2 * time.Second)
	fmt.Println("Vous sentez vos forces ce décupler.")
	time.Sleep(2 * time.Second)
	fmt.Println("Le Dieu Lucas prend possession de votre corps.")
	time.Sleep(2 * time.Second)
	C.name = "Lucas"
	C.class = "God"
	C.lvl = 118
	C.maxlife = 100000
	C.life = C.maxlife / 2
	C.manamax = 1000
	C.mana = 1000
	C.initiative = 100
	C.skill[0] = "La punition divine"
}
