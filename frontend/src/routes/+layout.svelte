<script>
    import Header from "$lib/Header.svelte";

    const { children } = $props();
</script>
<Header/>
{@render children()}