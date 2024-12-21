<script lang="ts">
    import { createGame, defaultGameBase, type GameBase } from "$lib/api";
    import { expoOut } from "svelte/easing";
    import { fade } from "svelte/transition";
    

    let { value = $bindable(defaultGameBase()) }: { value?: GameBase } =
        $props();

    let nextSymbol = "X";

    function click(x: number, y: number) {
        value.board[y][x] = nextSymbol;
        nextSymbol = nextSymbol == "X" ? "O" : "X";
    }
</script>

<div class="wrapper">
    <div class="board">
        {#each value.board as row, y}
            {#each row as elem, x}
                <button
                    onclick={(evt) => click(x, y)}
                    disabled={elem != ""}
                    aria-label="place your symbol">
                    {#if elem != ""}
                        <div
                            in:fade={{ easing: expoOut, duration: 500 }}
                            class="moveicon"
                            class:o={elem == "O"}
                            class:x={elem == "X"}
                        ></div>
                    {/if}
                </button>
            {/each}
        {/each}
    </div>
</div>



<style lang="scss">
    .wrapper {
        
        aspect-ratio: 1;

        .board {
            display: grid;
            grid-template-columns: repeat(15, 1fr);
            grid-template-rows: repeat(15, 1fr);
            aspect-ratio: 1;
            
            background-color: #1a1a1a;

            gap: 2px;
            button {
                padding: 12%;
                aspect-ratio: 1;
                border: none;
                border-radius: 0;
                background: #F6F6F6;
                color: black;
            }
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
