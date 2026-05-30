package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// comparable adalah constraint untuk tipe data yang bisa dibandingkan menggunakan operator == dan !=
// T bisa diganti dengan tipe data apa saja asalkan bisa dibandingkan
func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	result := IsSame[int](10, 10)
	assert.True(t, result)

	resultString := IsSame[string]("rifai", "rifai")
	assert.True(t, resultString)
}
