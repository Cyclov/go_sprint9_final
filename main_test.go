package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestRandomGenerator(t *testing.T) {

	var randomGeneratortests = []struct {
		name string
		in   int
		out  any
	}{
		{"Отрицательное число", -1, nil},
		{"Ноль", 0, nil},
		{"Очень маленький слайс", 10, 10},
		{"Маленький слайс", 100, 100},
		{"Нормальный слай", 1000, 1000},
		{"Большой слайс", 100_000, 100_000},
		{"Очень большой слайс", 1000_000, 100_0000},
	}

	for _, tt := range randomGeneratortests {

		t.Run(tt.name, func(t *testing.T) {

			randomSlice := generateRandomElements(tt.in)

			if tt.out == nil {
				assert.Nil(t, randomSlice, "generateRandomElements(%d), должен вернуть nil", tt.in)
				return
			}

			assert.NotNil(t, randomSlice, "generateRandomElements(%d), должен вернуть слайс 1+", tt.in)
			assert.Equal(t, tt.out, len(randomSlice), "размеры слайсов не равны", tt.in)

		})
	}
}

func TestMaximum(t *testing.T) {

	var testMaximumTests = []struct {
		name string
		in   []int
		out  int
	}{
		{"Пустой слайс", []int{}, 0},
		{"Слайс с одним элементов", []int{1}, 1},
		{"Просто слайсл", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}

	for _, tt := range testMaximumTests {

		t.Run(tt.name, func(t *testing.T) {

			max := maximum(tt.in)

			assert.Equal(t, max, tt.out, "Получено неверное значение")

		})
	}
}
