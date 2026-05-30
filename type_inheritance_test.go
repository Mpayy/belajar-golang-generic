package belajar_golang_generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresident() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresident() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "rifaih", GetName[Manager](&MyManager{Name: "rifaih"}))
	assert.Equal(t, "rifaih", GetName[VicePresident](&MyVicePresident{Name: "rifaih"}))
}