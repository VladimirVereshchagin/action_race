package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var sharedCounter int // Переменная, доступная из всех горутин
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			sharedCounter += i // Изменение общей переменной, создающее гонку
			fmt.Println(sharedCounter)
		}(i) // Передаем i в качестве аргумента, чтобы избежать гонки на уровне i, но оставить гонку на уровне sharedCounter
	}

	wg.Wait()
}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
