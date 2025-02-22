package utils

import "math"

func expectedGameResult(oldElo int, opponentElo int) float64 {
	return 1 / (1 + math.Pow(10, (float64(opponentElo)-float64(oldElo))/400))
}

// users stats are from before the game
func CalculateNewElo(oldElo int, opponentElo int, gameResult float64, wins int, draws int, losses int) int {
	if gameResult == 1.0 {
		wins++
	} else if gameResult == 0.5 {
		draws++
	} else {
		losses++
	}
	res := float64(oldElo)
	res += 40 * (gameResult - expectedGameResult(oldElo, opponentElo)) * (1 + 0.5*(0.5-float64(wins+draws)/float64(wins+draws+losses)))

	return int(float64(math.Ceil(res)))
}
