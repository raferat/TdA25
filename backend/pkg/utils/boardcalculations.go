package utils

func CalculateGameState([][]string) string {
	return "unknown"
}

// reutrns true, if the board is OK, string is error message
func CheckIntegrity(board [][]string) (bool, string) {
	if !checkSize(board) {
		return false, "Board is not 15x15"
	}

	if !checkContent(board) {
		return false, "Board contains unrecognizable characters"
	}

	if !checkMoves(board) {
		return false, "Board contains invalid number of player moves"
	}

	return true, ""
}

func checkMoves(board [][]string) bool {
	x := 0
	o := 0
	for _, r := range board {
		for _, c := range r {
			if len(c) == 0 {
				continue
			}
			if c == "X" {
				x++
			} else {
				o++
			}
		}
	}

	return x == o || x == o+1
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
