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

export function translateDifficulty(difficulty: Difficulty): string {
    switch (difficulty) {
        case "beginner":
            return "Začátečnická";
        case "easy":
            return "Jednoduchá";
        case "medium":
            return "Středně těžká";
        case "hard":
            return "Těžká";
        case "extreme":
            return "Extrémní";
    }
}