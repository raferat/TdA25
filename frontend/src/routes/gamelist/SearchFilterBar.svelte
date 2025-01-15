<script lang="ts">
    import Dropdown from "$lib/components/Dropdown.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import type { Filters } from "./filter";

    let { 
        filters = $bindable({}),
    }: {
        filters?: Filters
    } = $props();

    filters.nameSearch = "";
</script>

<div id="outer">
    <div id="nameSearch"><TextInput focuseffect={false} placeholder="Hledaný výraz" bind:value={filters.nameSearch}/></div>
    <div id="difficultyDropdown">
        <Dropdown>
            <p>A</p>
            <p>B</p>
            <p>C</p>

            {#snippet btn(toggle, icon)}
                <div style="display: flex; width: 100%; padding: 10px 30px 10px 30px; justify-content: space-between;"><span>Obtížnost</span> {@render icon()}</div>
                
            {/snippet}
        </Dropdown>
    </div>
</div>

<style lang="scss">
    #outer {
        position: sticky;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 20px;
        width: 100%;
        top: 0;
    }

    #nameSearch {
        width: 500px;
        border: 2px solid #0070BB;
        border-radius: 10px;
    }

    #difficultyDropdown {
        width: 200px;
        height: 54px;
    }
</style>