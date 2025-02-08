<script lang="ts">
    import "../app.css";
    import Header from "$lib/components/Header.svelte";

    const { children } = $props();
</script>

<div class="absolute overflow-auto size-full z-0 bg-tcream text-tblack dark:bg-tblack dark:text-tcream">
    <Header/>
    {@render children?.()}
</div>
<!--
<div id="outer">
    <div bind:clientWidth={clientWidth} bind:clientHeight={clientHeight}>
        <Header/>
    </div>
    <div id="content" style="height: calc(100vh - {clientHeight}px); --header-height: {clientHeight}px;">
        <!--<Background height={childrenHeight} width={clientWidth}/>--><!--
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
</style>-->