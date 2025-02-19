import { findGame } from '$lib/api';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ params, fetch }) => {
    const gameData = findGame(params.slug, fetch);
    return {
        slug: params.slug,
        gameData: gameData,
    }
}