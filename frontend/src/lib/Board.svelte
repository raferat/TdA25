<script lang="ts">
    import { createGame, defaultGameBase, type GameBase } from '$lib/api';
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
        <button onclick={(evt) => click(x,y)} disabled={elem != ''}>{elem}</button>
    {/each}
{/each}
</div>
{value.board}
<button onclick={() => createGame(value)}>Send!</button>

<style lang="scss">
    .board {
        display: grid;
        grid-template-columns: repeat(15, 1fr);
        grid-template-rows: repeat(15, 1fr);
        max-width: 1000px;
        background-color: black;
        gap: 2px;
        button {
            padding: 5px;
            aspect-ratio: 1;
            border: none;
            border-radius: 0;
            background: white;
            color: black;
        }
    }
</style>