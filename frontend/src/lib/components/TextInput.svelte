<script lang="ts">
    let { 
        value = $bindable(""), 
        placeholder = "",
        focuseffect = true,
        autofocus = false,
    }: { 
        value?: string, 
        placeholder?: string,
        autofocus?: boolean,
        focuseffect?: boolean,
    } = $props();
</script>

<label for="input">
    <input id="input" class:focuseffect={focuseffect} class:longtext={value.length > 50} type="text" bind:value={value}>
    <div class="placeholder" class:topleft={value.length > 0}>{placeholder}</div>
</label>

<style lang="scss">
    label {
        display: block;
        position: relative;
    }

    input {
        width: 100%;
        height: 100%;
        font-size: 13pt;
        font-weight: 500;
        border: 2px solid transparent;
        outline: none;
        background-color: #e7e7e7;
        padding: 20px;
        padding-bottom: 8px;
        border-radius: 10px;
        transition: all 150ms cubic-bezier(0.445, 0.05, 0.55, 0.95);
        

        &.focuseffect:focus {
            border-bottom-left-radius: 0px;
            border-bottom-right-radius: 0px;
            border-bottom: 2px solid #0070BB;
            background-color: #cacaca;
        }

        &:hover {
            background-color: #c9c9c9;
        }
    }

    .longtext:focus {
        border-bottom: 2px solid #E31837;
    }

    .placeholder {
        cursor: text;
        position: absolute;
        font-size: 13pt;
        top: 16px;
        left: 22px;
        color: #080808AA;

        transition: all 150ms cubic-bezier(0.445, 0.05, 0.55, 0.95);
    }

    .topleft {
        top: 4px;
        left: 10px;
        font-size: 9.5pt;
    }
</style>