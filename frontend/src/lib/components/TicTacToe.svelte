<script lang="ts">
    import type { Board } from "$lib/api";
    import { calculateNextSymbol, getWinBounds, isWinMove } from "$lib/boardutil";
    import { sineInOut } from "svelte/easing";
    import { draw } from "svelte/transition";
    let {
        board = $bindable(),
        playing = true,
        onwin = (symbol: "X" | "O") => {},
        placedSymbol = () => {},
    }: { 
        board: Board,
        playing?: boolean,
        onwin?: (symbol: "X" | "O") => void;
        placedSymbol?: (x:number, y:number, symbol: "X" | "O") => void,
    } = $props();


    let boardWidth = $state(0);
    let boardHeight = $state(0);
    let boardSize = $derived(Math.min(boardWidth, boardHeight) - 20);

    let nextSymbol = $derived(calculateNextSymbol(board));

    const placeSymbolAt = (x: number, y:number) => {
        if (!playing) {
            return () => {};
        }
        return () => {
            if (board[y][x] != "" ) return;
            placedSymbol(x, y, nextSymbol);

            if (isWinMove(board, x, y, nextSymbol)) {
                playing = false;
                const bounds = getWinBounds(board, x, y, nextSymbol);
                drawWinLine(bounds.x1, bounds.y1, bounds.x2, bounds.y2);

                setTimeout(() => {
                    const a = nextSymbol;
                    onwin(a); 
                }, 500);
            }

            board[y][x] = nextSymbol;
            board = board;
            
        }
    }

    let winSvgLineData: {x1: number, y1: number, x2: number, y2: number, color: string} = $state({x1: 0, x2: 0, y1: 0, y2: 0, color: "red"});
    let gameFinished: boolean = $state(false);

    function drawWinLine(x1: number, y1: number, x2: number, y2: number) {
        const startElem = document.getElementById(`btn-${x1}x${y1}`);
        const endElem = document.getElementById(`btn-${x2}x${y2}`);

        if (!startElem) return;
        if (!endElem) return;

        const startBox = startElem.getBoundingClientRect();
        const endBox = endElem.getBoundingClientRect();
        const parentBox = startElem.parentElement?.getBoundingClientRect();

        if (!parentBox) return;

        const sX = startBox.x + startBox.width/2 - parentBox.x;
        const sY = startBox.y + startBox.height/2 - parentBox.y;

        const eX = endBox.x + endBox.width/2 - parentBox.x;
        const eY = endBox.y + endBox.height/2 - parentBox.y;

        winSvgLineData.x1 = sX;
        winSvgLineData.y1 = sY;
        winSvgLineData.x2 = eX;
        winSvgLineData.y2 = eY;

        if (board[y1][x1] == "O") {
            winSvgLineData.color = "var(--color-tblue);"
        } else {
            winSvgLineData.color = "var(--color-tred);"
        }

        gameFinished = true;
    }
</script>

<div class="size-full flex justify-center items-center" bind:clientHeight={boardHeight} bind:clientWidth={boardWidth}>
    <div class="bg-tcream dark:bg-tlgrey p-5 lg:p-10 rounded-xl relative" style="width: {boardSize}px; height: {boardSize}px">
        <div class="grid grid-cols-15 aspect-square m-auto relative">
            {#each board as row, y}
                {#each row as cell, x}
                    <button
                        id="btn-{x}x{y}" 
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
            <div class="absolute top-0 left-0 size-full pointer-events-none">
                <svg width="100%" height="100%" xmlns="http://www.w3.org/2000/svg">
                    {#if gameFinished}
                    <line 
                        x1={winSvgLineData.x1} 
                        y1={winSvgLineData.y1} 
                        x2={winSvgLineData.x2} 
                        y2={winSvgLineData.y2} 
                        style="stroke: {winSvgLineData.color}; border: 1px solid pink;"
                        stroke-width="6"
                        stroke-linecap="round" 
                        in:draw={{duration: 500, easing: sineInOut}}/>
                    {/if}
                </svg>
            </div>
        </div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute lg:top-4 lg:left-4 top-2 left-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute lg:top-4 lg:right-4 top-2 right-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute lg:bottom-4 lg:right-4 bottom-2 right-2"
            class:x={nextSymbol == "X"} 
            class:o={nextSymbol == "O"}
            style="width: {boardSize/30}px"></div>
        <div 
            class="icon aspect-square bg-no-repeat bg-center absolute lg:bottom-4 lg:left-4 bottom-2 left-2"
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