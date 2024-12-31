<script lang="ts">
    import { goto } from "$app/navigation";
    import { createGame, defaultGameBase } from "$lib/api";
    import Board from "$lib/components/Board.svelte";
    import Button from "$lib/components/Button.svelte";
    import Overlay from "$lib/components/Overlay.svelte";
    import RadioButtons from "$lib/components/RadioButtons.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import { circOut, sineInOut } from "svelte/easing";
    import { fade, fly, scale, slide } from "svelte/transition";

    let playing = $state(true);
    let savingOverlayVisibe = $state(false);

    let boardAnimationController = $state(0);

    const onwin = () => {
        playing = false;
    };

    let base = $state(defaultGameBase());

    const newGame = () => {
        base = defaultGameBase();
        if (playing) {
            boardAnimationController = (boardAnimationController + 1) % 5;
        }
        playing = true;
    };

    const saveGame = () => {
        savingOverlayVisibe = true;
    };

    const cancelSave = () => {
        savingOverlayVisibe = false;
    };

    const confirmSave = async () => {
        const [game, err] = await createGame(base);

        if (err != undefined) {
            goto(`/error/?obj=${encodeURIComponent(JSON.stringify(err))}`);
            return;
        }

        goto(`/game/${game.uuid}/`);
    };
</script>

<main>
    <Overlay bind:visible={savingOverlayVisibe}>
        <div class="save-wrapper" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
            <h2>Uložit hru:</h2>
            <TextInput bind:value={base.name} placeholder="Jméno" />
            <div>Obtížnost:</div>
            <RadioButtons
                bind:value={base.difficulty}
                options={[
                    { display: "Začátečnická", option: "beginner" },
                    { display: "Jednoduchá", option: "easy" },
                    { display: "Středně těžká", option: "medium" },
                    { display: "Těžká", option: "hard" },
                    { display: "Extrémní", option: "extreme" },
                ]}/>
            <div class="save-options">
                <Button scaleDown={true} onclick={cancelSave}>
                    <div class="font-14pt">Zrušit</div>
                </Button>
                <Button scaleDown={true} variant="blue" onclick={confirmSave}>
                    <div class="font-14pt">Uložit</div>
                </Button>
            </div>
        </div>
    </Overlay>

    <div class="game">
        <div></div>

        <div class="board-wrapper">
            {#key boardAnimationController}
                <div in:scale={{ delay: 250 }} out:fly={{ y: 500, duration: 250 }}>
                    <Board {onwin} bind:value={base} />
                </div>
            {/key}
        </div>

        <div class="controls">
            {#if playing}
                <div transition:slide={{ axis: "x", duration: 250 }}>
                    <Button variant="blue" onclick={saveGame}>Uložit</Button>
                </div>
            {/if}
            <Button onclick={newGame}>
                <div transition:fade>
                    {#if playing}
                        Zahodit
                    {:else}
                        Nová hra
                    {/if}
                </div>
            </Button>
        </div>
    </div>
</main>

<style lang="scss">
    .font-14pt {
        font-size: 14pt;
    }

    main {
        position: relative;
    }

    .save-wrapper {
        font-size: 16pt;
        max-width: 500px;
        max-height: 500px;
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

    .game {
        display: grid;
        grid-template-columns: 0.3fr 0.8fr 0.3fr;

        width: 100vw;
        height: calc(100vh - 130px);

        padding: 50px;
    }

    .board-wrapper {
        overflow: hidden;
        height: calc(100vh - 230px);

        & > div {
            height: calc(100vh - 230px);
            margin: auto;
        }
    }

    .controls {
        display: flex;
        flex-direction: column;
        justify-content: end;
        align-items: end;
        gap: 20px;
    }
</style>
