<script lang="ts">
    import "../app.css";
    import Header from "$lib/components/Header.svelte";

    const { children } = $props();
</script>

<div class="absolute overflow-auto size-full z-0 bg-tcream text-tblack dark:bg-tblack dark:text-tcream">
    <Header/>
    {@render children?.()}
</div>
