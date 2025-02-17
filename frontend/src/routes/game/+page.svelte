<script lang="ts">
    import { goto } from "$app/navigation";
    import { createGame, defaultGameBase } from "$lib/api";
    import Background from "$lib/components/Background.svelte";
    import Overlay from "$lib/components/Overlay.svelte";
    import RadioButtons from "$lib/components/RadioButtons.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import TicTacToe from "$lib/components/TicTacToe.svelte";
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
        boardAnimationController = (boardAnimationController + 1) % 5;
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

<Background width={0} height={0}/>
<main class="relative size-full overflow-hidden">
    <Overlay bind:visible={savingOverlayVisibe}>
        <div class="bg-white dark:bg-tblack text-xl p-8 flex flex-col gap-3 rounded-xl" in:fly={{ y: -500, easing: circOut }} out:fly={{ y: -500, easing: sineInOut }}>
            <h2 class="leading-none">Uložit hru:</h2>
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
            <div class="grid grid-cols-2 gap-1 h-8">
                <button class="pbred px-2 py-1" onclick={cancelSave}>Zrušit</button>
                <button class="pbblue px-2 py-1" onclick={confirmSave}>Uložit</button>
            </div>
        </div>
    </Overlay>

    <div class="size-full md:grid md:grid-cols-[minmax(0,0.7fr)_0.3fr] flex flex-col">
        {#key boardAnimationController}
        <div class="size-full" in:scale={{ delay: 250 }} out:scale={{duration: 250}}>
            <TicTacToe {onwin} bind:board={base.board} />
        </div>
        {/key}

        <div class="grid grid-cols-2 p-3 md:flex md:flex-col md:items-end md:justify-end md:gap-4">
            {#if playing}
                <button 
                    class="pbblue w-full text-xl" 
                    onclick={saveGame}
                    transition:slide={{ axis: "x", duration: 250 }}>Uložit</button>
            {:else}
                <div></div>
            {/if}
            <button class="pbred w-full text-xl" onclick={newGame}>
                <div transition:fade>
                    {#if playing}
                        Zahodit
                    {:else}
                        Nová hra
                    {/if}
                </div>
            </button>
        </div>
    </div>
</main>
