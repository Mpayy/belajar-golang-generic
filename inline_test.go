package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// kita bisa juga lansung membuat type set parameter generic lansung di function nya tanpa harus mendefinisikan type set nya terlebih dahulu
func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(100), FindMin[int64](200, 100))
	assert.Equal(t, float64(100.0), FindMin[float64](200.0, 100.0))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{"rifaih", "aminu", "mita"}
	first := GetFirst(names)
	assert.Equal(t, "rifaih", first)

	numbers := []int{1, 2, 3}
	firstInt := GetFirst(numbers)
	assert.Equal(t, 1, firstInt)
}