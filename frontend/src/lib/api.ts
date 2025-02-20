import { get, writable, type Writable } from "svelte/store";

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


//================================== Users API ==========================================

export type LoginUser =
    | { username: string; email: string; elo: number; token: string, isAdmin: boolean }
    | undefined;

export const loginState: Writable<LoginUser> = writable();


export function initLoginState() {
    loginState.set(loadLoginState());
    loginState.subscribe((value: LoginUser) => {
        if (!window) {
            return;
        }

        if (!value) {
            window.localStorage.removeItem("user");
            return;
        }

        window.localStorage.setItem("user", JSON.stringify(value));
    });

    window.addEventListener("storage", () => {
        const loaded = loadLoginState();
        if (loaded == get(loginState)) return;
        
        loginState.set(loaded);
    });

    window.addEventListener("userLoginChange", () => {
        const loaded = loadLoginState();
        if (loaded == get(loginState)) return;
        
        loginState.set(loaded);
    });
}

function loadLoginState(): LoginUser {
    if (!window) {
        return undefined;
    }

    const userJSON: string | null = window.localStorage.getItem("user");

    if (!userJSON || userJSON == null) {

        return undefined;
    }

    const userObj = JSON.parse(userJSON);
    return userObj;
}

export async function login(username: string, password: string, fetchFunc=fetch): Promise<LoginUser> {
    const loginStr = JSON.stringify({
        username: username,
        password: password,
    });

    const resp = await fetchFunc(`/api/realtime/login`, {
        method: "POST",
        body: loginStr,
    })

    if (!resp.ok) {
        return undefined;
    }

    const rawJWT = await resp.text();
    
    const payload = JSON.parse(atob(rawJWT.split(".")[1]));

    const tmp = {username: payload["username"], email: payload["email"], elo: payload["elo"], token: rawJWT, isAdmin: payload["isAdmin"]};
    loginState.set(tmp);
    window.dispatchEvent(new Event("userLoginChange"));
    return tmp;
}

export async function logout() {
    loginState.set(undefined);
    window.dispatchEvent(new Event("userLoginChange"));
}

