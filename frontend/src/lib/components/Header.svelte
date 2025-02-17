<script>
    import Overlay from "./Overlay.svelte";

    let overlayVisible = $state(false);

</script>
<header class="w-screen p-7 flex justify-between">
    <a href="/" aria-label="Logo link pointing about page">
        <span class="block md:h-21 md:w-67 h-7 w-22 
            bg-[url(/logo/Think-different-Academy_LOGO_oficialni_1.svg)]
            dark:bg-[url(/logo/Think-different-Academy_LOGO_oficialni_1_dark-mode.svg)]"></span>
    </a>

    <button class="sm:hidden text-3xl cursor-pointer" onclick={() => {overlayVisible = !overlayVisible;}}>...</button>

    <nav class="hidden flex-1 sm:flex justify-end items-center lg:gap-18 lg:text-2xl pr-7 font-bold text-base gap-8">
        <a href="/">Úvod</a>
        <a href="/game/">Hrát</a>
        <a href="/gamelist/">Úlohy</a>
        <a href="/login/">Přihlášení</a>
        <a href="/register/">Registrace</a>
    </nav>

    <Overlay bind:visible={overlayVisible}>
        <nav class="flex flex-col p-7 items-center font-bold text-xl gap-8">
            <a href="/">Úvod</a>
            <a href="/game/">Hrát</a>
            <a href="/gamelist/">Úlohy</a>
            <a href="/login/">Přihlášení</a>
            <a href="/register/">Registrace</a>
        </nav>
    </Overlay>
</header>