package belajar_golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tanpa Generic kita harus membuat 2 func berbeda untuk setiap tipe data, padahal isinya sama
// func LengthString(param string) string {
// 	fmt.Println(param)
// 	return param
// }

// func LengthInteger(param int) int {
// 	fmt.Println(param)
// 	return param
// }

//Dengan Generic kita cukup membuat 1 func saja, 
// Kita harus mendefinisikan T terlebih dahulu tipe constraintnya apa
// any adalah constraint bahwa tipe datanya bisa apa saja atau sama saja dengan interface{} 
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSampleConstraint(t *testing.T) {
	assert.True(t, true)
}

func TestLength(t *testing.T) {
	var result string = Length[string]("rifaih")
	assert.Equal(t, "rifaih", result)

	var resultNumber int = Length[int](10)
	assert.Equal(t, 10, resultNumber)
}