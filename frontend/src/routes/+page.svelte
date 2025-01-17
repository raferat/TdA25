<script lang="ts">
    import type { GameBase } from "$lib/api";
    import Background from "$lib/components/Background.svelte";
    import Board from "$lib/components/Board.svelte";

    let board1: GameBase = {
        name: "A",
        difficulty: "beginner",
        board: [
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "X", "", "", "", "X", "", "", "", "", ""],
            ["", "", "", "", "", "X", "", "", "", "X", "", "", "", "", ""],
            ["", "", "", "", "", "X", "", "", "", "X", "", "", "", "", ""],
            ["", "", "", "", "", "X", "", "", "", "X", "", "", "", "", ""],
            ["", "", "", "", "", "X", "", "", "", "X", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "O", "", "", "", "", "", "", "", "", "", "O", "", ""],
            ["", "", "O", "", "", "", "", "", "", "", "", "", "O", "", ""],
            ["", "", "", "O", "O", "O", "O", "O", "O", "O", "O", "O", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""]
        ],
    };

    let board2: GameBase = {
        name: "A",
        difficulty: "beginner",
        board: [
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "O", "O", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "O", "", "", "O", "", "", "", "", "", "", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "X", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "X", "", "", ""],
            ["", "O", "O", "O", "O", "O", "O", "", "", "X", "X", "X", "X", "X", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "X", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "X", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "", "", "", ""],
            ["", "O", "", "", "", "", "O", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
            ["", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
        ],
    };

</script>
<svelte:head>
    <meta 
        name="description" 
        content="Piškovrková platforma, která přináší zábavu i poučení. Rozdejte si partii piškvorek s přáteli, nebo zkuste vyřešit jednu z našich úloh.
            Organizace: Think different Academy">
</svelte:head>
<a href="/game/default" style="display: none;">Hello TdA</a>
<a href="/game/default/edit" style="display: none;">Hello TdA</a>
<Background width={0} height={0}/>

<main>
    <center class="hero">
        <h1>Jsme
            <div class="big-icon"></div>
        </h1>
    </center>

    <article>
        <div>
        <h2>Přinášíme Vám piškvorkovou platformu</h2>
        Piškvorková platforma je plná zábavy díky:
        <ul>
            <li>inovativnímu rozhraní</li>
            <li>skvělému grafickému designu</li>
            <li>velkému počtu úloh všech obtížností</li>
            <li>jednoduchosti hraní</li>
            <li>a mnohem víc...</li>
        </ul>
        Zahrajte si partii piškvorek s kamarády, klikněte na tlačítko <strong>Nová hra</strong>.
        </div>
        <div class="board-limiter1">
            <Board controls={false} value={board1}/>
        </div>
    </article>

    <article>
        <div class="board-limiter2">
            <Board controls={false} value={board2}/>
        </div>
        <div>
            <h2>Natrénujte analytické myšlení</h2>
            <p>Jedním z klíčových aspektů piškvorek je nutnost analyzovat možné tahy. Hráči musí zvažovat nejen své vlastní strategie, ale také se snažit odhadnout, jaké kroky podnikne jejich soupeř. Tím se rozvíjí schopnost kritického myšlení, protože hráči musí hodnotit různé scénáře a vybrat ten nejefektivnější. Tato dovednost je přenositelná do mnoha dalších oblastí života, včetně řešení problémů v akademickém nebo pracovním prostředí.</p>
            <p>Hraní piškvorek také podporuje trpělivost a sebekontrolu. Hráči se učí, že ne vždy vyhrají na první pokus, a musí se naučit analyzovat své chyby a zlepšovat se. Tato reflexe a schopnost učit se z neúspěchů jsou klíčové pro rozvoj analytického myšlení.</p>
            <p>V neposlední řadě, piškvorky mohou být skvělým nástrojem pro rozvoj sociálních dovedností. Hraní s ostatními lidmi vyžaduje komunikaci a spolupráci, což dále posiluje schopnost analyzovat situace z různých úhlů pohledu.</p>
            Procvičte se na některé z našich úloh, klikněte na tlačítko <strong>Úlohy</strong>.
        </div>
            
        
        
    </article>

    <div style="height: 1px;"></div>
</main>

<style lang="scss">
    main {
        width: 100%;
        height: calc(100vh - var(--header-height) - 25px);
        padding: 50px;
    }

    .hero {
        background-color: #F6F6F6;
    }
    .big-icon {
        width: 300px;
        height: 300px;

        background-image: url("/logo/Think-different-Academy_LOGO_oficialni.svg");
        background-repeat: no-repeat;
        background-position: center;
    }

    .board-limiter1 {
        margin-left: auto;
        margin-right: 50px;
        width: 375px;
    }

    .board-limiter2 {
        margin-left: 50px;
        margin-right: auto;
        width: 420px;
    }
    article {
        margin-bottom: 20px;
        background-color: #F6F6F6;
        border-radius: 10px;
        border: 2px solid;
        padding: 20px;
        display: grid;
        gap: 20px;
        grid-template-columns: repeat(2, minmax(0, 1fr));
        font-size: 17pt;

        &:nth-child(2n) {
            border-color: #E31837;
        }

        &:nth-child(2n+1) {
            border-color: #0070BB;
        }

        h2 {
            margin-top: 5px;
            font-size: 20pt;
        }
    }
</style>