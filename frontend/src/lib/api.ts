interface Game {
    uuid: string,
    createdAt: string,
    updatedAt: string,
    name: string,
    difficulty: string,
    gameState: string,
    board: string[][],
}

interface GameBase {
    name: string,
    difficulty: string,
    board: string[][],
}

export async function listGames(): Promise<Game[]> {
    const res = await fetch("/api/v1/games");
    const json = await res.json()

    return json
}

export async function findGame(uuid: string): Promise<Game> {
    const res = await fetch(`/api/v1/games/${uuid}`);
    const json = await res.json()

    return json
}

export async function createGame(base: GameBase): Promise<Game> {
    const result = await fetch('/api/v1/games', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(base),
    })

    return await result.json();
}