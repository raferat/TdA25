<script lang="ts">
    import TextInput from "$lib/components/TextInput.svelte";

    let username: string = $state("");
    let password: string = $state("");

</script>

<main class="p-3 flex flex-col size-full justify-center items-center">
    <form onsubmit={(e) => e.preventDefault()} class="rounded-xl flex flex-col gap-3 max-w-100">
        <h2 class="text-3xl font-bold text-left p-8">Přihlášení</h2>
        <TextInput bind:value={username} placeholder="Uživatelské jméno / e-mail" />
        <TextInput bind:value={password} placeholder="Heslo" inputtype="password" />
        <div class="grid grid-cols-2">
            <button class="pbred text-xl sm:col-start-2 col-span-2">Přihlásit</button>
        </div>
        <div class="h-30"></div>
    </form>
</main>