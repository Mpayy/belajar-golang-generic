package belajar_golang_generic

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, int64(100), ExperimentalMin(int64(100), int64(200)))
	assert.Equal(t, float64(100.0), ExperimentalMin(float64(100.0), float64(200.0)))
}

func TestExperimentalMaps(t *testing.T) {
	data1 := map[string]string{
		"nama": "rifaih",
	}

	data2 := map[string]string{
		"nama": "aminu",
	}

	result := maps.Equal(data1, data2)
	assert.False(t, result)

	data3 := map[string]string{
		"nama": "rifaih",
	}

	result2 := maps.Equal(data1, data3)
	assert.True(t, result2)
}

func TestExperimentalSlice(t *testing.T) {
	data1 := []int{1, 2, 3}
	data2 := []int{1, 2, 3}
	data3 := []int{1, 2, 4}

	result := slices.Equal(data1, data2)
	assert.True(t, result)

	result2 := slices.Equal(data1, data3)
	assert.False(t, result2)
}