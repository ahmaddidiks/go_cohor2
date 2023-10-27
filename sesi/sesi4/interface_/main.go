package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("...")
	var lingkaran1 hitung = lingkaran{7}
	var persegi1 hitung = persegi{5}

	fmt.Printf("Type of lingkaran1: %T \n", lingkaran1)
	fmt.Printf("Type of persegi1: %T \n", persegi1)

	fmt.Println("luas lingkaran: ", lingkaran1.keliling())

	fmt.Println("diameter lingkaran: ", lingkaran1.(lingkaran).diameterMethod()) // Interface (type assertion)

}

type hitung interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	jariJari float64
}

func (l lingkaran) luas() float64 {
	return math.Pi * l.jariJari * l.jariJari
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jariJari
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func (l lingkaran) diameterMethod() float64 {
	return l.jariJari * 2
}
