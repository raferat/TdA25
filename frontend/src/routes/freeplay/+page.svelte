<script lang="ts">
    import TicTacToe from "$lib/components/TicTacToe.svelte";
    import {calculateNextSymbol} from "$lib/boardutil";
    import Background from "$lib/components/Background.svelte";
    

    let gameState: "pregame" | "waiting" | "started" = $state("pregame");
    let selectedSymbol: "X" | "O" = $state("O");

    let board: ("X"|"O"|"")[][] = $state(Array.from({ length: 15 }, () => Array(15).fill("")));

    let currentTurnSymbol = $derived(calculateNextSymbol(board));

    let gameCode: string = $state("");

    let ws: WebSocket;

    function startGame() {
        try {
            ws = new WebSocket(`wss://${window.location.host}/realtime/freeplay`);
        } catch (e) {
            ws = new WebSocket(`ws://${window.location.host}/realtime/freeplay`);
        }
        ws.onopen = () => {
            ws.send("create " + selectedSymbol);
        }

        ws.onmessage = (ev) => {
            let msg = ev.data as string;

            if (msg.startsWith("code ")) {
                gameCode = msg.split(" ")[1];
                gameState = "waiting";
            } else if (msg.startsWith("start ")) {
                gameState = "started";
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

    function symbolPlaced(x: number, y: number, symbol: "X" | "O") {
        ws.send(`${x} ${y}`);
    }
</script>

<Background height={0} width={0}/>
<main class="size-full">
{#if gameState == "pregame"}
    <div class="size-full flex justify-center items-center">
        <div class="p-6 border-2 rounded-xl mb-10">
            <h2 class="text-2xl font-bold mb-3">Založit hru</h2>
            Vyberte symbol:
            <div class="grid grid-cols-2 auto-rows-fr gap-4 min-w-40 mt-2">
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
                <img onclick={() => selectedSymbol = "X"} style="border-color: {selectedSymbol == 'X' ? 'black' : 'transparent'}" class="cursor-pointer border-2 border-transparent rounded-xl p-2" src="/icons/X/X_cervene.svg" alt="krizek">
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
                <img onclick={() => selectedSymbol = "O"} style="border-color: {selectedSymbol == 'O' ? 'black' : 'transparent'}" class="cursor-pointer border-2 border-transparent rounded-xl p-2" src="/icons/O/O_modre.svg" alt="kolecko">
            </div>
            <button onclick={startGame} class="pbred mt-4">Spustit</button>
        </div>
    </div>
{:else if gameState == "waiting"}
    <div>
        <h2 class="text-xl font-bold text-center mt-8">Čekání na druhého hráče</h2>
        <div class="text-center">Kód hry je:</div>
        <div class="text-3xl text-center">{gameCode}</div>
    </div>
{:else if gameState == "started"}
    <div class="size-full" class:overlay={currentTurnSymbol != selectedSymbol}>
        <TicTacToe bind:board={board} placedSymbol={symbolPlaced} />
    </div>
{/if}
</main>

<style>
    .overlay {
        pointer-events: none;
        cursor: default;
    }
</style>