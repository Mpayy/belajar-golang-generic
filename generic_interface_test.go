package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// kalau mau membuat generic interface, maka implementasi struct juga harus generic
type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(value T) {
	m.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}
	result := ChangeValue(&myData, "Rifaih")
	assert.Equal(t, "Rifaih", result)
	assert.Equal(t, "Rifaih", myData.Value)
}