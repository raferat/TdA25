<script>
    import Button from "./Button.svelte";
</script>

<header>
    <nav>
        <a href="/">Úvod</a>
        <a href="/game/">Hrát</a>
        <a href="/game/">Seznam úloh</a>
    </nav>
</header>

<style lang="scss">
    .mainlink {
        display: flex;
        border: 2px solid #1A1A1A;
        border-radius: 10px;
        padding: 20px 50px;
        h1 {
            text-align: center;
            font-size: 40pt;
            line-height: 0;
            margin: auto;
            margin-left: 0;
        }
    }

    header {
        height: 130px;
        padding: 24px;
        //background-color: white;
    }
    h3 {
        line-height: 0;
        text-decoration: none;
    }

    nav {
        display: flex;
        justify-content: space-between;
        width: 100%;
    }

    .links {
        display: grid;
        gap: 20px;
        grid-template-columns: 1fr 1fr;
    }

    a {
        display: flex;
        justify-content: left;
        align-items: center;
        flex: 1;
        gap: 50px;
        text-decoration: none;
        color: black;

        .logo {
            height: 72px;
            aspect-ratio: 3.21;
            background-repeat: no-repeat;
            background-image: url("/logo/Think-different-Academy_LOGO_oficialni_1.svg");
        }
    }
</style>
