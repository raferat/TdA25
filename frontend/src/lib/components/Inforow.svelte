<script lang="ts">
    import type { Snippet } from "svelte";
    import { fade } from "svelte/transition";

    let { 
        key,
        text,
        minify = false,
        renderer,
        children,
    }: { 
        key: string,
        minify?: boolean,
        text?: string,
        renderer?: Snippet<[string]>,
        children?: Snippet<[]>,
    } = $props();


    const minification = (a: string) => {
        const len = 25;
        let result = a.slice(0, len/2);
        result += "....";
        result += a.slice(a.length-len/2);
        return result;
    }


    if (minify && text)
        text = minification(text);
</script>

<div class="relative w-full p-8 bg-tcream dark:bg-tblack pb-3 rounded-xl overflow-x-auto">
    <label for="child" class="absolute top-[8px] left-[8px]">{key}:</label>
    <div id="child" class="max-w-full text-xl font-medium text-nowrap" transition:fade>
        {#if renderer && text}
            {@render renderer(text)}
        {:else if children}
            {@render children()}
        {:else if text}
            {text}
        {/if}
    </div>
</div>

