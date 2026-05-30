package belajar_golang_generic

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

// Soal 1 — Type Parameter dasar
// Buat generic function Contains[T comparable](slice []T, item T) bool — cek apakah item ada di slice, works untuk tipe apapun.
// Contains([]int{1,2,3}, 2)        → true
// Contains([]string{"a","b"}, "c") → false

func Contains[T comparable](slice []T, item T) bool {
	// for _,v := range slice {
	// 	if v == item {
	// 		return true
	// 	}
	// }
	// return false

	return slices.Contains(slice, item)
}

// Soal 2 — Type Constraint
// Buat type constraint Number yang mencakup int, int32, int64, float32, float64. Lalu buat function Sum[T Number](angka []T) T yang jumlahkan semua elemen.

type Numbers interface {
	int | int32 | int64 | float32 | float64
}

func Sum[T Numbers](angka []T) T {
	var total T
	for _,v := range angka {
		total += v
	}
	return total
}

// Soal 3 — Generic Struct
// Buat generic struct Pair[A, B any] yang menyimpan dua nilai berbeda tipe. Tambahkan method Swap() Pair[B, A] yang menukar posisi keduanya.
// p := Pair[string, int]{"umur", 26}
// p.Swap() → Pair[int, string]{26, "umur"}

type Pair[A, B any] struct {
	First A
	Second B
}

func (p *Pair[A, B]) Swap() Pair[B, A] {
	return Pair[B, A]{
		p.Second,
		p.First,
	}
}

// Soal 4 — Generic yang berguna di real code
// Buat function Map[T, U any](slice []T, fn func(T) U) []U — transform setiap elemen slice dengan function yang diberikan.
// Contoh penggunaan:
// angka := []int{1, 2, 3, 4, 5}
// hasil := Map(angka, func(n int) string {
//     return strconv.Itoa(n * 2)
// })
// hasil: ["2", "4", "6", "8", "10"]

func Map[T, U any](slice []T, fn func(T) U) []U {
	var result []U
	for _, a := range slice {
		result = append(result, fn(a))
	}
	return result
}

func TestLatihan1(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3}, 2))
	assert.False(t, Contains([]string{"a", "b"}, "c"))

	assert.Equal(t, 6, Sum([]int{1, 2, 3}))

	p := &Pair[string, int]{"umur", 26}
	assert.Equal(t, Pair[int, string]{26, "umur"}, p.Swap() )

	angka := []int{1, 2, 3, 4, 5}
	hasil := Map(angka, func(n int) string {
		return strconv.Itoa(n*2)
	})

	assert.Equal(t, []string{"2", "4", "6", "8", "10"}, hasil)
}