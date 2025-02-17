<script lang="ts">
    let { 
        value = $bindable(""), 
        placeholder = "",
    }: { 
        value?: string, 
        placeholder?: string,
    } = $props();
</script>

<label class="block relative" for="input">
    <input 
        class="size-full text-xl font-semibold border-2 border-transparent outline-none bg-tcream2 dark:bg-tlgrey p-5 pb-2 rounded-xl transition-all duration-150 ease-[cubic-bezier(0.445,0.05,0.55,0.95)]
            focus:rounded-bl-[0px] focus:rounded-br-[0px] focus:border-b-tblue" 
        id="input" 
        type="text" 
        bind:value={value}>
    <div class="cursor-text absolute text-xl top-[16px] left-[22px] transition-all duration-150 ease-[cubic-bezier(0.445,0.05,0.55,0.95)]" 
        class:topleft={value.length > 0}>{placeholder}</div>
</label>

<style>
    .topleft {
        top: 4px;
        left: 10px;
        font-size: 9.5pt;
    }
</style>