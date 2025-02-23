package api

import (
	"encoding/json"
	"net/http"

	"tdaserver/pkg/db"
)

type adminChange struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Elo      int    `json:"elo"`
	Wins     int    `json:"wins"`
	Draws    int    `json:"draws"`
	Losses   int    `json:"losses"`
	Banned   bool   `json:"banned"`
}

func AdminUpdateUserData(w http.ResponseWriter, r *http.Request) {
	tkString := r.Header.Get("Auth")
	user, ok := CheckUser(tkString)

	if user != nil || !ok {
		w.WriteHeader(401)
		return
	}

	var change adminChange

	err := json.NewDecoder(r.Body).Decode(&change)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.AdminEdit(change.Uuid, change.Username, change.Email, change.Elo, change.Wins, change.Draws, change.Losses, change.Banned)
	w.WriteHeader(http.StatusNoContent)
}

func AdminGetData(w http.ResponseWriter, r *http.Request) {
	tkString := r.Header.Get("Auth")
	user, ok := CheckUser(tkString)

	if user != nil || !ok {
		w.WriteHeader(401)
		return
	}

	users := db.ListUsers()

	users2 := make([]*adminUser, 0, 50)
	for _, v := range users {
		tmp := adminUser(*v)
		users2 = append(users2, &tmp)
	}

	json.NewEncoder(w).Encode(users2)
}

type adminUser db.User

func (a *adminUser) MarshalJSON() ([]byte, error) {
	res := map[string]any{
		"uuid":      a.Uuid,
		"createdAt": a.CreatedAt,
		"username":  a.Username,
		"email":     a.Email,
		"elo":       a.Elo,
		"wins":      a.Wins,
		"draws":     a.Draws,
		"losses":    a.Losses,
		"banned":    a.Banned,
		"password":  a.Password,
	}

	return json.Marshal(res)
}
