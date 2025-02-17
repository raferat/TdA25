<script lang="ts">
    import type { Snippet } from "svelte";
    import { sineIn } from "svelte/easing";
    import { fade, slide } from "svelte/transition";

    let {
        btn,
        visible = $bindable(false),
        children,
    }: {
        btn: Snippet<[(evt: MouseEvent)=>void, Snippet<[]>]>,
        visible?: boolean,
        children: Snippet<[]>,
    } = $props();

    const toggle = (evt: MouseEvent) => {
        evt.stopPropagation();
        visible = !visible;
    }
</script>

{#snippet icon()}
<span 
    class="inline-block bg-[30px_30px] w-[24px] h-[24px] bg-no-repeat bg-center bg-[url(/material-icons/dropdown-menu/dropdown.svg)] dark:invert" 
    class:downarrow={!visible} 
    class:uparrow={visible}>&af;</span>
{/snippet}

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="relative text-base bg-tcream2 dark:bg-tlgrey size-full flex justify-center items-center rounded-xl cursor-pointer" onclick={toggle}>
    {@render btn(toggle, icon)}
    {#if visible}
        <div class="z-999 absolute w-full bg-tcream2 dark:bg-tlgrey top-[calc(100%_-_10px)] left-0 rounded-bl-xl rounded-br-xl p-2" transition:slide={{axis: "y", duration: 250, easing: sineIn}}>
            {@render children()}
        </div>
    {/if}
</div>

<style>
    .downarrow {
        transform: rotateZ(0deg);
    }

    .uparrow {
        transform: rotateZ(180deg);
    }
</style>

