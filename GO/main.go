package main

//   importation des bibliotèques   //
import (
	"fmt"
	_ "image/jpeg"
	"time"
)

//   structure personnage   //
type character struct {
	name       string
	class      string
	lvl        int
	xpac       int
	needxp     int
	maxlife    int
	life       int
	manamax    int
	mana       int
	initiative int
	damage     int
	money      int
	slotinv    int
	inventory  []string
	skill      []string
	equip      stuff
}

//   structure équipement   //
type stuff struct {
	head string
	body string
	foot string
	hand string
}

//   Vstructure monstre   //
type Monstre struct {
	name       string
	pvmax      int
	pv         int
	damage     int
	initiative int
}

//   function d'initialisation du personnage   //
func (C *character) Init(name string, class string, lvl int, xpac int, needxp int, maxlife int, life int, manamax int, mana int, initiative int, damage int, money int, slotinv int, inventory []string, skill []string, head string, body string, foot string, hand string) {
	C.name = name
	C.class = class
	C.lvl = lvl
	C.xpac = xpac
	C.needxp = needxp
	C.maxlife = maxlife
	C.life = life
	C.manamax = manamax
	C.mana = mana
	C.initiative = initiative
	C.slotinv = slotinv
	C.inventory = inventory
	C.money = money
	C.skill = skill
	C.equip.head = head
	C.equip.body = body
	C.equip.foot = foot
	C.equip.hand = hand
}

// la fontion ne sert pas car on init automatiquement les monstres dans TrainingFight //
/*func (M *Monstre) init(name string, pvmax int, pv int, damage int, initiative int) {
	M.name = name
	M.pvmax = pvmax
	M.pv = pv
	M.damage = damage
	M.initiative = initiative
}
*/

//   Variable globale   //
var ASCIISTR = "MND8OZ$7I?+=~:,._-"
var a = 0
var Tour = 1
var mort = 0
var Combat int
var popogratuite int
var C *character = &character{}
var C1 character = character{}
var M *Monstre = &Monstre{}
var M1 Monstre = Monstre{}

//   fonction principale   //
func main() {
	C1.Init(C.name, C.class, 1, 0, 1000, C.maxlife, 1, C.manamax, C.mana, C.initiative, 0, 100, 10, []string{"potion de santé", "potion de santé", "potion de santé", "casqu'ette [5 pv]", "armure en carton [10 pv]", "jambière en lin [5 pv]"}, []string{"coup de poing"}, "empty", "empty", "empty", "empty")
	C1.charcreation()
	for {
		if C1.life <= 0 {
			C1.dead()
			return
		} else if Combat >= 3 {
			fmt.Println("Fin du jeu")
			break
		} else {
			C1.menu()
			time.Sleep(2 * time.Second)
			Clear()
		}
	}
}
