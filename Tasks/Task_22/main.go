package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(2)
	a.Exp(a, big.NewInt(20), nil).Add(a, big.NewInt(100))
	b := big.NewInt(2)
	b.Exp(b, big.NewInt(20), nil).Add(b, big.NewInt(200))

	// Умножение
	multiplication := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s\n", multiplication.String())

	// Деление
	division := new(big.Rat).SetFrac(a, b)
	fmt.Printf("Деление: %s\n", division.FloatString(2))

	// Сложение
	addition := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s\n", addition.String())

	// Вычитание
	subtraction := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s\n", subtraction.String())

}
