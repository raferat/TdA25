<script lang="ts">
    import Board from "$lib/components/Board.svelte";
    import { blur, fly } from "svelte/transition";
    import { deleteGame, updateGame, type ApiError, type Game } from "$lib/api.js";
    import "$lib/utils";
    import Button from "$lib/components/Button.svelte";
    import Inforow from "$lib/components/Inforow.svelte";
    import { goto } from "$app/navigation";
    import { formatDate, translateGameState } from "$lib/format";
    import Overlay from "$lib/components/Overlay.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import RadioButtons from "$lib/components/RadioButtons.svelte";
    import { circOut, sineInOut } from "svelte/easing";
    import { isBoardCorrect } from "$lib/boardutil";
    import Background from "$lib/components/Background.svelte";

    const { data }: { data: { gameData: Promise<[Game, ApiError | undefined]>, slug: string } } =
        $props();

    let game: Game | undefined = $state(undefined);
    let deleteGameOverlayVisible: boolean = $state(false);
    let mainClientWidth: number = $state(0);
    let gameInfoButtonVisible: boolean = $derived(mainClientWidth <= 1062);
    let gameInfoOverlayVisible: boolean = $state(false);
    let loading: boolean = $state(false);

    async function processData(gameData: Promise<[Game, ApiError | undefined]>) {
        const [val, err] = await gameData;
        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }
        game = val;
    }

    async function saveData() {
        loading = true;
        await Promise.all([
            Promise.after(5000),
            saveGameData(),
        ]);
        loading = false;
    }

    async function saveGameData() {
        if (!game) return;

        
        const [val, err] = await updateGame({
            board: game.board,
            difficulty: game.difficulty,
            name: game.name,
        }, game.uuid);

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        game.gameState = val.gameState;
        game.name = val.name;
        game.difficulty = val.difficulty;
    }

    function back() {
        goto("../");
    }

    function toggleDeleteOverlay() {
        deleteGameOverlayVisible = !deleteGameOverlayVisible;
    }

    function toggleInfoOverlay() {
        gameInfoOverlayVisible = !gameInfoOverlayVisible;
    }

    async function confirmGameDelete() {
        if (!game) return;
        await deleteGame(game.uuid);
        goto("/gamelist/");
    }

    $effect(() => {
        processData(data.gameData);
    })
</script>

<Background width={0} height={0}/>

{#snippet stats(game: Game)}
<div class="stats">
    <TextInput placeholder="Jméno hry:" bind:value={game.name}/>
    <div>
        Obtížnost:
        <RadioButtons
            bind:value={game.difficulty}
            options={[
                { display: "Začátečnická", option: "beginner" },
                { display: "Jednoduchá", option: "easy" },
                { display: "Středně těžká", option: "medium" },
                { display: "Těžká", option: "hard" },
                { display: "Extrémní", option: "extreme" },
            ]}/>
    </div>
    <br>
    <br>
    <Inforow key="Stav hry" text={translateGameState(game.gameState)} />
    <Inforow key="Vytvořeno" text={formatDate(game.createdAt)} />
    <Inforow key="Úloha dne" text={formatDate(game.updatedAt)} />
</div>
{/snippet}

{#if loading}
<div>Loading</div>

{:else if game}
    <main bind:clientWidth={mainClientWidth}>
        <Overlay bind:visible={deleteGameOverlayVisible}>
            <div class="edit-wrapper" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
                <h2>Potvrďte smazání.</h2>
                <div class="save-options">
                    <Button scaleDown={true} onclick={toggleDeleteOverlay}>
                        <div class="font-14pt">Zrušit</div>
                    </Button>
                    <Button scaleDown={true} variant="blue" onclick={confirmGameDelete}>
                        <div class="font-14pt">Smazat</div>
                    </Button>
                </div>
            </div>
        </Overlay>
        <Overlay bind:visible={gameInfoOverlayVisible}>
            <div class="edit-wrapper" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
                {@render stats(game)}
                <Button disabled={!isBoardCorrect(game.board)} variant="blue" scaleDown={true} onclick={saveData}><div class="font-15pt">Uložit</div></Button>
            </div>
        </Overlay>


        <div class="game">
            <div class="board-wrapper">
                <div transition:blur>
                    <Board editMode={true} bind:value={game} />
                </div>
            </div>

            <div class="sidebar">
                {#if !gameInfoButtonVisible}
                {@render stats(game)}
                {/if}
                <div class="controls">
                    <Button scaleDown={true} onclick={toggleDeleteOverlay}><div class="font-15pt">Smazat</div></Button>
                    {#if gameInfoButtonVisible}
                        <Button variant="blue" scaleDown={true} onclick={toggleInfoOverlay}><div class="font-15pt">Info</div></Button>
                    {:else}
                        <Button disabled={!isBoardCorrect(game.board)} variant="blue" scaleDown={true} onclick={saveData}><div class="font-15pt">Uložit</div></Button>
                    {/if}
                    <Button scaleDown={true} onclick={back}><div class="font-15pt">Zpět</div></Button>
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

    .font-14pt {
        font-size: 14pt;
        @media screen and (max-width: 1062px) {
            font-size: 12pt;
        }
    }

    .font-15pt {
        font-size: 14pt;
        padding: 15px;

        @media screen and (max-width: 1163px) {
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

        @media screen and (max-width: 1062px) {
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

        & > div {
            height: calc(100dvh - var(--header-height) - 2 * var(--padding) - var(--button-bar-height));
            margin: auto;
        }
    }

    .edit-wrapper {
        font-size: 16pt;
        max-width: 500px;
        background-color: white;
        box-shadow: 2px 2px 20px #666;
        border-radius: 5px;
        display: flex;
        flex-direction: column;
        gap: 20px;
        padding: 20px;

        h2 {
            line-height: 0;
        }

        .save-options {
            display: grid;
            gap: 20px;

            grid-template-columns: 1fr 1fr;
            height: 40px;
            width: 300px;
        }
    }

    .sidebar {
        display: flex;
        flex-direction: column;
        justify-content: space-between;

        gap: 20px;
    }

    .stats {
        display: flex;
        flex-direction: column;
        gap: 20px;
        width: 100%;
        font-size: 16pt;

        & > div {
            background-color: #e7e7e7;
            padding: 10px 20px 20px 20px;
            border-radius: 10px;
        }
    }

    .controls {
        display: grid;
        grid-template-columns: 1fr 1fr 1fr;
        width: 100%;
        gap: 20px;
        justify-content: space-between;
    }
</style>
