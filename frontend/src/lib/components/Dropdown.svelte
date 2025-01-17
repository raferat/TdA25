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
<span class="icon" class:downarrow={!visible} class:uparrow={visible}>&af;</span>
{/snippet}

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="dropdown" onclick={toggle}>
    {@render btn(toggle, icon)}
    {#if visible}
        <div class="drop" transition:slide={{axis: "y", duration: 250, easing: sineIn}}>
            {@render children()}
        </div>
    {/if}
</div>

<style>
    .dropdown {
        position: relative;
        background-color: #e7e7e7;
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        border-radius: 10px;
        cursor: pointer;
    }

    .drop {
        position: absolute;
        top: calc(100% - 10px);
        left: 0px;
        width: 100%;
        background-color: #e7e7e7;
        border-bottom-left-radius: 10px;
        border-bottom-right-radius: 10px;
        padding: 10px;
    }

    .icon {
        display: inline-block;
        background-size: 30px 30px;
        background-repeat: repeat;
        background-position: center;
        width: 24px;
        height: 24px;

        transition: all 250ms ease-in-out;
        background-image: url("/material-icons/dropdown-menu/dropdown.svg");
    }
    .downarrow {
        transform: rotateZ(0deg);
    }

    .uparrow {
        transform: rotateZ(180deg);
    }
</style>

