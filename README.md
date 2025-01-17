# TdA25 tým `Pentacosihexacontahexadecahecatondodecaexon`

## Technologie
- backend: Go - standartní net/http
- frontend: SvelteKit - statický adapter a fallback
- databáze: SQLite - driver mattn/go-sqlite3

## Požadavky
- [Go](https://go.dev) (min. 1.22)
- [Docker](https://docker.com) (na odeslání)
- [Node.js a NPM](https://nodejs.org)

## Setup lokálního dev prostředí

Společné příkazy
```bash
git clone https://github.com/raferat/TdA25.git
cd TdA25
```

### Backend
```bash
go work init && go work use ./backend
go run tdaserver
```

### Frontend
```bash
ln -s website frontend/build
cd frontend
npm install
npm run dev
```

### Sestavení aplikace
```bash
go build tdaserver
cd frontend && npm run build
```

### Spuštění
```bash
PORT=:4242 ./tdaserver
```
Server běží na adrese [127.0.0.1:4242](http://127.0.0.1:4242)

## Docker
### Sestavení
```bash
docker build . -t tda25
```

### Spuštění
```bash
docker run -p 8080:80 -t tda25
```

Server běží na adrese [127.0.0.1:8080](http://127.0.0.1:8080)

## Struktura

- `backend/` - kořenová složka serveru (Go)
- `frontend/` - kořenová složka SvelteKit

### Algoritmus počítání gameState
Algoritmus se nachází v `backend/pkg/utils/boardcalculations.go`. Je to poměrně naivně udělané. Pro každé prázdné pole se spočítá, zda-li vyhraje, když se na něj položí symbol.
Hráč, který je aktuálně na řadě má možnost vždy vyhrát, hráč po něm může být jedenkrát zablokován. To řeším tak, že od možných výher odečtu 1.

