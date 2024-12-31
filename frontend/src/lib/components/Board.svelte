<script lang="ts">
    import { defaultGameBase, type Game, type GameBase } from "$lib/api";
    import { expoOut, sineOut } from "svelte/easing";
    import { fade, fly } from "svelte/transition";
    import { calculateNextSymbol, isWinMove } from "../boardutil";
    

    let { 
        value = $bindable(defaultGameBase()),
        controls = true,
        onwin = () => {},
        onmove = () => {},
    }: { 
        value?: GameBase | Game, 
        controls?: boolean,
        onwin?: () => void,
        onmove?: () => void,
    } = $props();


    let nextSymbol: "X" | "O" = $derived(calculateNextSymbol(value.board));
    let isPlaying = $state(true);

    function click(x: number, y: number) {
        if (!controls) return;

        const isWin = isWinMove(value.board, x, y, nextSymbol);
        
        if ( isWin ) {
            isPlaying = false;
            onwin();
        } else {
            value.board[y][x] = nextSymbol;
            onmove();
        }

        
    }

    $effect(() => {
        if ( value != undefined )
            isPlaying = true;
    });
</script>


{#if isPlaying}
<div class="wrapper" in:fade={{delay: 300}} out:fade={{duration: 250}}>
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
    {#if controls}
    <div class="indicator">
        <h3>Na tahu: 
            <span class="moveicon"
                class:o={nextSymbol == "O"}
                class:x={nextSymbol == "X"}>
            </span>
        </h3>
    </div>
    {/if}
</div>
{:else}
<div class="winner-wrapper" 
    out:fly={{y: -50, duration: 250, easing: sineOut}} 
    in:fly={{y: -50, duration: 250, delay: 250, easing: sineOut}}
    >
    <div class="winner">
        {#if nextSymbol == "X"}
        Vyhrál:
        {:else}
        Vyhrálo:
        {/if}
        <span 
            out:fly={{y: -50, duration: 300, easing: sineOut}} 
            in:fly={{y: -50, duration: 300, delay: 250, easing: sineOut}}
            class="moveicon"
            class:o={nextSymbol == "O"}
            class:x={nextSymbol == "X"}>
        </span>
    </div>
</div>
{/if}



<style lang="scss">
    .winner-wrapper {
        display: flex;
        justify-content: center;
        align-items: center;
        .winner {
            margin: auto;
            font-size: 30pt;


            .moveicon {
                
                display: block;
                margin: auto;
                margin-top: 40px;
                width: 100px;
                height: 100px;

                background-repeat: no-repeat;
            }
        }
    }


    .indicator {
        text-align: center;

        h3 {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 12px;
            transition: all 250ms cubic-bezier(1, 0, 0, 1);

            span {
                width: 1.4rem;
                height: 1.4rem;
                transition: all 250ms ease-in-out;
            }
        }
        
    }

    .wrapper {
        max-height: 100%;
        max-width: 100%;
        height: 100%;
        width: 100%;
        .board {
            max-width: 100%;
            max-height: calc(100% - 2.6em);
            margin: auto;
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
                cursor: pointer;

                &:disabled {
                    cursor: default;
                }
            }
        }
    }

    .moveicon {
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
