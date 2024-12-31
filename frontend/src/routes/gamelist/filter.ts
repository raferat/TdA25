import type { Game } from "$lib/api";

export interface Filters {
    
}

export const filterList = (list: Game[] | undefined, filter: Filters | undefined) => {
    if (list === undefined) return undefined;
    if (filter === undefined) return list;

    return list;
}