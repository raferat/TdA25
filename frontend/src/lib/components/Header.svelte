<header class="w-screen p-7 flex justify-between">
    <a href="/" aria-label="Logo link pointing about page">
        <span class="block h-21 w-67 
            bg-[url(/logo/Think-different-Academy_LOGO_oficialni_1.svg)]
            dark:bg-[url(/logo/Think-different-Academy_LOGO_oficialni_1_dark-mode.svg)]"></span>
    </a>

    <nav class="flex-1 flex justify-end items-center gap-50 text-2xl pr-7 font-[600]">
        <a href="/">Úvod</a>
        <a href="/game/">Nová hra</a>
        <a href="/gamelist/">Úlohy</a>
    </nav>
</header>

<!-- Random background -->
{#each [
        [10,100],
        [20,18],
        [30,80],
        [50,2],
        [75,80],
        [90,30],
    ] as a}
    <img style="top: {a[1]}px; left: {a[0]}%" class="absolute size-10 blur-2xl" src="/icons/O/O_modre.svg" alt="">
{/each}

{#each [
        [15,10],
        [25,90],
        [40,70],
        [50,100],
        [62,50],
        [93,95],
    ] as a}
    <img style="top: {a[1]}px; left: {a[0]}%" class="absolute size-10 blur-2xl" src="/icons/X/X_cervene.svg" alt="">
{/each}