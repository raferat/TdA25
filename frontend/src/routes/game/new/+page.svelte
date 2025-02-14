<script>
    import { defaultGameBase } from "$lib/api";
    import TicTacToe from "$lib/components/TicTacToe.svelte";

</script>
<TicTacToe board={defaultGameBase().board}/>