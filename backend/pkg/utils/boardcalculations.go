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

func CanWinInNextTurn(board [][]string, x, o int) bool {
	return false
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
