package main

import "fmt"

type computer struct{}
type celphone struct{}
type smartv struct{}
type device interface {
	showme() string
}

func (computer) showme() string {
	return "Soy un computador"
}
func (celphone) showme() string {
	return "Soy un celular"
}
func (smartv) showme() string {
	return "Soy un tv inteligente"
}

func info(d device) {
	fmt.Println(d.showme())
}

func main() {
	pc := computer{}
	info(pc)
	cel := celphone{}
	info(cel)
	tv := smartv{}
	info(tv)
}
