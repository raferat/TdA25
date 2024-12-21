<script>
    import Header from "$lib/Header.svelte";

    const { children } = $props();
</script>
<div id="content">
<Header/>
{@render children()}
</div>

<style>
    #content {
        width: 100vw;
        height: 100vh;
    }
</style>