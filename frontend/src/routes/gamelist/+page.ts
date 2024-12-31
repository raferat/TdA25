import { listGames } from "$lib/api";
import type { PageLoad } from "../game/[slug]/$types";

export const prerender = true;

export const load: PageLoad = ({ fetch }) => {
    const gameData = listGames(fetch);
    return {
        gameData: gameData,
    }
}