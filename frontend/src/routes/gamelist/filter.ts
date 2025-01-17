import type { Difficulty, Game, GameState } from "$lib/api";

import '$lib/utils';

export interface Filters {
    nameSearch?: string,
    difficulty?: Difficulty,
    gameState?: GameState,
    updatedAgo?: number,
    sortProperty?: "name" | "created" | "updated" | "difficulty" | "gameState",
    sortDesc?: boolean,
}

export const filterList = (list: Game[] | undefined, filter: Filters | undefined) => {
    if (list === undefined) return undefined;
    if (filter === undefined) return list;

    let filtered = list.filter((value: Game, idx: number) => {
        if (filter.nameSearch && filter.nameSearch.length > 0) {
            if (!value.name.normalized().includes(filter.nameSearch.normalized()))
                return false;
        }
        if (filter.difficulty && filter.difficulty != value.difficulty) {
            return false;
        }
        if (filter.gameState && filter.gameState != value.gameState) {
            return false
        }
        if (filter.updatedAgo && Date.now() - filter.updatedAgo > new Date(value.updatedAt).getTime()) {
            return false
        }
        return true;
    });

    return filtered;
}