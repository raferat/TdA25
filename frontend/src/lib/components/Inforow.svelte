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

<div class="outer">
    <label for="child">{key}:</label>
    <div id="child" transition:fade>

        {#if renderer && text}
            {@render renderer(text)}
        {:else if children}
            {@render children()}
        {:else if text}
            {text}
        {/if}
    </div>
</div>

<style lang="scss">
    .outer {
        width: 100%;
        position: relative;
        background-color: #00000010;
        padding: 25px;
        padding-bottom: 10px;
        border-radius: 10px;
        overflow-x: auto;

        label {
            position: absolute;
            top: 8px;
            left: 8px;
            
            font-size: 10pt;
        }

        div {
            max-width: 100%;
            font-size: 15pt;
            text-wrap: nowrap;
            font-weight: 500;
        }
    }
</style>

