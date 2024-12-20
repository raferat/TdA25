<script lang="ts">
    import { listGames } from "$lib/api";
    import Board from "$lib/Board.svelte";

    let x = listGames();
</script>

<h1>Game</h1>

<Board/>

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