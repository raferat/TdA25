<script lang="ts">
    import { fly, slide } from "svelte/transition";
    import { page } from "$app/state";
    import { findGame, type ApiError, type Game } from "$lib/api";
    import Inforow from "$lib/components/Inforow.svelte";
    import { goto } from "$app/navigation";
    import { formatDate, translateDifficulty, translateGameState } from "$lib/format";
    import Overlay from "$lib/components/Overlay.svelte";
    import { circOut, sineInOut } from "svelte/easing";
    import TicTacToe from "$lib/components/TicTacToe.svelte";

    const { data }: { data: { slug: string } } = $props();

    let game: Game | undefined = $state(undefined);
    let mainClientWidth: number = $state(0);
    let gameInfoButtonVisible: boolean = $derived(mainClientWidth <= 1024);
    let gameInfoOverlayVisible: boolean = $state(false);
    let winSymbol: "X" | "O" = $state("X");

    let winScreenShown: boolean = $state(false);

    async function processData(gameData: Promise<[Game, ApiError | undefined]>) {
        const [val, err] = await gameData;
        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }
        game = val;
    }

    function copyText(text: string): () => void {
        return () => {
            navigator.clipboard.writeText(text);
        };
    }

    function toggleInfoOverlay() {
        gameInfoOverlayVisible = !gameInfoOverlayVisible;
    }

    function onwin(winner: "X" | "O") {
        winSymbol = winner;
        winScreenShown = true;
    }

    $effect(() => {
        const gameData = findGame(data.slug);
        processData(gameData);
    })

    const permanentURL = page.url.host + page.url.pathname;
</script>


{#snippet stats(game: Game)}
<div class="grid grid-cols-1 auto-rows-fr gap-2">
    <Inforow key="Jméno hry" text={game.name} />
    <Inforow key="Obtížnost" text={translateDifficulty(game.difficulty)} />
    <Inforow key="Stav hry" text={translateGameState(game.gameState)} />
    <Inforow key="Úloha dne" text={formatDate(game.updatedAt)} />
    <Inforow key="Permanentní odkaz" text={permanentURL} minify={true}>
        {#snippet renderer(text)}
            <a href="{page.url.protocol}//{permanentURL}">{text}</a>
            <!-- svelte-ignore a11y_no_static_element_interactions -->
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <span onclick={copyText(`${page.url.protocol}//${permanentURL}`)} class="copy-icon"></span>
        {/snippet}
    </Inforow>
</div>
{/snippet}


{#if game}
    <main class="relative size-full grid lg:grid-cols-[minmax(0,0.7fr)_minmax(0,0.3fr)] p-2 grid-cols-1" bind:clientWidth={mainClientWidth}>
        <Overlay bind:visible={gameInfoOverlayVisible}>
            <div class="p-4 bg-twhite dark:bg-tlgrey rounded-xl flex flex-col gap-3" 
                in:fly={{ y: -500, easing: circOut }} 
                out:fly={{ y: -500, easing: sineInOut }}>
                {@render stats(game)}
            </div>
        </Overlay>
        
        <TicTacToe bind:board={game.board} {onwin}/>
        
        <div class="flex flex-col gap-1 lg:justify-between justify-end">
            {#if !gameInfoButtonVisible}
                <div class="bg-tlgrey rounded-xl p-4 mt-3">
                {@render stats(game)}
                </div>
            {/if}
            <div class="grid grid-cols-2 gap-2 w-full">  
                {#if gameInfoButtonVisible && !winScreenShown }
                    <button class="pbblue text-xl" onclick={toggleInfoOverlay}>Info</button>
                {:else}
                    <div></div>
                {/if}
                {#if winScreenShown}
                    <button class="pbred text-xl" onclick={() => goto("/gamelist/")}>Zpět na seznam úloh</button>
                {:else}
                    <button class="pbred text-xl" onclick={() => goto("./edit")}>Upravit</button>
                {/if}
            </div>
        </div>
    </main>
{:else}
<center class="text-3xl font-bold pt-20">Načítání</center>
{/if}

<style lang="scss">
    .copy-icon {
        position: absolute;
        right: 10px;
        top: 10px;
        background: url("/material-icons/copy/cervene2.svg");
        background-repeat: no-repeat;
        background-size: 1.4rem 1.4rem;
        background-position: 50% 50%;

        
        display: inline-block;
        margin-left: auto;
        width: 1.4rem;
        height: 1.4rem;

        cursor: pointer;
        border-radius: 100px;
        padding: 20px;
        transition: all 250ms cubic-bezier(1, 0, 0, 1);
        transition-delay: 200ms;

        &:active {
            transform: scale(1.2);
            transition-delay: 10ms;
            transition-duration: 20ms;
            background-color: #00000020;
            
        }
    }
</style>
