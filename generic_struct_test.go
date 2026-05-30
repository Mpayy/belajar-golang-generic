package belajar_golang_generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T { //method tidak bisa menggunakan generic
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	// untuk struct tidak mendukung type inference jadi harus ditulis type parameternya
	data := Data[string]{
		First:  "Achmad",
		Second: "Rifaih",
	}

	fmt.Println(data)

	assert.Equal(t, "Hello Budi", data.SayHello("Budi"))
	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
}