package belajar_golang_generic

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[E any](bag Bag[E]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestPrintBag(t *testing.T) {
	names := Bag[string]{"rifaih", "aminu", "mita"}
	PrintBag[string](names)

	numbers := Bag[int]{1, 2, 3}
	PrintBag(numbers) // ini menggunakan typeinference (type tidak perlu ditulis)

}