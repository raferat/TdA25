<script lang="ts">
    import { listGames } from "$lib/api";

    let x = listGames();
</script>

<h1>Game</h1>

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