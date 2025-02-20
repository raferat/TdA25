<script lang="ts">
    import "../app.css";
    import Header from "$lib/components/Header.svelte";

    const { children } = $props();

    $effect(() => {
        const themePreference = window.localStorage.getItem("theme");
        const darkTheme = (themePreference === "dark" || 
                          (themePreference === null && window.matchMedia("(prefers-color-scheme: dark)").matches));
        if (darkTheme) {
            document.documentElement.setAttribute("data-theme", "light")
        } else {
            document.documentElement.setAttribute("data-theme", "light")
        }
    })

    let headerHeight = $state(0);
</script>

<div class="absolute overflow-auto size-full z-0 bg-tcream text-tblack dark:bg-tblack dark:text-tcream">
    <div bind:clientHeight={headerHeight}>
        <Header/>
    </div>
    <div style="height: calc(100vh - {headerHeight}px)" class="relative w-full">
        {@render children?.()}
    </div>
</div>
