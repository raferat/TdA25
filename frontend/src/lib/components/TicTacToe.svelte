<script lang="ts">
    import type { Board } from "$lib/api";
    import { calculateNextSymbol } from "$lib/boardutil";
    let {
        board = $bindable(),
        playing = true,
    }: { 
        board: Board,
        playing?: boolean,
    } = $props();


    let boardWidth = $state(0);
    let boardHeight = $state(0);
    let boardSize = $derived(Math.min(boardWidth, boardHeight) - 20);

    let nextSymbol = $derived(calculateNextSymbol(board));

    const placeSymbolAt = (x: number, y:number) => {
        if (!playing) {
            return null;
        }
        return () => {
            if (board[y][x] == "" ) {
                board[y][x] = nextSymbol;
                board = board;
            }
        }
    }
</script>

<div class="size-full flex justify-center items-center" bind:clientHeight={boardHeight} bind:clientWidth={boardWidth}>
    <div class="bg-tcream dark:bg-tlgrey p-5 md:p-10 lg:p-16 rounded-xl relative" style="width: {boardSize}px; height: {boardSize}px">
        <div class="grid grid-cols-15 aspect-square m-auto">
            {#each board as row, y}
                {#each row as cell, x}
                    <button 
                        aria-label="Place next symbol on position [{x}, {y}]"
                        class="icon aspect-square border border-solid border-tlgrey dark:border-tdark-grid bg-no-repeat bg-center cursor-pointer"
                        class:no-top-border={y == 0}
                        class:no-left-border={x == 0}
                        class:no-bottom-border={y == 14}
                        class:no-right-border={x == 14}
                        class:x={cell == "X"}
                        class:o={cell == "O"}
                        onclick={placeSymbolAt(x,y)}></button>
                {/each}
            {/each}
        </div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute saturate-0 lg:top-8 lg:left-8 md:top-4 md:left-4 top-2 left-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute saturate-0 lg:top-8 lg:right-8 md:top-4 md:right-4 top-2 right-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute saturate-0 lg:bottom-8 lg:right-8 md:bottom-4 md:right-4 bottom-2 right-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute saturate-0 lg:bottom-8 lg:left-8 md:bottom-4 md:left-4 bottom-2 left-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
    </div>
</div>


<style>
    .no-top-border {
        border-top: none;
    }

    .no-left-border {
        border-left: none;
    }

    .no-right-border {
        border-right: none;
    }

    .no-bottom-border {
        border-bottom: none;
    }

    .icon {
        background-image: url("/icons/transparent.svg");
        background-size: 0%;
        transition: background-size 1s cubic-bezier(0.19, 1, 0.22, 1);
        transition: background-image 1s cubic-bezier(0.19, 1, 0.22, 1);
    }

    .x {
        background-image: url("/icons/X/X_cervene.svg");
        background-size: 75%;
    }

    .o {
        background-image: url("/icons/O/O_modre.svg");
        background-size: 75%;
    }
</style>