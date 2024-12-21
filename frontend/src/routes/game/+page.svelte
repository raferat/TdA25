<script lang="ts">
    import { listGames } from "$lib/api";
    import Board from "$lib/Board.svelte";

</script>

<main>
    <Board/>
</main>

<style lang="scss">
    main {
        padding: 15px;
        max-width: 100%;
        max-height: calc(100% - 130px);
        aspect-ratio: 1;
    }
</style>

<!--
{#await x}
    <p>Loading</p>
{:then val}
    {#each val as elem (elem.uuid) }
        <p><a href="/game/{elem.uuid}">{elem.name}</a></p>
        <p>{elem.createdAt}</p>
        <p>{elem.updatedAt}</p>
        <p>{elem.uuid}</p>
        <p>{elem.difficulty}</p>
        <p>{elem.gameState}</p>
        <p>{elem.board}</p>
        <hr>
    {/each}
{/await}
-->