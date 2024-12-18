interface Game {
    uuid: string,
    createdAt: Date,
    updatedAt: Date,
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