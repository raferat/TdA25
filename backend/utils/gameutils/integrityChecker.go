package gameutils

func (c *MandatoryGameInfo) HasEmptyField() (bool, string) {
  if c.Name == *new(string) {
    return false, "name"
  } else if c.Board == nil || len(c.Board) == 0 {
    return false, "board"
  } else if c.Difficulty == *new(string) {
    return false, "difficulty"
  }

  return true, ""
}
