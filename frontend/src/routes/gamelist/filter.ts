import type { Difficulty, Game, GameState } from "$lib/api";

import '$lib/utils';

export interface Filters {
    nameSearch?: string,
    difficulty?: Difficulty,
    gameState?: Array<GameState>,
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
        if (filter.difficulty) {
            if (filter.difficulty != value.difficulty)
                return false;
        }
        return true;
    });

    return filtered;
}