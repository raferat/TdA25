package utils

func CalculateGameState(board [][]string) string {
	x, o := countMoves(board)

	if CanWinInNextTurn(board, x, o) {
		return "endgame"
	}

	if o < 5 {
		return "opening"
	}

	return "midgame"
}

type Direction [2]int

type cachePos struct {
	x      int
	y      int
	dir    Direction
	symbol string
}

func countSymbolInDirection(startX, startY int, board [][]string, symbol string, dir Direction, cache map[cachePos]int) int {
	x, y := startX+dir[0], startY+dir[1]

	if val, ok := cache[cachePos{startX, startY, dir, symbol}]; ok {
		return val
	}

	cnt := 0

	for y >= 0 && x >= 0 && y < len(board) && x < len(board[y]) && board[y][x] == symbol {
		x += dir[0]
		y += dir[1]
		cnt++
	}

	if y >= 0 && x >= 0 && y < len(board) && x < len(board[y]) {
		cache[cachePos{x, y, Direction{dir[0] * -1, dir[1] * -1}, symbol}] = cnt
	}

	return cnt
}

func IsGameFinished(board [][]string) bool {
	cache := make(map[cachePos]int)
	for y, r := range board {
		for x, c := range r {
			if c == "" {
				continue
			}

			row := countSymbolInDirection(x, y, board, c, Direction{1, 0}, cache) + countSymbolInDirection(x, y, board, c, Direction{-1, 0}, cache) + 1
			col := countSymbolInDirection(x, y, board, c, Direction{0, 1}, cache) + countSymbolInDirection(x, y, board, c, Direction{0, -1}, cache) + 1
			diag1 := countSymbolInDirection(x, y, board, c, Direction{1, 1}, cache) + countSymbolInDirection(x, y, board, c, Direction{-1, -1}, cache) + 1
			diag2 := countSymbolInDirection(x, y, board, c, Direction{-1, 1}, cache) + countSymbolInDirection(x, y, board, c, Direction{1, -1}, cache) + 1

			if row >= 5 || col >= 5 || diag1 >= 5 || diag2 >= 5 {
				return true
			}
		}
	}
	return false
}

func CanWinInNextTurn(board [][]string, x, o int) bool {
	type pos struct {
		x int
		y int
	}
	xWinOpt := 0
	oWinOpt := 0
	cache := make(map[cachePos]int)

	if x == o {
		oWinOpt--
	} else {
		xWinOpt--
	}

	for y, r := range board {
		for x, c := range r {
			if c != "" {
				continue
			}

			xRow := countSymbolInDirection(x, y, board, "X", Direction{1, 0}, cache) + countSymbolInDirection(x, y, board, "X", Direction{-1, 0}, cache)
			oRow := countSymbolInDirection(x, y, board, "O", Direction{1, 0}, cache) + countSymbolInDirection(x, y, board, "O", Direction{-1, 0}, cache)
			xCol := countSymbolInDirection(x, y, board, "X", Direction{0, 1}, cache) + countSymbolInDirection(x, y, board, "X", Direction{0, -1}, cache)
			oCol := countSymbolInDirection(x, y, board, "O", Direction{0, 1}, cache) + countSymbolInDirection(x, y, board, "O", Direction{0, -1}, cache)
			xDiag1 := countSymbolInDirection(x, y, board, "X", Direction{1, 1}, cache) + countSymbolInDirection(x, y, board, "X", Direction{-1, -1}, cache)
			oDiag1 := countSymbolInDirection(x, y, board, "O", Direction{1, 1}, cache) + countSymbolInDirection(x, y, board, "O", Direction{-1, -1}, cache)
			xDiag2 := countSymbolInDirection(x, y, board, "X", Direction{-1, 1}, cache) + countSymbolInDirection(x, y, board, "X", Direction{1, -1}, cache)
			oDiag2 := countSymbolInDirection(x, y, board, "O", Direction{-1, 1}, cache) + countSymbolInDirection(x, y, board, "O", Direction{1, -1}, cache)

			if xRow >= 4 || xCol >= 4 || xDiag1 >= 4 || xDiag2 >= 4 {
				xWinOpt++
			}
			if oRow >= 4 || oCol >= 4 || oDiag1 >= 4 || oDiag2 >= 4 {
				oWinOpt++
			}
		}
	}
	return xWinOpt > 0 || oWinOpt > 0
}

// returns true when Difficulty is OK
func CheckDifficulty(str string) bool {
	states := []string{
		"beginner",
		"easy",
		"medium",
		"hard",
		"extreme",
	}

	for _, v := range states {
		if v == str {
			return true
		}
	}

	return false
}

// reutrns true, if the board is OK, string is error message
func CheckIntegrity(board [][]string) (bool, string) {
	if !checkSize(board) {
		return false, "Board is not 15x15"
	}

	if !checkContent(board) {
		return false, "Board contains unrecognizable characters"
	}

	if x, o := countMoves(board); x != o && x != o+1 {
		return false, "Board contains invalid number of player moves"
	}

	return true, ""
}

func countMoves(board [][]string) (int, int) {
	x := 0
	o := 0
	for _, r := range board {
		for _, c := range r {
			if c == "X" {
				x++
			} else if c == "O" {
				o++
			}
		}
	}

	return x, o
}

func checkContent(board [][]string) bool {
	for _, x := range board {
		for _, y := range x {
			if len(y) > 1 {
				return false
			}

			if len(y) == 0 {
				continue
			}

			switch y[0] {
			case 'X':
				fallthrough
			case 'O':
				continue

			default:
				return false
			}
		}
	}

	return true
}

func checkSize(board [][]string) bool {
	if len(board) != 15 {
		return false
	}

	for _, v := range board {
		if len(v) != 15 {
			return false
		}
	}

	return true
}
