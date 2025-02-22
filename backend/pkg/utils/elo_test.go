package utils_test

import (
	"testing"

	. "tdaserver/pkg/utils"
)

func TestCalculateELO(t *testing.T) {
	if CalculateNewElo(400, 400, 1, 0, 0, 0) != 415 {
		t.Fatal("#1")
	}

	if CalculateNewElo(400, 400, 0, 0, 0, 0) != 375 {
		t.Fatal("#2")
	}

	if CalculateNewElo(400, 400, 0.5, 0, 0, 0) != 400 {
		t.Fatal("#3")
	}
}
