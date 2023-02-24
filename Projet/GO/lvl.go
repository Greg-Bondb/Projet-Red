package main

import (
	"fmt"
	"time"
)

//   Lvl up le personnage   //
func (C *character) xp(nb int) {
	for nb > 0 {
		if nb >= C.needxp-C.xpac {
			nb -= (C.needxp - C.xpac)
			C.lvl += 1
			C.needxp += 250
			C.maxlife += 5
			C.manamax += 5
			C.initiative += 1
			C.damage += 1
			time.Sleep(1 * time.Second)
			fmt.Println("Vous venez de gagner un niveau, vos stats augmentent légèrement.")
			time.Sleep(1 * time.Second)
			fmt.Println("Vous êtes désormais lvl", C.lvl)
			C.xpac = 0
		} else {
			C.xpac += nb
			nb = 0
		}
	}
}
