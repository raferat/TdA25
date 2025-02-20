package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"strings"
	"time"

	"tdaserver/pkg/db"
	"tdaserver/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var hmacSecret [500]byte

func init() {
	rand.Read(hmacSecret[:])
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginAdmin(w http.ResponseWriter, r *http.Request, password string) {
	if password != "StudentCyberGames25!" {
		w.WriteHeader(404)
		w.Write([]byte("User not found"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "TdA",
		"email":    "tda@scg.cz",
		"elo":      42,
		"isAdmin":  true,
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(hmacSecret[:])
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tokenString))
}

func Login(w http.ResponseWriter, r *http.Request) {
	loginReq := &UserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(loginReq)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	if loginReq.Username == "TdA" || loginReq.Username == "tda@scg.cz" {
		loginAdmin(w, r, loginReq.Password)
		return
	}

	user := db.LoginUser(loginReq.Username, loginReq.Password)
	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("User not found"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"elo":      user.Elo,
		"isAdmin":  false,
		"iat":      time.Now().Unix(),
	})

	tokenString, err := token.SignedString(hmacSecret[:])
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tokenString))
}

// ============================= Freeplay games ================================
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand2.Intn(len(letterBytes))]
	}
	return string(b)
}

type FreeplayGame struct {
	X     *websocket.Conn
	O     *websocket.Conn
	Board [][]string
	Moves int
}

var freeplayGames = make(map[string]*FreeplayGame)

func startFreeplayGame(id string) {
	conns := freeplayGames[id]

	defer func() {
		conns.X.Close()
		conns.O.Close()
	}()

	conns.X.WriteMessage(websocket.TextMessage, []byte("start X"))
	conns.O.WriteMessage(websocket.TextMessage, []byte("start O"))

	xwon, owon := false, false

	for conns.Moves < 15*15 {
		for {
			_, msg, err := conns.X.ReadMessage()
			if err != nil {
				return
			}

			var (
				x int
				y int
			)

			num, err := fmt.Sscanf(string(msg), "%d %d", &x, &y)
			if err != nil || num != 2 {
				conns.X.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			if x < 0 || x > 14 || y < 0 || y > 14 || conns.Board[x][y] != "" {
				conns.X.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			conns.Board[x][y] = "X"
			conns.O.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("placeX %d %d", x, y)))
			break
		}

		conns.Moves++

		xwon = utils.IsGameFinished(conns.Board)
		if xwon {
			break
		}

		for {
			_, msg, err := conns.O.ReadMessage()
			if err != nil {
				return
			}

			var (
				x int
				y int
			)

			num, err := fmt.Sscanf(string(msg), "%d %d", &x, &y)
			if err != nil || num != 2 {
				conns.O.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			if x < 0 || x > 14 || y < 0 || y > 14 || conns.Board[x][y] != "" {
				conns.O.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			conns.Board[x][y] = "O"
			conns.X.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("placeO %d %d", x, y)))
			break
		}

		conns.Moves++
		owon = utils.IsGameFinished(conns.Board)
		if owon {
			break
		}

	}

	if xwon {
		conns.X.WriteMessage(websocket.TextMessage, []byte("Xwon"))
		conns.O.WriteMessage(websocket.TextMessage, []byte("Xwon"))
	} else if owon {
		conns.X.WriteMessage(websocket.TextMessage, []byte("Owon"))
		conns.O.WriteMessage(websocket.TextMessage, []byte("Owon"))
	} else {
		conns.X.WriteMessage(websocket.TextMessage, []byte("draw"))
		conns.O.WriteMessage(websocket.TextMessage, []byte("draw"))
	}
}

func HandleFreeplayGame(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		w.Write([]byte("Websocket upgrade needed"))
		return
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WS freeplay: %#v\n", err)
		return
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading WS message: %#v\n", err)
	}

	args := strings.Split(strings.TrimSpace(string(data)), " ")
	if args[0] == "create" {
		id := ""
		for {
			id = randStringBytes(6)
			if _, ok := freeplayGames[id]; ok {
				continue
			}

			break
		}

		board := make([][]string, 0, 15)
		for range 15 {
			board = append(board, make([]string, 15))
		}

		freeplayGames[id] = &FreeplayGame{
			Board: board,
		}

		if args[1] == "X" {
			freeplayGames[id].X = conn
		} else {
			freeplayGames[id].O = conn
		}

		conn.WriteMessage(websocket.TextMessage, []byte("code "+id))

	} else if args[0] == "connect" {
		if freeplayGames[args[1]].O == nil {
			freeplayGames[args[1]].O = conn
		} else {
			freeplayGames[args[1]].X = conn
		}

		startFreeplayGame(args[1])
	}
}
