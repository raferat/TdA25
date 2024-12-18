<script lang="ts">
    import { findGame } from '$lib/api.js';

    const { data } = $props();

    const x = findGame(data.slug);
</script>

<p>Slug: `{data.slug}`</p>

{#await x}
    <p>Loading</p>
{:then elem}
    <p><a href="/game/{elem.uuid}">{elem.name}</a></p>
    <p>{elem.createdAt}</p>
    <p>{elem.updatedAt}</p>
    <p>{elem.uuid}</p>
    <p>{elem.difficulty}</p>
    <p>{elem.gameState}</p>
    <p>{elem.board}</p>
    <hr>
{/await}