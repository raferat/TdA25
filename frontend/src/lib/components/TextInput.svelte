<script lang="ts">
    let { 
        value = $bindable(""), 
        placeholder = "",
        autofocus = false,
    }: { 
        value?: string, 
        placeholder?: string,
        autofocus?: boolean,
    } = $props();
</script>

<label for="input">
    <input id="input" class:longtext={value.length > 50} type="text" bind:value={value}>
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
        background-color: #00000010;
        padding: 20px;
        padding-bottom: 8px;
        border-radius: 10px;
        transition: all 150ms cubic-bezier(0.445, 0.05, 0.55, 0.95);
        

        &:focus {
            border-bottom-left-radius: 0px;
            border-bottom-right-radius: 0px;
            border-bottom: 2px solid #0070BB;
            background-color: #0808081B;
        }

        &:hover {
            background-color: #0808081A;
        }
    }

    .longtext:focus {
        border-bottom: 2px solid #E31837;
    }

    .placeholder {
        cursor: text;
        position: absolute;
        font-size: 13pt;
        top: 22px;
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