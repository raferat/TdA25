<script lang="ts">
    import { initLoginState, login } from "$lib/api";
    import TextInput from "$lib/components/TextInput.svelte";

    $effect(() => {
        initLoginState();
    });

    let username: string = $state("");
    let password: string = $state("");

    async function loginWithCreds() {
        console.log(await login(username, password));
    }

</script>

<main class="p-3 flex flex-col size-full justify-center items-center">
    <form onsubmit={(e) => e.preventDefault()} class="rounded-xl flex flex-col gap-3 max-w-100">
        <h2 class="text-3xl font-bold text-left p-8">Přihlášení</h2>
        <TextInput bind:value={username} placeholder="Uživatelské jméno / e-mail" />
        <TextInput bind:value={password} placeholder="Heslo" inputtype="password" />
        <div class="grid grid-cols-2">
            <button class="pbred text-xl sm:col-start-2 col-span-2" onclick={loginWithCreds}>Přihlásit</button>
        </div>
        <div class="h-30"></div>
    </form>
</main>