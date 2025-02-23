<script lang="ts">
    import { goto } from "$app/navigation";
    import { initLoginState, loginState } from "$lib/api";
    import Background from "$lib/components/Background.svelte";
    import Overlay from "$lib/components/Overlay.svelte";
    import TextInput from "$lib/components/TextInput.svelte";

    const gotoLinkFunctor = (link: string) => {
        return () => {
            goto(link);
        };
    };

    let joiningGame: boolean = $state(false);
    let joinGameCode: string = $state("");

    $effect(initLoginState);
</script>

<Background height={0} width={0}/>

<a href="/freeplay/default/" class="hidden">Pentacosihexacontahexadecahecatondodecaexon</a>
<main class="size-full grid grid-cols-2 auto-rows-fr p-6 gap-6">
    <Overlay bind:visible={joiningGame}>
        <div class="bg-tcream p-6 rounded-xl">
            <h2 class="text-xl font-bold mb-5">Zadejte kód hry</h2>
            <TextInput placeholder="Kód" bind:value={joinGameCode}/>
            <div class="flex justify-end mt-5">
                <button class="pbblue text-[13pt]" onclick={_=> goto(`/freeplay/${joinGameCode}`)}>Přípojit</button>
            </div>
        </div>
    </Overlay>
    <div class="border-2 rounded-xl flex flex-col justify-between">
        <h2 class="text-2xl text-center leading-10 font-semibold">Lokální multiplayer</h2>
        <button class="pbblue m-8 text-xl" onclick={gotoLinkFunctor("/game/")}>Hrát</button>
    </div>
    <div class="border-2 rounded-xl flex flex-col justify-between">
        <h2 class="text-2xl text-center leading-10 font-semibold">Úlohy</h2>
        <button class="pbred m-8 text-xl" onclick={gotoLinkFunctor("/gamelist/")}>Přejít na seznam úloh</button>
    </div>
    <div class="border-2 rounded-xl flex flex-col justify-between">
        <h2 class="text-2xl text-center leading-10 font-semibold">Volná hra</h2>
        <div class="grid grid-cols-2 auto-rows-fr">
            <button class="pbblue m-8 text-xl disabled:bg-gray-500" disabled={$loginState == undefined} onclick={gotoLinkFunctor("/freeplay/")}>Založit hru</button>
            <button class="pbred m-8 text-xl" onclick={_ => joiningGame = !joiningGame}>Připojit se do hry</button>
        </div>
    </div>
    <div class="border-2 rounded-xl flex flex-col justify-between">
        <h2 class="text-2xl text-center leading-10 font-semibold">Hodnocená hra</h2>
        <button class="pbblue m-8 text-xl disabled:bg-gray-500" disabled={$loginState == undefined} onclick={gotoLinkFunctor("/match/")}>Hrát</button>
    </div>
</main>