<script lang="ts">
    import Board from "$lib/components/Board.svelte";
    import { blur, fly } from "svelte/transition";
    import { page } from "$app/state";
    import { deleteGame, updateGame, type ApiError, type Game, type GameBase } from "$lib/api.js";
    import Button from "$lib/components/Button.svelte";
    import Inforow from "$lib/components/Inforow.svelte";
    import { goto } from "$app/navigation";
    import { formatDate, translateDifficulty, translateGameState } from "$lib/format";
    import Overlay from "$lib/components/Overlay.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import RadioButtons from "$lib/components/RadioButtons.svelte";
    import { circOut, sineInOut } from "svelte/easing";

    const { data }: { data: { gameData: Promise<[Game, ApiError | undefined]>, slug: string } } =
        $props();

    let game: Game | undefined = $state(undefined);
    let editingGameOverlayVisible: boolean = $state(false);
    let editedGameBase: GameBase | undefined = $state(undefined);
    let deleteGameOverlayVisible: boolean = $state(false);
    let mainClientWidth: number = $state(0);
    let gameInfoButtonVisible: boolean = $derived(mainClientWidth <= 1022);
    let gameInfoOverlayVisible: boolean = $state(false);

    async function processData(gameData: Promise<[Game, ApiError | undefined]>) {
        const [val, err] = await gameData;
        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }
        game = val;
    }

    async function onmove() {
        if (!game) return;

        const [val, err] = await updateGame({
            board: game?.board,
            difficulty: game?.difficulty,
            name: game?.name,
        }, game?.uuid);

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        game.updatedAt = val.updatedAt;
        game.gameState = val.gameState;

    }

    function copyText(text: string): () => void {
        return () => {
            navigator.clipboard.writeText(text);
        };
    }

    function startEditing() {
        if (!game) return;

        gameInfoOverlayVisible = false;

        editedGameBase = {
            board: game.board,
            difficulty: game.difficulty,
            name: game.name,
        }

        editingGameOverlayVisible = true;
    }

    function stopEditing() {
        editingGameOverlayVisible = false;
    }

    async function saveGameData() {
        if (!editedGameBase) return;
        if (!game) return;

        const [val, err] = await updateGame({
            board: editedGameBase.board,
            difficulty: editedGameBase.difficulty,
            name: editedGameBase.name,
        }, game.uuid);

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        game.gameState = val.gameState;
        game.name = val.name;
        game.difficulty = val.difficulty;
        stopEditing();
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
        goto("/game/");
    }

    $effect(() => {
        processData(data.gameData);
    })
    

    const permanentURL = page.url.host + page.url.pathname;
</script>


{#snippet stats(game: Game)}
<div class="stats">
    <Inforow key="Jméno hry" text={game.name} />
    <Inforow key="Obtížnost" text={translateDifficulty(game.difficulty)} />
    <Inforow key="Stav hry" text={translateGameState(game.gameState)} />
    <Inforow key="Vytvořeno" text={formatDate(game.createdAt)} />
    <Inforow key="Poslední úprava" text={formatDate(game.updatedAt)} />
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
        <Overlay bind:visible={editingGameOverlayVisible}>
            <div class="edit-wrapper" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
                <h2>Upravit údaje:</h2>
                <TextInput bind:value={editedGameBase!.name} placeholder="Jméno" />
                <div>Obtížnost:</div>
                <RadioButtons
                    bind:value={editedGameBase!.difficulty}
                    options={[
                        { display: "Začátečnická", option: "beginner" },
                        { display: "Jednoduchá", option: "easy" },
                        { display: "Středně těžká", option: "medium" },
                        { display: "Těžká", option: "hard" },
                        { display: "Extrémní", option: "extreme" },
                    ]}/>
                <div class="save-options">
                    <Button scaleDown={true} onclick={stopEditing}>
                        <div class="font-14pt">Zrušit</div>
                    </Button>
                    <Button scaleDown={true} variant="blue" onclick={saveGameData}>
                        <div class="font-14pt">Uložit</div>
                    </Button>
                </div>
            </div>
        </Overlay>
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
                <Button variant="blue" scaleDown={true} onclick={startEditing}><div class="font-15pt">Upravit</div></Button>
            </div>
        </Overlay>


        <div class="game">
            <div class="board-wrapper">
                <div transition:blur>
                    <Board {onmove} onwin={() => console.log("Win")} bind:value={game} />
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
                        <Button variant="blue" scaleDown={true} onclick={startEditing}><div class="font-15pt">Upravit</div></Button>
                    {/if}
                    <Button scaleDown={true} onclick={toggleDeleteOverlay}><div class="font-15pt">Smazat</div></Button>
                    
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
        @media screen and (max-width: 1022px) {
            font-size: 12pt;
        }
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

        @media screen and (max-width: 1022px) {
            display: flex;
            flex-direction: column;

            padding: 20px;
            --padding: 20px;
            --button-bar-height: 68px;
        }

        width: 100vw;
        height: calc(100vh - var(--header-height));
        --button-bar-height: 0px;
        --padding: 50px;
        padding: 50px;
    }

    .board-wrapper {
        overflow: hidden;
        height: calc(100vh - var(--header-height) - 2 * var(--padding) - var(--button-bar-height));

        & > div {
            height: calc(100vh - var(--header-height) - 2 * var(--padding) - var(--button-bar-height));
            margin: auto;
        }
    }

    .edit-wrapper {
        font-size: 16pt;
        max-width: 500px;
        //max-height: 500px;
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
</style>
