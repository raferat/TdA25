<script lang="ts">
    import { defaultGameBase } from "$lib/api";
    import Board from "$lib/Board.svelte";
    import Button from "$lib/Button.svelte";
    import { fade, fly, scale, slide } from "svelte/transition";

    let playing = $state(true);

    let boardAnimationController = $state(0);

    const onwin = () => {
        playing = false;
    };

    let base = $state(defaultGameBase());

    const newGame = () => {
        base = defaultGameBase();
        if (playing) {
            boardAnimationController = (boardAnimationController + 1) % 5;
        }
        playing = true;
    }

</script>

<main>
    <div></div>
    <div class="board-wrapper">
        {#key boardAnimationController}
        <div in:scale={{delay: 250}} out:fly={{y: 500, duration: 250}}>
            <Board {onwin} bind:value={base}/>
        </div>
        {/key}
    </div>
    <div class="controls">
        {#if playing}
            <div transition:slide={{axis: "x", duration: 250}}>
                <Button variant="blue">Uložit</Button>
            </div>
        {/if}
        <Button onclick={newGame}>
            <div transition:fade>
            {#if playing}
                Zahodit
            {:else}
                Nová hra
            {/if}
            </div>
        </Button>
    </div>
</main>

<style lang="scss">
    main {
        display: grid;
        grid-template-columns: 0.2fr 0.8fr 0.2fr;

        width: 100vw;
        height: calc(100vh - 130px);
        
        padding: 50px;
    }

    .board-wrapper {
        height: calc(100vh - 230px);


        & > div {
            height: calc(100vh - 230px);
            margin: auto;
        }
    }

    .controls {
        display: flex;
        flex-direction: column;
        justify-content: end;
        align-items: end;
        gap: 20px;
    }

    
</style>