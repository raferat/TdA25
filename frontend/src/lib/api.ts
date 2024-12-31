export type Board = ("X" | "O" | "")[][];
export type Difficulty = "beginner" | "easy" | "medium" | "hard" | "extreme";
export type GameState = "opening" | "midgame" | "endgame" | "unknown";

export interface Game {
    uuid: string,
    createdAt: string,
    updatedAt: string,
    name: string,
    difficulty: Difficulty,
    gameState: GameState,
    board: Board,
}

export interface GameBase {
    name: string,
    difficulty: Difficulty,
    board: Board,
}

export interface ApiError {
    code: number,
    message: string,
}

export function defaultGameBase(): GameBase {
    return {
        name: "Untitled",
        difficulty: "beginner",
        board: Array.from({length: 15}, () => Array(15).fill('')),
    };
}

export async function listGames(fetchFunc=fetch): Promise<[Game[], ApiError | undefined]> {
    const res = await fetchFunc("/api/v1/games");

    let json;

    try {
        json = await res.json();
    } catch (e) {
        json = {code: res.status, message: "Not json response"};
    }

    if (json["code"] === undefined) {
        return [json, undefined]
    }
    return [[], json]
}

export async function findGame(uuid: string, fetchFunc=fetch): Promise<[Game, ApiError | undefined]> {
    const res = await fetchFunc(`/api/v1/games/${uuid}`);

    let json;

    try {
        json = await res.json();
    } catch (e) {
        json = {code: res.status, message: "Not json response"};
    }

    if (json["code"] === undefined) {
        return [json, undefined]
    }
    return [json, json]
}

export async function deleteGame(uuid: string, fetchFunc=fetch): Promise<void> {
    await fetchFunc(`/api/v1/games/${uuid}`, {
        method: 'DELETE'
    });
}

export async function createGame(base: GameBase, fetchFunc=fetch): Promise<[Game, ApiError | undefined]> {
    const result = await fetchFunc('/api/v1/games', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(base),
    })

    const json = await result.json();

    if (json["code"] === undefined) {
        return [json, undefined]
    }

    return [json, json]
}

export async function updateGame(base: GameBase, uuid: string, fetchFunc=fetch): Promise<[Game, ApiError | undefined]> {
    const result = await fetchFunc(`/api/v1/games/${uuid}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(base),
    })

    const json = await result.json();

    if (json["code"] === undefined) {
        return [json, undefined]
    }

    return [json, json]
}