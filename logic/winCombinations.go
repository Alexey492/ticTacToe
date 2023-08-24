package logic

import (
	"math/rand"
	"time"
)

type Combination struct {
	coords [3][2]int // массив координат комбинации
}

var Combinations = []Combination{
	{coords: [3][2]int{{0, 0}, {1, 1}, {2, 2}}},
	{coords: [3][2]int{{2, 0}, {1, 1}, {0, 2}}},
	{coords: [3][2]int{{0, 0}, {0, 1}, {0, 2}}},
	{coords: [3][2]int{{1, 0}, {1, 1}, {1, 2}}},
	{coords: [3][2]int{{2, 0}, {2, 1}, {2, 2}}},
	{coords: [3][2]int{{0, 0}, {1, 0}, {2, 0}}},
	{coords: [3][2]int{{0, 1}, {1, 1}, {2, 1}}},
	{coords: [3][2]int{{0, 2}, {1, 2}, {2, 2}}},
}

func RandomChoice(size int) (x, y int) {
	// Создаем генератор случайных чисел с использованием текущего времени
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	// Генерируем два случайных числа от 0 до size
	x = generator.Intn(size)
	y = generator.Intn(size)

	return x, y
}
