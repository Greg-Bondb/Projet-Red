package main

//   Apprendre un nouveau sort  //
func (C *character) spellBook(sort string, rang int) {
	for i := 0; i < len(C.skill); i++ {
	}
	C.skill = append(C.skill, sort)
	println("Vous avez appris un nouveau sort")
	C.removestuff(rang)
}
