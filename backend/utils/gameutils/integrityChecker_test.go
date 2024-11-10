package gameutils_test

import (
	"tdaserver/utils/gameutils"
	"testing"
)

func TestHasEmptyField(t *testing.T) {
  s := &gameutils.MandatoryGameInfo{}
  if ok, val := s.HasEmptyField(); ok || val != "name" {
    t.Fatalf("Struct: %+v, first empty value should be `name`, but returned `%s`\n", s, val)
  }

  s.Name = "Ignac"

  if ok, val := s.HasEmptyField(); ok || val != "board" {
    t.Fatalf("Struct: %+v, second empty value should be `board`, but returned `%s`\n", s, val)
  }

  s.Board = make([][]string, 2)

  if ok, val := s.HasEmptyField(); ok || val != "difficulty" {
    t.Fatalf("Struct: %+v, third empty value should be `difficulty`, but returned `%s`\n", s, val)
  }

  s.Difficulty = "beginner"

  if ok, val := s.HasEmptyField(); !ok || val != "" {
    t.Fatalf("Struct: %+v, should be full, but returned `%s`\n", s, val)
  }
}
