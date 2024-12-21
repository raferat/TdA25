<script lang="ts">
    import type { Snippet } from "svelte";

    let {
        children,
        variant = "red",
        href = undefined,
        onclick = undefined,
    }: {
        children: Snippet<[]>,
        variant?: "red" | "blue";
        href?: string;
        onclick?: (a: MouseEvent) => void;
        
    } = $props();
</script>

{#if href}
<a {href} class:blue={variant == "blue"} class:red={variant == "red"}>
    <center>
        {@render children()}
    </center>
</a>
{:else if onclick}
<button {onclick} class:blue={variant == "blue"} class:red={variant == "red"}>
    <center>
        {@render children()}
    </center>
</button>
{:else}
<div class:blue={variant == "blue"} class:red={variant == "red"}>
    <center>
        {@render children()}
    </center>
</div>
{/if}

<style lang="scss">
    a, button, div {
        display: block;
        border: none;
        background: none;
        padding: 20px 60px;
        margin: 0;


        color: #F6F6F6;
        //border: 2px solid #F6F6F6;
        text-decoration: none;
        cursor: pointer;
        border-radius: 10px;
        font-size: large;
    }

    .red {
        background-color: #E31837;
    }

    .blue {
        background-color: #0070BB;
    }
</style>