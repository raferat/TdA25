<script>
    import Button from "./Button.svelte";

</script>
<header>
    <nav>
        <a href="/" aria-label="úvod">
            <div class="logo"></div>
        </a>
        
        
        <div class="links">
            <Button href="/game/">
                <h3>Hrát</h3>
            </Button>

            <Button variant="blue" href="/gamelist/">
                <h3>Seznam úloh</h3>
            </Button>
        </div>
    </nav>
</header>

<style lang="scss">
    header {
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