<script lang="ts">
    import { goto } from "$app/navigation";
    import { listGames, type ApiError, type Game } from "$lib/api";
    import { flip } from "svelte/animate";
    import { filterList, type Filters } from "./filter";
    import SearchFilterBar from "./SearchFilterBar.svelte";
    import { fade, fly } from "svelte/transition";
    import { untrack } from "svelte";
    import { formatDate, shortenText, translateDifficulty, translateGameState } from "$lib/format";
    import Inforow from "$lib/components/Inforow.svelte";
    import TicTacToe from "$lib/components/TicTacToe.svelte";

    let list: Game[] | undefined = $state();
    let filter: Filters = $state({});
    let filteredList: Game[] | undefined = $state([]);
    let cardlist: HTMLDivElement | undefined = $state();
    let noGames: boolean = $state(false);

    async function loadGameData(gameList: Promise<[Game[], ApiError | undefined]>) {
        try {
        const [val, err] = await gameList;

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        list = val;
        } catch (e) {
        if (list == undefined)
            noGames = true;
        else
            noGames = false;
        }
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
    {#if noGames}
        <center class="font-semibold text-2xl mt-50">Zatím nejsou zveřejněny žádné úlohy</center>
    {:else if !filteredList}
        <center class="font-semibold text-2xl mt-50">Načítání</center>
    {:else if filteredList}
        <div class="p-8 md:p-16 grid grid-cols-1 lg:grid-cols-2 gap-2" bind:this={cardlist} in:fade>
            {#each filteredList as elem (elem.uuid)}
                <div class="bg-tcream dark:bg-tlgrey p-8 rounded-xl border-2 grid grid-cols-1 gap-2" animate:flip={{duration: 550}} transition:fly={{duration: 250, x: -200}}>
                    <h2 class="text-xl font-bold text-center">{shortenText(elem.name, 30)}</h2>
                    <div class="cursor-default pointer-events-none w-full h-[450px]"><TicTacToe board={elem.board} playing={false}/></div>
                    <Inforow key="Obtížnost">{translateDifficulty(elem.difficulty)}</Inforow>
                    <Inforow key="Stav hry">{translateGameState(elem.gameState)}</Inforow>
                    <Inforow key="Úprava">{formatDate(elem.updatedAt)}</Inforow>
                    <button class="pbred text-xl" onclick={() => {goto(`/game/${elem.uuid}`);}}>Hrát</button>
                    <button class="pbblue text-xl" onclick={() => {goto(`/game/${elem.uuid}/edit`);}}>Upravit</button>
                </div>
            {/each}
        </div>
    {/if}
</main>