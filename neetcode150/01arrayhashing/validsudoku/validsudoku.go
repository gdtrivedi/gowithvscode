package main

import "fmt"

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	// same row, all columns
	rm := make(map[byte]struct{})
	// same column, all rows
	cm := make(map[byte]struct{})
	for r := 0; r < 9; r++ {
		clear(rm)
		clear(cm)
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				if _, ok := rm[board[r][c]]; ok {
					return false
				} else {
					rm[board[r][c]] = struct{}{}
				}
			}

			if board[c][r] != '.' {
				if _, ok := cm[board[c][r]]; ok {
					return false
				} else {
					cm[board[c][r]] = struct{}{}
				}
			}
		}
	}

	// squares
	sm := make(map[byte]struct{})
	for cnt := 0; cnt < 9; cnt += 3 {
		clear(sm)

		// first 3 columns
		for c := 0; c < 3; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

		clear(sm)
		// next 3 columns
		for c := 3; c < 6; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

		clear(sm)
		// last 3 columns
		for c := 6; c < 9; c++ {
			if !validSquare(cnt, sm, board, c) {
				return false
			}
		}

	}

	return true
}

func validSquare(cnt int, sm map[byte]struct{}, board [][]byte, c int) bool {
	for r := cnt; r < cnt+3; r++ {
		if board[r][c] != '.' {
			if _, ok := sm[board[r][c]]; ok {
				return false
			} else {
				sm[board[r][c]] = struct{}{}
			}
		}
	}
	return true
}

func isValidSudoku1(board [][]byte) bool {
	tr := len(board)
	tc := len(board[0])

	fmt.Println(tr, tc)

	// row wise check
	for i := 0; i < tr; i++ {
		tra := make([]int, 9)
		for j := 0; j < tc; j++ {
			if board[i][j] == '.' {
				continue
			}
			if tra[board[i][j]-'1'] > 1 {
				fmt.Println("aaaaaaaaa")
				return false
			}
			tra[board[i][j]-'1']++
		}
	}

	// column wise check
	for i := 0; i < tc; i++ {
		tca := make([]int, 9)
		for j := 0; j < tr; j++ {
			if board[i][j] == '.' {
				continue
			}
			if tca[board[i][j]-'1'] > 1 {
				fmt.Println("bbbbbbbbbb")
				return false
			}
			tca[board[i][j]-'1']++
		}
	}

	return true
}
