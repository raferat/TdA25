# TdA25 tým `ještě nevíme`

## Požadavky
- [Go](https://go.dev) (min. 1.22)
- [Docker](https://docker.com) (na odeslání)

## Setup lokálního dev prostředí
```bash
git clone https://github.com/raferat/TdA25.git
cd TdA25 
go work init && go work use ./backend
go run tdaserver
```

### Sestavení aplikace
```bash
go build tdaserver
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
