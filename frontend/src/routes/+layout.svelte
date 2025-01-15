<script lang="ts">
    import Background from "$lib/components/Background.svelte";
    import Header from "$lib/components/Header.svelte";

    const { children } = $props();

    let clientWidth: number = $state(0);
    let clientHeight: number = $state(0);

    let childrenHeight: number = $state(0);
</script>

<div id="outer">
    <div bind:clientWidth={clientWidth} bind:clientHeight={clientHeight}>
        <Header/>
    </div>
    <div id="content" style="height: calc(100vh - {clientHeight}px); --header-height: {clientHeight}px;">
        <Background height={childrenHeight} width={clientWidth}/>
        <div bind:clientHeight={childrenHeight}>
            {@render children()}
        </div>
    </div>
</div>

<style>
    #outer {
        width: 100vw;
        height: 100vh;
    }

    #content {
        width: 100vw;
        overflow: auto;
        position: relative;
    }
</style>