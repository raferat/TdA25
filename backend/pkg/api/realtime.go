package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"slices"
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
	go HandleMatches()
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

func RefreshLogin(w http.ResponseWriter, r *http.Request) {
	tkString := r.Header.Get("Auth")

	user, ok := CheckUser(tkString)
	if !ok {
		w.WriteHeader(401)
		w.Write([]byte("Wrong token"))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "TdA",
		"email":    "tda@scg.cz",
		"elo":      42,
		"isAdmin":  true,
		"iat":      time.Now().Unix(),
	})

	if user != nil {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"email":    user.Email,
			"elo":      user.Elo,
			"isAdmin":  false,
			"iat":      time.Now().Unix(),
		})
	}

	tokenString, err := token.SignedString(hmacSecret[:])
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tokenString))
}

func CheckUser(tokenStr string) (*db.User, bool) {
	token, err := jwt.Parse(tokenStr, func(_ *jwt.Token) (any, error) {
		return hmacSecret[:], nil
	})
	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if isAdmin := claims["isAdmin"].(bool); isAdmin {
			return nil, true
		}

		user := db.SearchUser(claims["username"].(string))
		if user == nil {
			return nil, false
		}

		return user, true
	} else {
		return nil, false
	}
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

//================================================================ Matchmaking ====================================================

type WaitingUser struct {
	User *db.User
	Conn *websocket.Conn
}

var waitQueue []WaitingUser = make([]WaitingUser, 0, 100)

func HandleMatchmakingRequest(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		w.Write([]byte("Websocket upgrade needed"))
		return
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WS matchmaking: %#v\n", err)
		return
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		return
	}

	authMsg := strings.Split(strings.TrimSpace(string(data)), " ")
	if len(authMsg) != 2 || authMsg[0] != "auth" {
		conn.Close()
		return
	}

	user, ok := CheckUser(authMsg[1])

	if !ok || user == nil {
		conn.WriteMessage(websocket.TextMessage, []byte("not-auth"))
		conn.Close()
		return
	}

	if user.Banned {
		conn.WriteMessage(websocket.TextMessage, []byte("You are banned"))
		conn.Close()
		return
	}

	waitQueue = append(waitQueue, WaitingUser{
		User: user,
		Conn: conn,
	})
}

func HandleMatches() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		<-ticker.C

		slices.SortFunc(waitQueue, func(a, b WaitingUser) int {
			if a.User.Elo == b.User.Elo {
				return 0
			} else if a.User.Elo < b.User.Elo {
				return -1
			}

			return 1
		})

		for len(waitQueue) >= 2 {
			go RunMatch(waitQueue[0], waitQueue[1])
			waitQueue = waitQueue[2:]
		}
	}
}

func RunMatch(user1, user2 WaitingUser) {
	defer user1.Conn.Close()
	defer user2.Conn.Close()

	X, O := user1, user2

	if rand2.Float64() >= 0.5 {
		X, O = user2, user1
	}

	X.Conn.WriteMessage(websocket.TextMessage, []byte("start X"))
	O.Conn.WriteMessage(websocket.TextMessage, []byte("start O"))

	xwon, owon := false, false

	moves := 0
	board := make([][]string, 0, 15)
	for range 15 {
		board = append(board, make([]string, 15))
	}

	for moves < 15*15 {
		for {
			_, msg, err := X.Conn.ReadMessage()
			if err != nil {
				return
			}

			var (
				x int
				y int
			)

			num, err := fmt.Sscanf(string(msg), "%d %d", &x, &y)
			if err != nil || num != 2 {
				X.Conn.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			if x < 0 || x > 14 || y < 0 || y > 14 || board[x][y] != "" {
				X.Conn.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			board[x][y] = "X"
			O.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("placeX %d %d", x, y)))
			break
		}

		moves++

		xwon = utils.IsGameFinished(board)
		if xwon {
			break
		}

		for {
			_, msg, err := O.Conn.ReadMessage()
			if err != nil {
				return
			}

			var (
				x int
				y int
			)

			num, err := fmt.Sscanf(string(msg), "%d %d", &x, &y)
			if err != nil || num != 2 {
				O.Conn.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			if x < 0 || x > 14 || y < 0 || y > 14 || board[x][y] != "" {
				O.Conn.WriteMessage(websocket.TextMessage, []byte("badplacement"))
				continue
			}

			board[x][y] = "O"
			X.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("placeO %d %d", x, y)))
			break
		}

		moves++
		owon = utils.IsGameFinished(board)
		if owon {
			break
		}

	}

	xGameResult := 0.5
	if owon {
		xGameResult = 0
	} else if xwon {
		xGameResult = 1
	}

	newXElo := utils.CalculateNewElo(X.User.Elo, O.User.Elo, xGameResult, X.User.Wins, X.User.Draws, X.User.Losses)
	newOElo := utils.CalculateNewElo(O.User.Elo, X.User.Elo, 1-xGameResult, O.User.Wins, O.User.Draws, O.User.Losses)

	if xwon {
		X.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Xwon %d", newXElo)))
		O.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Xwon %d", newOElo)))

		X.User.Wins++
		O.User.Losses++
	} else if owon {
		X.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Owon %d", newXElo)))
		O.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Owon %d", newOElo)))

		X.User.Losses++
		O.User.Wins++
	} else {
		X.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("draw %d", newXElo)))
		O.Conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("draw %d", newOElo)))

		X.User.Draws++
		O.User.Draws++
	}

	db.EditUserScore(X.User.Uuid, newXElo, X.User.Wins, X.User.Draws, X.User.Losses)
	db.EditUserScore(O.User.Uuid, newOElo, O.User.Wins, O.User.Draws, O.User.Losses)
}
