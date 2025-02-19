package db

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type GameBase struct {
	Name       string     `json:"name"`
	Board      [][]string `json:"board"`
	Difficulty string     `json:"difficulty"`
}

type Game struct {
	GameBase

	Uuid      string `json:"uuid"`
	CreatedAt MyTime `json:"createdAt"`
	UpdatedAt MyTime `json:"updatedAt"`
	GameState string `json:"gameState"`
}

func (g *GameBase) UnmarshalJSON(arr []byte) error {
	dst := make(map[string]any)
	err := json.Unmarshal(arr, &dst)
	if err != nil {
		return errors.New("Invalid JSON object")
	}

	if val, ok := dst["name"]; ok {
		g.Name, ok = val.(string)
		if !ok {
			return errors.New("JSON 'name' field is not string")
		}
	} else {
		return errors.New("JSON has no 'name' field")
	}

	if val, ok := dst["difficulty"]; ok {
		g.Difficulty, ok = val.(string)
		if !ok {
			return errors.New("JSON 'difficulty' field is not string")
		}
	} else {
		return errors.New("JSON has no 'difficulty' field")
	}

	if val, ok := dst["board"]; ok {
		g.Board, err = convertAnyTo2DSlice[string](val)
		if err != nil {
			return errors.New("JSON 'board' field is not 2D array of string")
		}
	} else {
		return errors.New("JSON has no 'board' field")
	}

	delete(dst, "name")
	delete(dst, "board")
	delete(dst, "difficulty")

	if len(dst) == 0 {
		return nil
	}

	names := bytes.NewBufferString("sent JSON object has too many attributes:")

	for i := range dst {
		names.WriteString(" '" + i + "'" + ",")
	}

	return errors.New(names.String())
}

type MyTime struct {
	time.Time
}

func (t MyTime) String() string {
	return t.UTC().Format(time.RFC3339Nano)
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.UTC().Format(time.RFC3339Nano))
}

func convertAnyTo2DSlice[T any](a any) (result [][]T, err error) {
	ar, ok := a.([]any)
	if !ok {
		return result, errors.New("Not a slice")
	}

	for i := 0; i < len(ar); i++ {
		src, ok := ar[i].([]any)
		if !ok {
			return result, errors.New("Not 2D slice")
		}

		row := make([]T, 0, len(src))
		for _, v := range src {
			s, ok := v.(T)
			if !ok {
				return result, errors.New("Element error")
			}

			row = append(row, s)
		}

		result = append(result, row)
	}

	return result, nil
}

type UserBase struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Elo      int    `json:"elo"`
	Email    string `json:"email"`
}

type User struct {
	UserBase

	Uuid      string `json:"uuid"`
	CreatedAt MyTime `json:"created_at"`
	Wins      int    `json:"wins"`
	Draws     int    `json:"draws"`
	Losses    int    `json:"losses"`
	Banned    bool   `json:"-"`
}

func (b *UserBase) UnmarshalJSON(arr []byte) error {
	dst := make(map[string]any)
	err := json.Unmarshal(arr, &dst)
	if err != nil {
		return errors.New("Invalid JSON object")
	}

	if val, ok := dst["username"]; ok {
		b.Username, ok = val.(string)
		if !ok {
			return errors.New("JSON 'username' field is not string")
		}
	} else {
		return errors.New("JSON has no 'username' field")
	}

	if val, ok := dst["password"]; ok {
		b.Password, ok = val.(string)
		if !ok {
			return errors.New("JSON 'password' field is not string")
		}
	} else {
		return errors.New("JSON has no 'password' field")
	}

	if val, ok := dst["email"]; ok {
		b.Email, ok = val.(string)
		if !ok {
			return errors.New("JSON 'email' field is not string")
		}
	} else {
		return errors.New("JSON has no 'email' field")
	}

	if val, ok := dst["elo"]; ok {
		elo, ok := val.(float64)
		b.Elo = int(elo)
		if !ok {
			return errors.New("JSON 'elo' field is not int")
		}
	} else {
		return errors.New("JSON has no 'elo' field")
	}

	delete(dst, "username")
	delete(dst, "password")
	delete(dst, "elo")
	delete(dst, "email")

	if len(dst) == 0 {
		return nil
	}

	names := bytes.NewBufferString("Sent JSON object has too many attributes:")

	for i := range dst {
		names.WriteString(" '" + i + "'" + ",")
	}

	return errors.New(names.String())
}
