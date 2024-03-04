package main

import (
	"fmt"
	"math"
)
type Test interface {
	Luas() float64
	keliling() float64
}

type Lingkaran struct {
	Diameter float32
}
func (l Lingkaran) JariJari()float32  {
	return l.Diameter / 2
}
func (l Lingkaran) Luas()float64  {
	return math.Pi * math.Pow(float64(l.JariJari()),float64(2))
}
func (l Lingkaran) keliling()float64  {
	return math.Pi * float64(math.Pi)
}
type Persegi_panjang struct {
	Panjang uint
	Lebar uint
}
func (p Persegi_panjang) Luas()float64  {
	return float64(p.Panjang) * float64(p.Lebar)
}
func (p Persegi_panjang) keliling()float64  {
	return float64(2 * (p.Lebar + p.Panjang))
}
func main() {
	var bangunDatar Test
	bangunDatar = Lingkaran{
		Diameter: 10,
	}
	fmt.Println("Luas: ",bangunDatar.Luas())
	fmt.Println("Keliling: ", bangunDatar.keliling())
	fmt.Println("Jari-jari: ",bangunDatar.(Lingkaran).JariJari())
	fmt.Println("Persegi")
	bangunDatar = Persegi_panjang{
		Panjang: 10,
		Lebar: 5,
	}
	fmt.Println("Luas: ",bangunDatar.Luas())
	fmt.Println("Keliling", bangunDatar.keliling())
	fmt.Println("Panjang persegi: ",bangunDatar.(Persegi_panjang).Panjang)
}