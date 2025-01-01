<script lang="ts">
    import { goto } from "$app/navigation";
    import { listGames, type ApiError, type Game } from "$lib/api";
    import { filterList, type Filters } from "./filter";

    let list: Game[] | undefined = $state();
    let filter: Filters = $state({});
    let filteredList: Game[] | undefined = $derived(filterList(list, filter));


    async function loadGameData(gameList: Promise<[Game[], ApiError | undefined]>) {
        const [val, err] = await gameList;

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        list = val;
    }

    $effect(()=>{
        const gameData = listGames(fetch);
        loadGameData(gameData);
    })
</script>

<main>
{#if filteredList}
    {filteredList}
{:else}
    Loading...    
{/if}
</main>