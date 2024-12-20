<script lang="ts">
    import { createGame, defaultGameBase, type GameBase } from '$lib/api';
    import { expoOut } from 'svelte/easing';
    import { fly } from 'svelte/transition';
    let { value = $bindable(defaultGameBase()) } : {value?: GameBase} = $props();

    let nextSymbol = "X";

    function click(x: number, y: number) {
        value.board[y][x] = nextSymbol;
        nextSymbol = nextSymbol == "X" ? "O" : "X";
    }
 </script>

<div class="board">
{#each value.board as row, y}
    {#each row as elem, x}
        <button onclick={(evt) => click(x,y)} disabled={elem != ''} aria-label="place your symbol">
            {#if elem != ''}
                <div in:fly={{x: 150, y: 150, easing: expoOut}} class="moveicon" class:o={elem == 'O'} class:x={elem == 'X'}></div>
            {/if}
        </button>
    {/each}
{/each}
</div>



<button onclick={() => createGame(value)}>Send!</button>

<style lang="scss">
    .board {
        display: grid;
        grid-template-columns: repeat(15, 1fr);
        grid-template-rows: repeat(15, 1fr);
        max-height: 100vh;
        max-width: 100vw;
        aspect-ratio: 1;
        
        background-color: #1A1A1A;
        gap: 1px;
        button {
            padding: 12%;
            aspect-ratio: 1;
            border: none;
            border-radius: 0;
            background: white;
            color: black;
        }
    }

    div.moveicon {
        width: 100%;
        height: 100%;

        &.x {
            background-image: url("/icons/X/X_cervene.svg");
        }

        &.o {
            background-image: url("/icons/O/O_modre.svg");
        }
    }
</style>