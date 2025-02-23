<script lang="ts">
    import { goto } from '$app/navigation';
    import { initLoginState, loginState, refreshlogin } from '$lib/api';
    import { calculateNextSymbol } from '$lib/boardutil';
    import Background from '$lib/components/Background.svelte';
    import Overlay from '$lib/components/Overlay.svelte';
    import TicTacToe from '$lib/components/TicTacToe.svelte';
    import { untrack } from 'svelte';
    import { get } from 'svelte/store';


    let ws: WebSocket;
    
    let selectedSymbol: "X" | "O" | undefined = $state(undefined);

    let board: ("X"|"O"|"")[][] = $state(Array.from({ length: 15 }, () => Array(15).fill("")));

    let currentTurnSymbol = $derived(calculateNextSymbol(board));

    let newElo = $state();

    function connectToGame() {
        if (window.location.host.endsWith(".app")) {
            ws = new WebSocket(`wss://${window.location.host}/realtime/joinmatch`);
        } else {
            ws = new WebSocket(`ws://${window.location.host}/realtime/joinmatch`);
        }
        
        ws.onopen = () => {
            ws.send("auth " + get(loginState)?.token);
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
            } else if (msg.startsWith("Xwon")) {
                newElo = msg.split(" ")[1];
                winSymbol = "X";
                setTimeout(() => {
                    finished = true;
                }, 1000);
                refreshlogin();
                
            } else if (msg.startsWith("Owon")) {
                newElo = msg.split(" ")[1];
                winSymbol = "O";
                setTimeout(() => {
                    finished = true;
                }, 1000);
                refreshlogin();
            }
        };
    }

    $effect(() => {
        initLoginState();
        untrack(() => {
            if (get(loginState) == undefined) {
                goto("/play/");
                return;
            }

            connectToGame();
        });
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

    let finished = $state(false);
    let winSymbol = $state();
</script>

<Background height={0} width={0}/>
{#if finished}
<center class="w-50 m-auto p-6">
    <div class="text-2xl mb-2 font-bold">Vyhrál symbol</div>
    {#if winSymbol == "X"}
        <img class="m-auto w-20" src="/icons/X/X_cervene.svg" alt="">
    {:else if winSymbol == "O"}
        <img class="m-auto w-20" src="/icons/O/O_modre.svg" alt="">
    {:else}
        <div class="text-xl">Byla to remíza</div>
    {/if}
    <div class="text-xl">Vaše nové ELO je:</div>
    <div class="text-3xl font-semibold">{newElo}</div>
    <button onclick={_ => goto("/play/")} class="pbred w-full mt-2">Zpět</button>
</center>
{:else if selectedSymbol == undefined}
<center class="p-6"><h2 class="text-3xl">Čekání na oponenta</h2></center>
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