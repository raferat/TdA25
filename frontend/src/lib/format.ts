import type { Difficulty, GameState } from "./api";

export function formatDate(date: string | Date): string {
    if (typeof date === "string") {
        date = new Date(date);
    }

    const hh = `${date.getHours()}`.padStart(2, '0');
    const mm = `${date.getMinutes()}`.padStart(2, '0');
    const ss = `${date.getSeconds()}`.padStart(2, '0');

    const dd = `${date.getDate()}`;
    const MM = `${date.getMonth() + 1}`;
    const YY = `${date.getFullYear()}`;


    return `${dd}.${MM}.${YY}\xa0\xa0${hh}:${mm}:${ss}`;
}

export function translateGameStateUnsafe(state: GameState | undefined): string | undefined {
    if (state == undefined) {
        return undefined
    }
    return translateGameState(state);
}

export function translateGameState(gameState: GameState): string {
    switch (gameState) {
        case "opening":
            return "Zahájení";
        case "midgame":
            return "V průběhu";
        case "endgame":
            return "Koncovka";
        case "unknown":
            return "Neznámý stav";
    }
}
export function translateDifficultyUnsafe(difficulty: Difficulty | undefined): string | undefined {
    if (difficulty == undefined) {
        return undefined
    }
    return translateDifficulty(difficulty);
}
export function translateDifficulty(difficulty: Difficulty): string {
    switch (difficulty) {
        case "beginner":
            return "Začátečnická";
        case "easy":
            return "Jednoduchá";
        case "medium":
            return "Pokročilá";
        case "hard":
            return "Těžká";
        case "extreme":
            return "Nejtěžší";
    }
}

export function translateTimeFrameUnsafe(updatedAgo: number | undefined): string | undefined {
    if (updatedAgo == undefined)
        return undefined;
    
    switch (updatedAgo) {
        case 1000*60*60*24:
            return "Za den";
        case 1000*60*60*24*7:
            return "Za týden";
        case 1000*60*60*24*31:
            return "Za měsíc";
        case 1000*60*60*24*31*3:
            return "Za 3 Měsíce";
    }
}

export function shortenText (a: string, length: number): string {
    if (a.length <= length) return a;
    const len = length;
    let result = a.slice(0, len/2);
    result += "...";
    result += a.slice(a.length-len/2);

    if (result.length >= a.length) return a;
    return result;
}