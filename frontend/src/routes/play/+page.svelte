<script lang="ts">
    import { goto } from "$app/navigation";
    import { initLoginState, loginState } from "$lib/api";
    import Background from "$lib/components/Background.svelte";

    const gotoLinkFunctor = (link: string) => {
        return () => {
            goto(link);
        };
    };

    $effect(initLoginState);
</script>

<Background height={0} width={0}/>

<main class="size-full grid grid-cols-2 auto-rows-fr p-6 gap-6">
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
            <button class="pbblue m-8 text-xl disabled:bg-gray-500" disabled={$loginState == undefined}>Založit hru</button>
            <button class="pbred m-8 text-xl">Připojit se do hry</button>
        </div>
    </div>
    <div class="border-2 rounded-xl flex flex-col justify-between">
        <h2 class="text-2xl text-center leading-10 font-semibold">Hodnocená hra</h2>
        <button class="pbblue m-8 text-xl disabled:bg-gray-500" disabled={$loginState == undefined}>Hrát</button>
    </div>
</main>