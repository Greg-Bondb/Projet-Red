package main

import (
	"fmt"
	"time"
)

//   Création du personnage : nom   //
func (C *character) start() {
	fmt.Println("Veuillez patienter le temps que nous transferons votre conscience dans le jeu.")
	time.Sleep(1 * time.Second)
	fmt.Println("importation des données...")
	for b := 0; b <= 100; b++ {
		Clear()
		fmt.Println("Veuillez patienter le temps que nous transferons votre conscience dans le jeu.")
		fmt.Println("importation des données :", b, "%")
		if b < 10 {
			fmt.Println(">.........")
		} else if b < 20 {
			fmt.Println("|>........")
		} else if b < 30 {
			fmt.Println("||>.......")
		} else if b < 40 {
			fmt.Println("|||>......")
		} else if b < 50 {
			fmt.Println("||||>.....")
		} else if b < 60 {
			fmt.Println("|||||>....")
		} else if b < 70 {
			fmt.Println("||||||>...")
		} else if b < 80 {
			fmt.Println("|||||||>..")
		} else if b < 90 {
			fmt.Println("||||||||>.")
		} else if b < 99 {
			fmt.Println("|||||||||>")
		} else if b == 100 {
			fmt.Println("||||||||||")
		}
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("chargement de la fluctlit : utilisateur #7851")
	time.Sleep(1 * time.Second)
	fmt.Println("----------/Syncronisation terminé/----------")
	time.Sleep(2 * time.Second)
	Clear()
	g := 0
	a := 0
	for g == 0 {
		var interaction string
		fmt.Println("Comment vous appelez vous ? (seulement les lettres sont acceptées, tout le reste sera supprimé)")
		fmt.Scanf("%s\n", &interaction)
		res := []rune(interaction)
		for i := range res {
			if res[i] >= 65 && res[i] <= 90 || res[i] >= 97 && res[i] <= 122 {
				if a > 0 && res[i] >= 'A' && res[i] <= 'Z' {
					res[i] += 32
				}
				for a <= 0 {
					if res[i] >= 'a' && res[i] <= 'z' {
						res[i] -= 32
						a++
					} else if res[i] >= 'A' && res[i] <= 'Z' {
						a++
					}
				}
			}
			if res[i] > 90 && res[i] < 97 || res[i] < 65 || res[i] > 122 {
				res[i] = 0
			}
			g++
		}
		C.name = string(res)
	}
	C.start2()
}

//   Création du personnage : classe   //
func (C *character) start2() {
	var interaction string
	a := 0
	for {
		if a <= 0 {
			fmt.Println("choisissez votre classe : 'Humain', 'Elfe', 'Nain'")
			fmt.Scanf("%s\n", &interaction)

			switch interaction {
			case "Humain":
				C.class = "Humain"
				C.maxlife = 100
				C.manamax = 50
				C.mana = 45
				C.initiative = 2
				a++
			case "Elfe":
				C.class = "Elfe"
				C.maxlife = 80
				C.manamax = 70
				C.mana = 65
				C.initiative = 1
				a++
			case "Nain":
				C.class = "Nain"
				C.maxlife = 120
				C.manamax = 30
				C.mana = 25
				C.initiative = 3
				a++
			}
		} else {
			fmt.Println("choisissez votre classe : 'Humain', 'Elfe', 'Nain'")
		}
		if C.maxlife >= 80 && a > 0 {
			Clear()
			C.life = C.maxlife / 2
			slow("------------------------------------------------\n")
			time.Sleep(2 * time.Second)
			slow("Après la rupture du pacte 'Synthesis' entre les ")
			fmt.Print(C.class, "s")
			slow(" et les forces du Dark Territory,\n")
			time.Sleep(2 * time.Second)
			slow("les gobelins ont commencé a attaquer l'Yggdrasil, l'arbre donnant sur votre monde...\n")
			time.Sleep(2 * time.Second)
			slow("Avec pour objectif de semer la destruction et la terreur, le roi des gobelins Oryx, érigea une armée destinée a punir les responsables de son exil.\n")
			time.Sleep(2 * time.Second)
			slow("Soucieux du danger que cela représentait, vous aviez pour ambition de trouver la gemme du pouvoir,\n")
			time.Sleep(2 * time.Second)
			slow("une ancienne relique des Walkers, peuple des temps anciens a l'origine du pacte Synthesis.\n")
			time.Sleep(2 * time.Second)
			slow("Après de nombreuses recherche dans le temple perdu, vous avez finalement mit la main sur la gemme du pouvoir,\n")
			time.Sleep(2 * time.Second)
			slow("mais en touchant cette dernière rien ne se produisit.\n")
			time.Sleep(2 * time.Second)
			slow("A l'heure où je vous parle, les gobelins viennent de franchir l'yggdrasil et vous ressentez une vive douleur vous parcourir le corps,\n")
			time.Sleep(2 * time.Second)
			slow("vous entendez alors une voix vous parler...\n")
			time.Sleep(2 * time.Second)
			slow("'???' : ")
			fmt.Print(C.name)
			slow(" ton destin est d'éliminer le roi gobelin, lui seul peut controler ses subordonnés, une fois éliminé, la guerre cessera.\n")
			time.Sleep(2 * time.Second)
			slow("------------------------------------------------\n")
			time.Sleep(2 * time.Second)
			slow("Bonjour héros, je me présente, je suis GLaDOS et je serai votre guide durant votre aventure.\n")
			time.Sleep(2 * time.Second)
			slow("Je vous donne quelques objets histoire de vous souhaiter la bienvenu, tâchez de regarder votre inventaire pour les utilisez.\n")
			time.Sleep(2 * time.Second)
			slow("Laissez moi vous aider, pour interagir avec la console, il vous suffit d'écrire le mot qui correspond a vos attentes :)\n")
			time.Sleep(2 * time.Second)
			break
		}
	}
}
