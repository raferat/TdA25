<script lang="ts">
    import { expoInOut, quadInOut } from "svelte/easing";
    import { crossfade, fade, scale } from "svelte/transition";
    import { appearAfterDelay } from "../transitions";

    let { 
        value = $bindable(""), 
        options = [],
    }: {
        value?: string, 
        options?: Array<{display: string, option: string}>
    } = $props();

    /*let selected = $derived(options.findIndex( (val, idx) => {
        return val.option == value;
    }));*/

    const select = (idx: number) => {
        return () => {
            value = options[idx].option;
        }
    }

    const [send,receive] = crossfade({
        duration: 250,
        easing: quadInOut,
    });
</script>

<div class="wrapper">
    {#each options as opt, idx (opt.option)}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div class="opt" onclick={select(idx)}>
        {#if value == opt.option}
        <div in:receive={{key: opt.option}}
             out:send={{key: value}}
            class="checkbox" 
            class:selectedO={idx <= options.length/2} 
            class:selectedX={idx > options.length/2}>
        </div>
        {:else}
        <div in:appearAfterDelay={{delay: 250}} class="checkbox"></div>
        {/if}
        <div class="option-content">{opt.display}</div>
    </div>
    {/each}
</div>

<style lang="scss">
    .wrapper {
        font-size: 85%;
    }
    .opt {
        cursor: pointer;
        display: flex;
        gap: 10px;
        justify-content: left;
        align-items: center;
    }

    .checkbox {
        width: 15px;
        height: 15px;
        border-radius: 3px;
    }

    .selectedX {
        background-image: url("/icons/X/X_cervene.svg");
    }

    .selectedO {
        background-image: url("/icons/O/O_modre.svg");
    }
</style>

