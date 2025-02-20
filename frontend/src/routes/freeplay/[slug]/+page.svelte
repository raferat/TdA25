<script lang="ts">
    import { calculateNextSymbol } from '$lib/boardutil';
    import Background from '$lib/components/Background.svelte';
    import Overlay from '$lib/components/Overlay.svelte';
    import TicTacToe from '$lib/components/TicTacToe.svelte';
    import { untrack } from 'svelte';


    let { data } = $props();

    let ws: WebSocket;
    
    let selectedSymbol: "X" | "O" | undefined = $state(undefined);

    let board: ("X"|"O"|"")[][] = $state(Array.from({ length: 15 }, () => Array(15).fill("")));

    let currentTurnSymbol = $derived(calculateNextSymbol(board));

    let badConnection: boolean = $state(false);

    function connectToGame() {
        if (window.location.host.endsWith(".app")) {
            ws = new WebSocket(`wss://${window.location.host}/realtime/freeplay`);
        } else {
            ws = new WebSocket(`ws://${window.location.host}/realtime/freeplay`);
        }
        
        ws.onopen = () => {
            ws.send("connect " + data.slug);
        }

        ws.onmessage = (ev) => {
            let msg = ev.data as string;

            if (msg.startsWith("start ")) {
                selectedSymbol = msg.split(" ")[1] as ("X" | "O");
            } else if (msg.startsWith("placeX ")) {
                const cmd = msg.split(" ");
                let x = parseInt(cmd[1]);
                let y = parseInt(cmd[2]);
                board[y][x] = "X";
            } else if (msg.startsWith("placeO ")) {
                const cmd = msg.split(" ");
                let x = parseInt(cmd[1]);
                let y = parseInt(cmd[2]);
                board[y][x] = "O";
            }
        };
    }

    $effect(() => {
        untrack(() => {
            if (data.slug.length != 6) return;
            connectToGame();
            setTimeout(() => {
                if (selectedSymbol == undefined)
                    badConnection = true;
            }, 5000);
        })
    });

    let overlayVisible = $state(true);

    $effect(() => {
        if (selectedSymbol != undefined) {
            untrack(() => {
                setTimeout(() => {
                    overlayVisible = false;
                }, 5000);
            });
        }
    });

    function symbolPlaced(x: number, y: number, symbol: "X" | "O") {
        ws.send(`${x} ${y}`);
    }
</script>

<Background height={0} width={0}/>
{#if data.slug.length != 6 || badConnection}
<h2 class="text-center text-3xl">Neplatný kód</h2>
{:else if selectedSymbol}
<main class="size-full" class:overlay={currentTurnSymbol != selectedSymbol}>
    <Overlay bind:visible={overlayVisible}>
        <div class="bg-tcream p-6 rounded-xl">
            <h2 class="text-xl font-semibold mb-2">Hrajete jako</h2>
            {#if selectedSymbol == "X"}
                <img class="max-w-15 m-auto" src="/icons/X/X_cervene.svg" alt="">
            {:else}
                <img class="max-w-15 m-auto" src="/icons/O/O_modre.svg" alt="">
            {/if}
        </div>
    </Overlay>
    <TicTacToe bind:board={board} placedSymbol={symbolPlaced} />
</main>
{/if}

<style>
    .overlay {
        pointer-events: none;
        cursor: default;
    }
</style>