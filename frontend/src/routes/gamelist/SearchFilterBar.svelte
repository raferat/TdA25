<script lang="ts">
    import type { Filters } from "./filter";

    let { 
        filters = $bindable(),
    }: {
        filters?: Filters
    } = $props();
</script>

<div id="outer">
    <div id="inner">
        <input type="text">
    </div>
</div>

<style lang="scss">
    #outer {
        position: sticky;
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        top: 0;
    }

    #inner {
        padding: 25px;
        display: flex;
        background-color: #f6f6f6;
        border: 1px solid #000;
    }
</style>