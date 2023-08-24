package logic

import "fmt"

const dotEmpty = "•"
const X = "X"
const Y = "Y"

const Size = 3

type Map struct { //игровая доска
	board [][]string
}

type MapInterface interface {
	InitMap(Size int)
	PrintMap()
	IsGameFinished() bool
	CheckWin(playerSymbol string) bool
	IsMapFool() bool
	HumanTurn()
	IsCellValidHuman(x, y int) bool
	IsCellValidPC(x, y int) bool
	ComputerAmatteurModeTurn()
}

func (m *Map) HumanTurn() {
	var x, y int
	for {
		fmt.Println("Введите ход по горизонтали от 1 до 3 и нажмите enter. Затем введите ход по вертикали от 1 до 3 и нажмите enter.")
		fmt.Scan(&x, &y)
		x = x - 1
		y = y - 1
		if m.IsCellValidHuman(x, y) {
			break
		}
	}
	m.board[x][y] = X
	fmt.Scanln()
}

func (m *Map) IsCellValidHuman(x, y int) bool {
	if !(x >= 0 && x <= 2) || !(y >= 0 && y <= 2) {
		fmt.Println("Введены некорректные данные. Введите ход по горизонтали от 1 до 3 и нажмите enter. Затем введите ход по вертикали от 1 до 3 и нажмите enter.")
		return false
	} else if m.board[x][y] != dotEmpty {
		fmt.Println("Выбранная ячейка уже занята. Пожалуйста, выберите другую ячейку")
		return false
	}

	return true
}

func (m *Map) ComputerAmatteurModeTurn() {
	var x, y int

	// Проверяем, есть ли возможность выигрыша компьютера
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if m.IsCellValidPC(i, j) {
				m.board[i][j] = Y
				if m.CheckWin(Y) {
					fmt.Printf("Компьютер выбрал ячейку %d %d\n", i+1, j+1)
					return
				}
				m.board[i][j] = dotEmpty
			}
		}
	}

	// Проверяем, есть ли возможность выигрыша человека и если есть - блокируем
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if m.IsCellValidPC(i, j) {
				m.board[i][j] = X
				if m.CheckWin(X) {
					m.board[i][j] = Y
					fmt.Printf("Компьютер выбрал ячейку %d %d\n", i+1, j+1)
					return
				}
				m.board[i][j] = dotEmpty
			}
		}
	}

	for {
		x, y = RandomChoice(Size)
		if m.IsCellValidPC(x, y) {
			break
		}
	}
	fmt.Printf("Компьютер выбрал ячейку %d %d \n", x+1, y+1)
	m.board[x][y] = Y
}

func (m *Map) IsCellValidPC(x, y int) bool {
	if !(x >= 0 && x <= 2) || !(y >= 0 && y <= 2) {
		return false
	} else if m.board[x][y] != dotEmpty {
		return false
	}

	return true
}

/*
func (m *Map) ComputerSillyModeTurn() {
	var x, y int
	for {
		x, y = RandomChoice(size)
		if m.IsCellValidPC(x, y) {
			break
		}
	}
	fmt.Printf("Компьютер выбрал ячейку %d %d \n", x+1, y+1)
	m.board[x][y] = Y
}
*/

func (m *Map) IsGameFinished() bool {
	m.PrintMap()
	if m.CheckWin(X) {
		fmt.Println("Победил человек!")
		return true
	} else if m.CheckWin(Y) {
		fmt.Println("Победил компьютер!")
		return true
	} else if m.IsMapFool() {
		fmt.Println("Ничья")
		return true
	}
	return false
}

func (m *Map) CheckWin(playerSymbol string) bool {
	for _, comb := range Combinations {
		match := true
		for _, coord := range comb.coords {
			if m.board[coord[0]][coord[1]] != playerSymbol {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func (m *Map) InitMap(Size int) { // Метод для создания игровой доски
	m.board = make([][]string, Size)
	for i := 0; i < Size; i++ {
		m.board[i] = make([]string, Size)
		for j := 0; j < Size; j++ {
			m.board[i][j] = dotEmpty
		}
	}
}

func (m *Map) PrintMap() { // выводим карту

	// Выводим номера столбцов
	fmt.Printf("   ")
	for i := 0; i < len(m.board); i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()

	// Выводим значения по вертикали и горизонтали
	for i, row := range m.board {
		fmt.Printf("%d  ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func (m *Map) IsMapFool() bool {
	for _, row := range m.board {
		for _, cell := range row {
			if cell == dotEmpty {
				return false
			}
		}
	}
	return true
}
