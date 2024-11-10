# TdA25 tým `ještě nevíme`

## Požadavky
- [Go](https://go.dev) (min. 1.22)
- [Docker](https://docker.com) (na odeslání)
- [Postgres](https://www.postgresql.org/) (na lokální testovní)

## Setup lokálního prostředí
```bash
git clone https://github.com/raferat/TdA25.git
cd TdA25 
go work init && go work use ./backend
sh ./postgresql.conf.d/init.sh
```

### Sestavení aplikace
```bash
go build tdaserver
```
### Spuštění
```bash
./tdaserver
```

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
- `frontend/` - kořenová složka klienta (SvelteKit; kompilován na statickou stránku)
- `frontend/build/` - sestavěná stránka od SvelteKit