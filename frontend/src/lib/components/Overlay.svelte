<script lang="ts">
    import type { Snippet } from "svelte";
    import { fade } from "svelte/transition";

    let {
        visible = $bindable(false),
        children,
    }: {
        visible?: boolean,
        children?: Snippet<[]>,
    } = $props();

    const hide = () => {
        visible = false;
    }
</script>


{#if visible}
<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="overlay" onclick={hide} transition:fade>
    <div onclick={e => e.stopPropagation()}>
        {@render children?.()}
    </div>
</div>
{/if}

<style lang="scss">
    .overlay {
        position: absolute;
        z-index: 50;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        backdrop-filter: blur(5px);

        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>