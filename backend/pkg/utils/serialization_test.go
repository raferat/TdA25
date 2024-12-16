package utils_test

import (
	"fmt"
	"testing"

	. "tdaserver/pkg/utils"
)

func TestSmallBoard(t *testing.T) {
	// Define a sample board to test
	originalBoard := [][]string{
		{"X", "O", "X"},
		{"O", "X", "O"},
		{"X", "O", "X"},
	}

	// Serialize the board
	serializedBoard := SerializeBoard(originalBoard)

	// Deserialize the board
	deserializedBoard, err := DeserializeBoard(serializedBoard)
	if err != nil {
		t.Fatalf("Failed to deserialize board: %v", err)
	}

	// Check if the deserialized board matches the original board
	if !equalBoards(originalBoard, deserializedBoard) {
		t.Errorf("Deserialized board does not match original. Got: %v, Want: %v", deserializedBoard, originalBoard)
	}
}

func TestLargeArray(t *testing.T) {
	rows := 10000 // Number of rows
	cols := 5000  // Number of columns

	// Create a 2D slice of strings
	string2DSlice := make([][]string, rows)

	// Initialize the 2D slice with some values
	for i := range string2DSlice {
		string2DSlice[i] = make([]string, cols)
		for j := range string2DSlice[i] {
			string2DSlice[i][j] = fmt.Sprintf("Row%d_Col%d", i, j)
		}
	}

	seSlice := SerializeBoard(string2DSlice)

	// Deserialize the board
	deSlice, err := DeserializeBoard(seSlice)
	if err != nil {
		t.Fatalf("Failed to deserialize board: %v", err)
	}

	// Check if the deserialized board matches the original board
	if !equalBoards(deSlice, string2DSlice) {
		t.Errorf("Deserialized board does not match original. Got: %v, Want: %v", deSlice, string2DSlice)
	}
}

// Helper function to compare two boards
func equalBoards(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
