package toolkit

import "fmt"

var Option int

type Tools struct {
	name string
}

func (t *Tools) SetName(n string) {
	t.name = n
}

func (t *Tools) GetName() string {
	return t.name
}

func Print_number() {
	a := 5
	fmt.Println("Hellow there!", a)
}

func PrintOption() {

	fmt.Println("Option: ", Option)
}
