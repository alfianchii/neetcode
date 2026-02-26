package main

func IsValidSudoku(board [][]byte) bool {
	currRow := make(map[byte]int)
	currColumn := make(map[byte]int)

	for _, rows := range board {
		for _, row := range rows {
			_, exists := currRow[row]
			if exists && row != '.' {
				return false
			}
			currRow[row]++
		}

		currRow = make(map[byte]int)
	}

	for i, rows := range board {
		for j := range rows {
			column := board[j][i]
			_, exists := currColumn[column]
			if exists && column != '.' {
				return false
			}
			currColumn[column]++
		}

		currColumn = make(map[byte]int)
	}

	countColumn := 0
	for i := 0; i < len(board); i++ {
		rows := board[i]

		for j := countColumn; j < countColumn+3; j++ {
			row := rows[j]
			_, exists := currRow[row]
			if exists && row != '.' {
				return false
			}
			currRow[row]++
		}

		if i%3 == 2 {
			currRow = make(map[byte]int)
			if countColumn < 6 {
				i = -1
				countColumn += 3
			}
		}
	}

	return true
}
