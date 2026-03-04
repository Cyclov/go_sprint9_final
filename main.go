package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	if size <= 0 { //Обреботаем крайние
		return nil
	}

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {

	if len(data) < 1 {
		return 0
	}

	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	if CHUNKS < 1 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	if len(data) < 1 {
		return 0
	}

	var wg sync.WaitGroup

	maxInChanks := make([]int, CHUNKS)
	chunkSize := SIZE / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		chunkStart := chunkSize * i
		chunkEnd := min(chunkStart+chunkSize, len(data)) //проверяем чтобы не вылезти за границы массива.

		wg.Add(1)
		go func(data []int, i int) {
			defer wg.Done()

			maxInChanks[i] = maximum(data)

		}(data[chunkStart:chunkEnd], i)

	}
	wg.Wait()
	return maximum(maxInChanks)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(slice)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()
	max = maxChunks(slice)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
