<script lang="ts">
    import { goto } from "$app/navigation";
    import { listGames, type ApiError, type Game } from "$lib/api";
    import { flip } from "svelte/animate";
    import { filterList, type Filters } from "./filter";
    import SearchFilterBar from "./SearchFilterBar.svelte";
    import { fade, fly, slide } from "svelte/transition";
    import { untrack } from "svelte";
    import { formatDate, shortenText, translateDifficulty, translateGameState } from "$lib/format";
    import Button from "$lib/components/Button.svelte";
    import Board from "$lib/components/Board.svelte";
    import Inforow from "$lib/components/Inforow.svelte";

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
        <center style="font-size:20pt; font-weight: 600; margin-top: 150px;">Načítání</center>
    {:else if filteredList}
        <div class="cardlist" bind:this={cardlist}>
            {#each filteredList as elem (elem.uuid)}
                <div class="gamecard" animate:flip={{duration: 550}} transition:fly={{duration: 250, x: -200}}>
                    <div class="board"><Board value={elem} controls={false}/></div>
                    <h2 class="name">{shortenText(elem.name, 30)}</h2>
                    <div class="difficulty">
                        <Inforow key="Obtížnost">{translateDifficulty(elem.difficulty)}</Inforow>
                    </div>
                    <div class="gameState">
                        <Inforow key="Stav hry">{translateGameState(elem.gameState)}</Inforow>
                    </div>
                    <div class="date">
                        <Inforow key="Úprava">{formatDate(elem.updatedAt)}</Inforow>
                    </div>
                    <div class="btn1"><Button scaleDown={true} href="/game/{elem.uuid}">
                        <div style="padding: 15px; font-weight: 500;">Hrát</div>
                    </Button></div>
                    <div class="btn2"><Button variant="blue" scaleDown={true} href="/game/{elem.uuid}/edit">
                        <div style="padding: 15px; font-weight: 500;">Upravit</div>
                    </Button></div>
                </div>
            {/each}
        </div>
    {/if}
</main>

<style lang="scss">
    .cardlist {
        padding: 50px;
        display: grid;
        grid-template-columns: repeat(2, minmax(0, 1fr));
        gap: 20px;

        @media screen and (max-width: 1260px) {
            grid-template-columns: minmax(0, 1fr);
        }
    }
    .gamecard {
        background-color: #F6F6F6;
        
        border-radius: 10px;
        padding: 20px;
        border: 2px solid;
        display: grid;
        gap: 10px;

        grid-template-areas: 
            "name name name"
            "board diff diff"
            "board state state"
            "board date date"
            "board btn1 btn2";

        @media screen and (max-width: 766px) {
            grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
            grid-template-areas: 
                "name name"
                "board board"
                "board board"
                "diff diff"
                "state state"
                "date date"
                "btn1 btn2";
        }
        
        .name {
            margin-top: 0;
            text-align: center;
            grid-area: name;
        }

        .board {
            cursor: default;
            pointer-events: none;
            grid-area: board;
            margin: 20px;
            
            height: 250px;
        }

        .difficulty {
            grid-area: diff;
        }

        .gameState {
            grid-area: state;
        }

        .date {
            grid-area: date;
        }

        .btn1 {
            grid-area: btn1;
        }

        .btn2 {
            grid-area: btn2;
        }

        &:nth-child(2n) {
            border-color: #E31837;
        }

        &:nth-child(2n+1) {
            border-color: #0070BB;
        }
    }
</style>
