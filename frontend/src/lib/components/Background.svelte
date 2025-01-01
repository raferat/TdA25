<script lang="ts">
    import { untrack } from "svelte";
    import { flip } from "svelte/animate";
    import { sineInOut } from "svelte/easing";
    import { fade } from "svelte/transition";

    type IconClass =
        | "beginner"
        | "duck"
        | "easy"
        | "extreme"
        | "hard"
        | "idea"
        | "medium"
        | "o"
        | "playing"
        | "thinking"
        | "x"
        | "logo-text1"
        | "logo-text2"
        | "logo-erb";

    interface MyIcon {
        x: number;
        y: number;
        iconClass: IconClass;
    }

    let { height: containerHeight, width: containerWidth }: { height: number, width: number } = $props();

    let lastH = 0;
    let lastW = 0;

    let icons: Array<MyIcon> = $state([]);
    const possibilities: ReadonlyArray<IconClass> = [
        "beginner",
        "duck",
        "duck",
        "easy",
        "extreme",
        "hard",
        "idea",
        "idea",
        "medium",
        "o",
        "o",
        "logo-text1",
        "logo-text2",
        "playing",
        "playing",
        "thinking",
        "thinking",
        "x",
        "x",
        "logo-erb",
        "logo-erb",
    ];

    async function createIconicBackground(containerHeight: number, containerWidth: number) {
        const cellWidth = 175;
        const cellHeight = 175;

        const iconWidth = 60;
        const iconHeight = 60;

        const rowCnt = Math.ceil(containerHeight/cellHeight);
        const colCnt = Math.ceil(containerWidth/cellWidth);

        const localIcons: Array<MyIcon> = [];

        for (let xidx = 0; xidx < colCnt; xidx++) {
            for (let yidx = 0; yidx < rowCnt; yidx++) {
                const cls = possibilities[Math.floor(Math.random() * possibilities.length)];
                const x = cellWidth * xidx + (cellWidth - iconWidth) * Math.random();
                const y = cellHeight * yidx + (cellHeight - iconHeight) * Math.random();
                localIcons.push({ x: x, y: y, iconClass: cls });
            }
        }

        icons = localIcons;
    }

    $effect(() => {
        let h = Math.max(containerHeight, window.screen.height);
        let w = Math.max(containerWidth, window.screen.width);
        untrack(() => {
            if ( h == lastH && w == lastW ) return;
            lastH = h;
            lastW = w;
            createIconicBackground(h, w);
        });
    });
</script>

<div class="background" style="height: {containerHeight}px">
    {#each icons as icon, idx (idx)}
        <div
            animate:flip={{delay: 5+idx, duration: 1000, easing: sineInOut}}
            transition:fade
            class="icon {icon.iconClass}"
            style="top: {icon.y}px; left: {icon.x}px"
        ></div>
    {/each}
</div>

<!-- svelte-ignore css_unused_selector -->
<style lang="scss">
    .background {
        position: absolute;
        user-select: none;
        top: 0;
        left: 0;
        width: 100%;
        min-height: 100%;
        z-index: -5000;
        overflow: hidden;
    }

    .icon {
        position: absolute;
        width: 48px;
        height: 48px;
        background-repeat: no-repeat;
        background-size: 48px 48px;
        filter: opacity(7%);
        background-position: center;
        transition: background-image 500ms ease-in-out;

        &.beginner {
            background-image: url("/icons/Beginner/zarivka_beginner_cerne.svg");
        }

        &.duck {
            background-image: url("/icons/Duck/flag{p1skw0rky-4r3-n0t-d34d}_2.svg");
            background-size: 35px 35px;
        }

        &.easy {
            background-image: url("/icons/Easy/zarivka_easy_cerne.svg");
        }

        &.extreme {
            background-image: url("/icons/Extreme/zarivka_extreme_cerne.svg");
        }

        &.hard {
            background-image: url("/icons/Hard/zarivka_hard_cerne.svg");
        }

        &.idea {
            background-image: url("/icons/Idea/zarivka_idea_cerne.svg");
        }

        &.medium {
            background-image: url("/icons/Medium/zarivka_medium_cerne.svg");
        }

        &.o {
            background-image: url("/icons/O/O_cerne.svg");
            background-size: 35px 35px;
        }

        &.playing {
            background-image: url("/icons/Playing/zarivka_playing_cerne.svg");
        }

        &.thinking {
            background-image: url("/icons/Thinking/zarivka_thinking_cerne.svg");
        }

        &.x {
            background-image: url("/icons/X/X_cerne.svg");
            background-size: 35px 35px;
        }

        &.logo-text1 {
            background-image: url("/logo/Think-different-Academy_LOGO_textove-cerne.svg");
            width: 60px;
            height: 60px;
            background-size: 60px 60px;
        }

        &.logo-text2 {
            background-image: url("/logo/Think-different-Academy_LOGO_textove-cerne2.svg");
            width: 60px;
            height: 60px;
            background-size: 60px 60px;
        }

        &.logo-erb {
            background-image: url("/logo/Think-different-Academy_LOGO_cerny.svg");
        }
    }
</style>
