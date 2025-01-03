<script lang="ts">
    import Board from "$lib/components/Board.svelte";
    import { blur, fly } from "svelte/transition";
    import { page } from "$app/state";
    import { findGame, type ApiError, type Game } from "$lib/api.js";
    import Button from "$lib/components/Button.svelte";
    import Inforow from "$lib/components/Inforow.svelte";
    import { goto } from "$app/navigation";
    import { formatDate, translateDifficulty, translateGameState } from "$lib/format";
    import Overlay from "$lib/components/Overlay.svelte";
    import { circOut, sineInOut } from "svelte/easing";

    const { data }: { data: { slug: string } } = $props();

    let game: Game | undefined = $state(undefined);
    let mainClientWidth: number = $state(0);
    let gameInfoButtonVisible: boolean = $derived(mainClientWidth <= 1022);
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
<div class="stats">
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
    <main bind:clientWidth={mainClientWidth}>
        <Overlay bind:visible={gameInfoOverlayVisible}>
            <div class="info-wrapper" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
                {@render stats(game)}
            </div>
        </Overlay>
        <div class="game">
            <div class="board-wrapper">
                <div class:afterwin={winScreenShown} transition:blur>
                    <Board {onwin} bind:value={game} />
                    {#if winScreenShown}
                        <Button href="/gamelist/" variant={winSymbol == "X" ? "red" : "blue"} scaleDown={true}>
                            <div class="font-15pt" style="padding: 12px 40px 12px 20px; display: flex; gap: 12px">
                                <span class="back-arrow"></span>
                                Zpět na seznam úloh
                            </div>
                        </Button>
                    {/if}
                </div>
            </div>
            <div class="sidebar">
                {#if !gameInfoButtonVisible}
                    {@render stats(game)}
                {/if}
                <div class="controls">  
                    {#if gameInfoButtonVisible}
                        <Button variant="blue" scaleDown={true} onclick={toggleInfoOverlay}><div class="font-15pt">Info</div></Button>
                    {:else}
                        <div></div>
                    {/if}
                    <Button href="./edit" scaleDown={true}><div class="font-15pt">Upravit</div></Button>
                </div>
            </div>
        </div>
    </main>
{:else}
    Loading
{/if}

<style lang="scss">
    main {
        position: relative;
    }

    .font-15pt {
        font-size: 15pt;
        padding: 20px;

        @media screen and (max-width: 1022px) {
            font-size: 13pt;
            padding: 12px;
        }
    }

    .game {
        display: grid;
        grid-template-columns: 0.8fr 0.3fr;
        gap: 20px;

        width: 100dvw;
        height: calc(100dvh - var(--header-height));
        --button-bar-height: 0px;
        --padding: 50px;
        padding: 50px;

        @media screen and (max-width: 1022px) {
            display: flex;
            flex-direction: column;

            padding: 20px;
            --padding: 20px;
            --button-bar-height: 68px;
        }
    }

    .board-wrapper {
        overflow: hidden;
        height: calc(100dvh - var(--header-height) - 2 * var(--padding) - var(--button-bar-height));

        padding-top: 0;
        transition: all 250ms ease-in-out;
        transition-delay: 250ms;

        & > div {
            height: calc(100dvh - var(--header-height) - 2 * var(--padding) - var(--button-bar-height));
            margin: auto;
        }

        .afterwin {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 175px;
        }
    }

    .info-wrapper {
        font-size: 16pt;
        max-width: 500px;
        background-color: white;
        box-shadow: 2px 2px 20px #666;
        border-radius: 5px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        padding: 20px;
    }

    .sidebar {
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        gap: 20px;
    }

    .stats {
        display: grid;
        grid-template-columns: 1fr;
        grid-auto-rows: 1fr;
        gap: 20px;
        width: 100%;
        height: min-content;
    }

    .controls {
        display: grid;
        grid-template-columns: 1fr 1fr;
        width: 100%;
        gap: 20px;
        justify-content: space-between;
    }

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

    .back-arrow {
        background: url("/material-icons/arrow-back/48x48-white.svg");
        background-size: 19px 19px;
        background-position: center;
        width: 22px;
        height: 22px;
        display: inline-block;
    }
</style>
