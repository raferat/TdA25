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
<div class="overlay backdrop-brightness-90 dark:backdrop-brightness-50 backdrop-blur-sm z-50 absolute top-0 left-0 size-full flex justify-center items-center" onclick={hide} transition:fade>
    <div onclick={e => e.stopPropagation()}>
        {@render children?.()}
    </div>
</div>
{/if}

<style lang="scss">
    
</style>