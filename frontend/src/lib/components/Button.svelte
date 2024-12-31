<script lang="ts">
    import type { Snippet } from "svelte";

    let {
        children,
        variant = "red",
        href = undefined,
        onclick = undefined,
        scaleDown = false,
    }: {
        children: Snippet<[]>,
        variant?: "red" | "blue";
        href?: string;
        scaleDown?: boolean,
        onclick?: (a: MouseEvent) => void,
    } = $props();
</script>

{#if href}
<a {href} class:blue={variant == "blue"} class:red={variant == "red"} class:size={!scaleDown}>
    <center>
        {@render children()}
    </center>
</a>
{:else if onclick}
<button {onclick} class:blue={variant == "blue"} class:red={variant == "red"} class:size={!scaleDown}>
    <center>
        {@render children()}
    </center>
</button>
{:else}
<div class:blue={variant == "blue"} class:red={variant == "red"} class:size={!scaleDown}>
    <center>
        {@render children()}
    </center>
</div>
{/if}

<style lang="scss">
    @use 'sass:color';

    .size {
        width: 215px;
        height: 60px;
        font-size: 15pt;
        padding: 20px 60px;
    }

    a, button, div {
        display: block;
        border: none;
        background: none;

        
        margin: 0;

        color: #F6F6F6;
        text-decoration: none;
        cursor: pointer;
        border-radius: 10px;
        
        transition: all 150ms cubic-bezier(0.445, 0.05, 0.55, 0.95);
    }

    .red {
        background-color: #E31837;

        &:hover {
            background-color: color.adjust(#E31837, $lightness: -5%, $space: hsl);
        }
    }

    .blue {
        background-color: #0070BB;

        &:hover {
            background-color: color.adjust(#0070BB, $lightness: -5%, $space: hsl);
        }
    }
</style>