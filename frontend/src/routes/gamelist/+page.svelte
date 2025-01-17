<script lang="ts">
    import { goto } from "$app/navigation";
    import { listGames, type ApiError, type Game } from "$lib/api";
    import { flip } from "svelte/animate";
    import { filterList, type Filters } from "./filter";
    import SearchFilterBar from "./SearchFilterBar.svelte";
    import { fade, fly, slide } from "svelte/transition";
    import { untrack } from "svelte";

    let list: Game[] | undefined = $state();
    let filter: Filters = $state({});
    let filteredList: Game[] | undefined = $state([]);
    let cardlist: HTMLDivElement | undefined = $state();

    async function loadGameData(gameList: Promise<[Game[], ApiError | undefined]>) {
        const [val, err] = await gameList;

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        list = val;
    }

    $effect(() => {
        const flist = filterList(list, filter);
        
        untrack(() => {
            filteredList = flist;
        });
    });

    $effect(() => {
        const gameData = listGames(fetch);
        loadGameData(gameData);
    });
</script>

<main>
    <SearchFilterBar bind:filters={filter}/>
    {#if !filteredList}
        loading
    {:else if filteredList}
        <div class="cardlist" bind:this={cardlist}>
            {#each filteredList as elem (elem.uuid)}
                <div class="gamecard" animate:flip={{duration: 550}} transition:fly={{duration: 250, x: -200}}>
                    <div>{elem.name}</div>
                    <div>{elem.createdAt}</div>
                    <div>{elem.difficulty}</div>
                    <div>{elem.gameState}</div>
                    <div>{elem.updatedAt}</div>
                </div>
            {/each}
        </div>
    {/if}
</main>

<style lang="scss">
    .cardlist {
        padding: 50px;
        display: flex;
        flex-direction: column;
        gap: 20px;
    }
    .gamecard {
        background-color: #F6F6F6;
        
        border-radius: 10px;
        padding: 20px;
        border: 2px dashed;
        min-width: 500px;
        flex: 1;

        &:nth-child(2n) {
            border-color: #E31837;
        }

        &:nth-child(2n+1) {
            border-color: #0070BB;
        }
    }
</style>
