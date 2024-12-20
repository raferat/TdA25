export interface Game {
    uuid: string,
    createdAt: string,
    updatedAt: string,
    name: string,
    difficulty: string,
    gameState: string,
    board: string[][],
}

export interface GameBase {
    name: string,
    difficulty: string,
    board: string[][],
}

export function defaultGameBase(): GameBase {
    return {
        name: "Untitled",
        difficulty: "beginner",
        board: Array.from({length: 15}, () => Array(15).fill('')),
    };
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

export async function deleteGame(uuid: string): Promise<void> {
    await fetch(`/api/v1/games/${uuid}`, {
        method: 'DELETE'
    });
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

export async function updateGame(base: GameBase, uuid: string): Promise<Game> {
    const result = await fetch(`/api/v1/games/${uuid}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(base),
    })

    return await result.json();
}