<script lang="ts">
    import { goto } from "$app/navigation";
    import { register } from "$lib/api";
    import TextInput from "$lib/components/TextInput.svelte";
    import Overlay from "$lib/components/Overlay.svelte";

    let username: string = $state("");
    let email: string = $state("");
    let password: string = $state("");
    let password2: string = $state("");

    let passwordChecks = $derived({
        min8Letters: password.length >= 8,
        oneLowercaseLetter: password.match("[a-z]") != null,
        oneUppercaseLetter: password.match("[A-Z]") != null,
        oneNumber: password.match("[0-9]") != null,
        oneSpecialLetter: password.match("[^0-9a-zA-Z]") != null,
    });

    let errorOverlay = $state(false);

    async function register2() {
        if (!passwordChecks.min8Letters 
         || !passwordChecks.oneLowercaseLetter
         || !passwordChecks.oneUppercaseLetter
         || !passwordChecks.oneNumber
         || !passwordChecks.oneSpecialLetter) return;

        if (password != password2) return;

        if (await register(username, email, password)) {
            goto("/");
            return;
        }

        username = "";
        email = "";

        errorOverlay = true;
    }

</script>

<Overlay bind:visible={errorOverlay}>
    <div class="bg-tcream p-8 rounded-xl">
        <h2 class="font-bold text-xl mb-2">Přihlášení selhalo</h2>
        <div>Přihlašovací jméno nebo email jsou už používány někým jiným.</div>
    </div>
</Overlay>
<main class="p-3 flex flex-col size-full justify-center items-center">
    <form onsubmit={(e) => e.preventDefault()} class="rounded-xl flex flex-col gap-3 max-w-100">
        <h2 class="text-3xl font-bold text-left p-8">Registrace</h2>
        <TextInput bind:value={username} placeholder="Uživatelské jméno" />
        <TextInput bind:value={email} placeholder="Email" inputtype="email" />
        <TextInput bind:value={password} placeholder="Heslo" inputtype="password" />
        <TextInput bind:value={password2} placeholder="Zopakovat heslo" inputtype="password" />

        {#if !passwordChecks.min8Letters}
            <div class="mark">Heslo musí obsahovat nejméně 8 znaků.</div>
        {/if}
        {#if !passwordChecks.oneLowercaseLetter}
            <div class="mark">Heslo musí obsahovat nejméně 1 malé písmeno.</div>
        {/if}
        {#if !passwordChecks.oneUppercaseLetter}
            <div class="mark">Heslo musí obsahovat nejméně 1 velké písmeno.</div>
        {/if}
        {#if !passwordChecks.oneNumber}
            <div class="mark">Heslo musí obsahovat nejméně 1 číslo.</div>
        {/if}
        {#if !passwordChecks.oneSpecialLetter}
            <div class="mark">Heslo musí obsahovat nejméně 1 speciální znak.</div>
        {/if}
        {#if password != password2}
            <div class="mark">Heslo a zopakované heslo musí být stejné.</div>
        {/if}

        <div class="grid grid-cols-2">
            <button class="pbred text-xl sm:col-start-2 col-span-2" onclick={register2}>Registrovat</button>
        </div>
        <div class="h-30"></div>
    </form>
</main>


<style>
    .mark {
        margin-left: 8px;
        padding-left: 29px;

        background-image: url("/icons/X/X_cervene.svg");
        background-repeat: no-repeat;
        background-position: left;
        background-size: 20px 20px;
    }
</style>