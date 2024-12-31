<header>
    <nav>
        <a href="/" aria-label="Logo link pointing about page"
            ><span class="logo"></span></a
        >
        <span></span>
        <a href="/">Úvod</a>
        <a href="/game/">Nová hra</a>
        <a href="/gamelist/">Úlohy</a>
    </nav>
</header>

<style lang="scss">
    @media (min-width: 930px) {
        header {
            height: 130px;
            padding: 26px;
            background-color: #f6f6f6;

            nav {
                display: flex;
                align-items: center;
                font-size: 20pt;
                gap: 50px;

                span:not(.logo) {
                    flex: 1;
                }

                a:not(:has(.logo)) {
                    text-decoration: none;
                    color: black;
                    width: 200px;
                    height: 100%;
                    text-align: center;
                }
            }
        }

        a:has(.logo) {
            display: inline-block;
            height: 77px;
            .logo {
                display: inline-block;
                width: 258px;
                height: 77px;

                background-repeat: no-repeat;
                background-image: url("/logo/Think-different-Academy_LOGO_oficialni_1.svg");

                background-size: 258px 77px;
            }
        }
    }

    @media (max-width: 930px) and (min-width: 640px) {
        header {
            height: 70px;
            padding: 14px;
            background-color: #f6f6f6;

            nav {
                display: flex;
                align-items: center;
                font-size: 16pt;
                gap: 20px;

                span:not(.logo) {
                    flex: 1;
                }

                a:not(:has(.logo)) {
                    text-decoration: none;
                    color: black;
                    width: 120px;
                    text-align: center;
                }
            }
        }

        a:has(.logo) {
            display: inline-block;
            height: 42px;
            .logo {
                display: inline-block;
                width: 140px;
                height: 42px;

                background-repeat: no-repeat;
                background-image: url("/logo/Think-different-Academy_LOGO_oficialni_1.svg");

                background-size: 140px 42px;
            }
        }
    }

    @media (max-width: 640px) {
        header {
            height: 70px;
            padding: 14px;
            background-color: #f6f6f6;

            nav {
                display: flex;
                align-items: center;
                font-size: 12pt;
                gap: 20px;

                span:not(.logo) {
                    flex: 1;
                }

                a:not(:has(.logo)) {
                    text-decoration: none;
                    color: black;
                    width: 90px;
                    text-align: center;
                }
            }
        }

        a:has(.logo) {
            display: inline-block;
            height: 42px;
            .logo {
                display: inline-block;
                width: 140px;
                height: 42px;

                background-repeat: no-repeat;
                background-image: url("/logo/Think-different-Academy_LOGO_oficialni_1.svg");

                background-size: 140px 42px;
            }
        }
    }
</style>
