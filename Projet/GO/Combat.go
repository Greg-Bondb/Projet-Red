package main

import (
	"fmt"
	"time"
)

//   Gère les monstres, le nombre de tours, la fin des combats et la fin du jeu   //
func (C *character) TrainingFight(M *Monstre) {
	var fight int
	fmt.Println("------------------------------------------------")
	fmt.Println("1) affronter un gobelin d'entrainement.")
	fmt.Println("2) affronter un gobelin supérieur.")
	fmt.Println("3) affronter Oryx le roi des gobelins")
	fmt.Println("------------------------------------------------")
	fmt.Println("utilisez '0' pour revenir en arrière")
	fmt.Println("------------------------------------------------")
	var k int
	fmt.Scanf("%d\n", &k)
	switch k {
	case 0:
		return
	case 1:
		fight = 1
		M1.name = "Gobelin d'entrainement"
		M1.pvmax = 40
		M1.pv = 40
		M1.damage = 5
		M1.initiative = 1
	case 2:
		fight = 2
		M1.name = "Gobelin supérieur"
		M1.pvmax = 80
		M1.pv = 80
		M1.damage = 10
		M1.initiative = 3
	case 3:
		fight = 3
		M1.name = "Oryx Roi des Gobelins"
		M1.pvmax = 100
		M1.pv = 100
		M1.damage = 15
		M1.initiative = 5
	case k:
		if k < 0 || k > 3 {
			fmt.Println("Je n'ai pas compris votre requête")
			return
		}
	default:
		fmt.Println("Je n'ai pas compris votre requête")
		return
	}
	if fight <= 2 {
		fmt.Println("Vous entrez dans un donjon, un", M1.name, "vous attaque.")
	} else {
		fmt.Println("Vous entrez dans un chateau, le", M1.name, "vous fait face.")
		time.Sleep(1 * time.Second)
		fmt.Println("Vous vous sentez intimidé par sa présence")
		time.Sleep(1 * time.Second)
		fmt.Println("Vos stats baissent légèrement")
		C.maxlife -= 10
		C.manamax -= 10
	}
	for {
		if C1.life > 0 && M1.pv > 0 {
			fmt.Println("------------------------------------------------")
			fmt.Println("Vous êtes au tour n°", Tour)
			if !C.charTurn() {
			} else {
				break
			}
		} else {
			if M1.pv == 0 {
				fmt.Println("Vous avez vaincu le", M1.name)
				time.Sleep(2 * time.Second)
				if fight == 1 {
					fmt.Println("Vous récuperez 5 pièces d'or.")
					C.money += 5
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 1000 points d'expérience")
					C.xp(1000)
				} else if fight == 2 {
					fmt.Println("Vous récuperez 10 pièces d'or.")
					C.money += 10
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 1250 points d'expérience")
					C.xp(1250)
				} else if fight == 3 {
					fmt.Println("Vous récuperez 20 pièces d'or.")
					C.money += 20
					time.Sleep(2 * time.Second)
					fmt.Println("Vous gagnez 2000 points d'expérience")
					C.xp(2000)
				}
				time.Sleep(2 * time.Second)
				fmt.Println("xp nécessaire pour lvl up :", C.needxp-C.xpac)
				fmt.Println("------------------------------------------------")
				if Combat == 0 && M1.name == "Gobelin d'entrainement" {
					time.Sleep(2 * time.Second)
					fmt.Println("Le Gobelin d'entrainement laisse tomber derrière lui le livre de sort : pichenette")
					time.Sleep(2 * time.Second)
					C.add("livre de sort (pichenette)", 0)
					Combat = 1
				} else if Combat == 1 && M1.name == "Gobelin supérieur" {
					time.Sleep(2 * time.Second)
					fmt.Println("Le Gobelin supérieur laisse tomber derrière lui le livre de sort : tatane celeste")
					time.Sleep(2 * time.Second)
					C.add("livre de sort (tatane celeste)", 0)
					Combat = 2
				}
				time.Sleep(2 * time.Second)
				if fight < 2 {
					fmt.Println("Vous vous dirigez victorieusement vers la sortie mais votre soif de combat n'a pas été rassasiée.")
					time.Sleep(2 * time.Second)
				}
				C.end(fight)
				Tour = 1
				break
			}
		}
		if C1.life == 0 {
			fmt.Println("------------------------------------------------")
			time.Sleep(2 * time.Second)
			fmt.Println("Le ", M1.name, "vous a mis une raclée.")
			time.Sleep(2 * time.Second)
			C.dead()
			time.Sleep(2 * time.Second)
			fmt.Println("Vous reprenez connaisance en dehors du donjon.")
			time.Sleep(2 * time.Second)
			Tour = 1
			if M1.name == "Gobelin supérieur" {
				C.maxlife += 10
				C.manamax += 10
			}
			break
		}
	}
}

//   Choisir ce que l'on veut faire dans un combat   //
func (C *character) charTurn() bool {
	fmt.Println("------------------------------------------------")
	time.Sleep(1 * time.Second)
	fmt.Println("1) Menu Attack")
	fmt.Println("2) acceder à l'inventaire")
	fmt.Println("3) prendre la fuite ")
	fmt.Println("------------------------------------------------")
	var interaction int
	fmt.Scanf("%d\n", &interaction)

	switch interaction {
	case 1:
		C.MenuAttack()
		return false
	case 2:
		C.use2()
		return false
	case 3:
		Tour = 1
		fmt.Println("Vous prenez la fuite tel un misérable gueux apeuré.")
		time.Sleep(2 * time.Second)
		return true
	}
	fmt.Println("Je n'ai pas compris votre requête")
	return false
}

//   Choisir une attaque   //
func (C *character) MenuAttack() {
	a := 0
	nb := 1
	for a == 0 {
		nb = 1
		fmt.Println("------------------------------------------------")
		fmt.Println("Quelle attaque allez vous lancer ?")
		fmt.Println("Votre mana actuel est de :", "[", C.mana, "]")
		fmt.Println("------------------------------------------------")
		for i := 0; i < len(C.skill); i++ {
			fmt.Print(nb, ")", C.skill[i])
			if C.skill[i] == "coup de poing" {
				fmt.Println(" ", C.damage+8, "damage [5 mana]")
			} else if C.skill[i] == "épée rouillée" {
				fmt.Println(" ", C.damage+10, "damage [10 mana]")
			} else if C.skill[i] == "épée empoisonnée" {
				fmt.Println(" ", C.damage+15, "damage [10 mana]")
			} else if C.skill[i] == "boule de feu" {
				fmt.Println(" ", C.damage+18, "damage [15 mana]")
			} else if C.skill[i] == "pichenette" {
				fmt.Println(" ", C.damage+5, "damage [0 mana]")
			} else if C.skill[i] == "tatane celeste" {
				fmt.Println(" ", C.damage+15, "damage [12 mana]")
			} else if C.skill[i] == "La punition divine" {
				fmt.Println(" 15000 damage [10 mana]")
			}
			nb++
		}
		fmt.Println("------------------------------------------------")
		fmt.Println("écrivez le numéro qui correspond à l'attaque que vous voulez utiliser.")
		fmt.Println("utilisez '0' pour revenir en arrière")
		fmt.Println("------------------------------------------------")
		var k int
		fmt.Scanf("%d\n", &k)

		switch k {
		case 0:
			return
		case k:
			if k > len(C.skill) || k < 1 {
				a = 0
			} else {
				C.verifsort(C.skill[k-1])
				a++
			}
		}
	}
}

//   Utiliser une attaque et vérifier le mana nécéssaire pour la lancer   //
func (C *character) verifsort(rang string) {
	mana := 0
	deg := 0
	if rang == "coup de poing" {
		mana = 5
		deg = C.damage + 8
	} else if rang == "épée rouillée" {
		mana = 5
		deg = C.damage + 10
	} else if rang == "épée empoisonnée" {
		mana = 5
		deg = C.damage + 5
	} else if rang == "boule de feu" {
		mana = 15
		deg = C.damage + 18
	} else if rang == "pichenette" {
		mana = 0
		deg = C.damage + 5
	} else if rang == "tatane celeste" {
		mana = 10
		deg = C.damage + 15
	} else if rang == "La punition divine" {
		mana = 10
		deg = 15000
	}
	if C.mana < mana {
		fmt.Println("Vous n'avez pas asser de mana.")
		time.Sleep(2 * time.Second)
		C.MenuAttack()
		return
	} else {
		if !C.ynitiative() {
			fmt.Println("Vous avez moins d'initiative que votre adversaire, vous attaquez en dernier.")
			fmt.Println("------------------------------------------------")
			M.gobelinPattern(C)
			fmt.Println("------------------------------------------------")
		} else {
			fmt.Println("Vous avez plus d'initiative que votre adversaire, vous attaquez en premier.")
			fmt.Println("------------------------------------------------")
		}
		C.mana -= mana
		time.Sleep(1 * time.Second)
		fmt.Println(C1.name, "inflige à", M1.name, deg, "pts de dégat", "votre Mana restant est de", C.mana, "/", C.manamax)
		time.Sleep(1 * time.Second)
		if M1.pv-deg < 0 {
			M1.pv = 0
		} else {
			M1.pv -= deg
			if rang == "épée empoisonnée" {
				seconds := 2
				fmt.Println("Vous avez empoisonnée le", M1.name)
				for s := seconds; s >= 1; s-- {
					if M1.pv-5 < 0 {
						M1.pv = 0
						time.Sleep(2 * time.Second)
						fmt.Println("-5 pv")
						fmt.Println("life :", M1.pv)
						break
					}
					time.Sleep(2 * time.Second)
					fmt.Println("-5 pv")
					M1.pv -= 5
					fmt.Println("life :", M1.pv)
				}
			}
		}
		fmt.Println("il reste", M1.pv, "/", M1.pvmax, "pts de vie à", M1.name)
		time.Sleep(1 * time.Second)
		if C.ynitiative() {
			fmt.Println("------------------------------------------------")
			M.gobelinPattern(C)
		}
		return
	}
}

//   Permet d'attaquer le joueur   //
func (M *Monstre) gobelinPattern(C *character) {
	if M1.pv <= 0 {
		return
	}
	if Tour%3 == 0 {
		if C.life-(M1.damage*2) <= 0 {
			C.life = 0
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, (M1.damage * 2), "pts de dégat")
			time.Sleep(1 * time.Second)
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
			time.Sleep(1 * time.Second)
			return
		} else {
			C.life -= (M1.damage * 2)
			fmt.Println(M1.name, "inflige un coup critique à", C1.name, (M1.damage * 2), "pts de dégat")
			time.Sleep(1 * time.Second)
			fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
			time.Sleep(1 * time.Second)
			Tour++
			return
		}
	}
	if C.life-M1.damage <= 0 {
		C.life = 0
		fmt.Println(M1.name, "inflige à", C1.name, M1.damage, "pts de dégat")
		time.Sleep(1 * time.Second)
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		time.Sleep(1 * time.Second)
		return
	} else {
		C.life -= M1.damage
		fmt.Println(M1.name, "inflige à", C1.name, M1.damage, "pts de dégat")
		time.Sleep(1 * time.Second)
		fmt.Println("il vous reste", C1.life, "/", C1.maxlife, "pts de vie")
		time.Sleep(1 * time.Second)
		Tour++
	}
}

//   Utiliser un objet en combat   //
func (C *character) use2() {
	nb := 1
	var k int
	a := 0
	for a == 0 {
		nb = 1
		fmt.Println("------------------------------------------------")
		for i := 0; i < len(C.inventory); i++ {
			if C.inventory[i] == "potion de santé" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "potion de mana" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "fiole de poison" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "livre de sort (boule de feu)" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			} else if C.inventory[i] == "livre de sort (pichenette)" {
				fmt.Println(nb, ")", C.inventory[i])
				nb++
			}
		}
		fmt.Println("------------------------------------------------")
		fmt.Println("écrivez le numéro qui correspond à l'objet que vous voulez utiliser")
		fmt.Println("utilisez '0' pour revenir en arrière")
		fmt.Println("------------------------------------------------")
		fmt.Scanf("%d\n", &k)
		switch k {
		case 0:
			return
		case k:
			if k > len(C.inventory) || k < 1 {
				a = 0
			} else if C.verf2(k - 1) {
				fmt.Println("------------------------------------------------")
				a++
				M.gobelinPattern(C)
			}
		default:
			fmt.Println("Je n'ai pas compris votre requête")
		}
	}
}

//   Pour savoir qui attaque en premier   //
func (C *character) ynitiative() bool {
	if C.initiative >= M1.initiative {
		return true
	} else {
		return false
	}
}

//   Trouver un objet en fonction de son rang dans l'inventaire puis l'utiliser dans un combat (éviter d'utiliser un objet inutile en combat)   //
func (C *character) verf2(rang int) bool {
	if C.inventory[rang] == "potion de santé" && C.life != C.maxlife {
		C.takeapot()
		return true
	} else if C.inventory[rang] == "potion de mana" && C.mana != C.manamax {
		C.PotiondeMana()
		return true
	} else if C.inventory[rang] == "fiole de poison" {
		C.poison()
		C.removestuff(rang)
		return true
	} else if C.inventory[rang] == "livre de sort (boule de feu)" {
		C.spellBook("boule de feu", rang)
		return true
	} else if C.inventory[rang] == "livre de sort (pichenette)" {
		C.spellBook("pichenette", rang)
		return true
	} else {
		fmt.Println("vous etes déjà full")
		return false
	}
}
