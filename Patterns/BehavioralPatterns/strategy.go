package main

type Strategy interface {
	Attack()
}

type AK47 struct {
}

func (a *AK47) Attack() {
}

type Knife struct {
}

func (k *Knife) Attack() {
}

type Hero struct {
	st Strategy
}

func (h *Hero) SetWeaponStrategy(st Strategy) {
	h.st = st
}

func (h *Hero) Fight() {
	h.st.Attack()
}

func main() {
	hero := new(Hero)
	hero.SetWeaponStrategy(new(AK47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()

}
