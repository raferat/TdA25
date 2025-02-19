<script lang="ts">
    import Dropdown from "$lib/components/Dropdown.svelte";
    import TextInput from "$lib/components/TextInput.svelte";
    import { translateDifficulty, translateDifficultyUnsafe, translateGameState, translateGameStateUnsafe, translateTimeFrameUnsafe } from "$lib/format";
    import type { Filters } from "./filter";

    let { 
        filters = $bindable({}),
    }: {
        filters?: Filters
    } = $props();

    filters.nameSearch = "";
</script>

<div class="grid grid-cols-1 sm:grid-cols-3 xl:flex xl:justify-center xl:items-center w-full gap-3 text-xs p-3">
    <div class="rounded-xl xl:min-w-[500px] sm:col-[1/4] col-1"><TextInput placeholder="Hledaný výraz" bind:value={filters.nameSearch}/></div>
    <div class="min-w-[200px] h-[54px]">
        <Dropdown>
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
            <ul class="dropdownul">
                <li onclick={_ => filters.difficulty = undefined}>Všechny</li>
                <li onclick={_ => filters.difficulty = "beginner"}>{translateDifficulty("beginner")}</li>
                <li onclick={_ => filters.difficulty = "easy"}>{translateDifficulty("easy")}</li>
                <li onclick={_ => filters.difficulty = "medium"}>{translateDifficulty("medium")}</li>
                <li onclick={_ => filters.difficulty = "hard"}>{translateDifficulty("hard")}</li>
                <li onclick={_ => filters.difficulty = "extreme"}>{translateDifficulty("extreme")}</li>
            </ul>
            {#snippet btn(toggle, icon)}
                <div style="display: flex; width: 100%; padding: 10px 30px 10px 30px; justify-content: space-between;">
                    <span>{translateDifficultyUnsafe(filters.difficulty) ?? "Obtížnost"}</span> {@render icon()}
                </div>
            {/snippet}
        </Dropdown>
    </div>
    <div class="min-w-[200px] h-[54px]">
        <Dropdown>
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
            <ul class="dropdownul">
                <li onclick={_ => filters.gameState = undefined}>Všechny</li>
                <li onclick={_ => filters.gameState = "opening"}>{translateGameState("opening")}</li>
                <li onclick={_ => filters.gameState = "midgame"}>{translateGameState("midgame")}</li>
                <li onclick={_ => filters.gameState = "endgame"}>{translateGameState("endgame")}</li>
            </ul>
            {#snippet btn(toggle, icon)}
                <div style="display: flex; width: 100%; padding: 10px 30px 10px 30px; justify-content: space-between;">
                    <span>{translateGameStateUnsafe(filters.gameState) ?? "Stav hry"}</span> {@render icon()}
                </div>
            {/snippet}
        </Dropdown>
    </div>
    <div class="min-w-[200px] h-[54px]">
        <Dropdown>
            <!-- svelte-ignore a11y_click_events_have_key_events -->
            <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
            <ul class="dropdownul">
                <li onclick={_ => filters.updatedAgo = undefined}>Všechny</li>
                <li onclick={_ => filters.updatedAgo = 1000*60*60*24}>Den</li>
                <li onclick={_ => filters.updatedAgo = 1000*60*60*24*7}>Týden</li>
                <li onclick={_ => filters.updatedAgo = 1000*60*60*24*31}>Měsíc</li>
                <li onclick={_ => filters.updatedAgo = 1000*60*60*24*31*3}>3 měsíce</li>
            </ul>
            {#snippet btn(toggle, icon)}
                <div style="display: flex; width: 100%; padding: 10px 30px 10px 30px; justify-content: space-between;">
                    <span>{translateTimeFrameUnsafe(filters?.updatedAgo)  ?? "Období"}</span> {@render icon()}
                </div>
            {/snippet}
        </Dropdown>
    </div>
</div>

<style lang="scss">
    .dropdownul {
        margin: 0;
        padding: 0;
        list-style-type: none;
        
        li {
            padding: 5px;
            margin-bottom: 2px;
            border-radius: 2px;


            &:first-child {
                border-top-left-radius: 7px;
                border-top-right-radius: 7px;
            }

            &:last-child {
                border-bottom-left-radius: 7px;
                border-bottom-right-radius: 7px;
            }

            &:hover {
                background-color: #72447920;
            }
        }    
    }
</style>