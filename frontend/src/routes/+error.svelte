<script lang="ts">
    import { goto } from "$app/navigation";
    import { page } from "$app/state";

    $effect(() => {
        goto(`/error/?obj=${encodeURIComponent(JSON.stringify(page.url))}`);
    })
</script>