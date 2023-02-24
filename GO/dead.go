package main

import "time"

//   Mort du perso   //
func (C *character) dead() {
	if C.life <= 0 {
		println("VOUS ETES MORT")
		time.Sleep(2 * time.Second)
		println("Vous revenez à la vie avec 50% de vos PV")
		C.life = C.maxlife / 2
		C.mana = C.manamax / 2
		mort += 1
	}
}
