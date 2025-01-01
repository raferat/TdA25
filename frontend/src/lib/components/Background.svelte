<script lang="ts">
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

    let { height: containerHeight }: { height: number } = $props();

    const icons: Array<MyIcon> = [];
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
        "thinking",
        "x",
        "x",
        "logo-erb",
        "logo-erb",
    ];

    const cellWidth = 175;
    const cellHeight = 175;

    const iconWidth = 60;
    const iconHeight = 60;

    for (let xidx = 0; xidx < 50; xidx++) {
        for (let yidx = 0; yidx < 50; yidx++) {
            const cls = possibilities[Math.floor(Math.random() * possibilities.length)];
            const x = cellWidth * xidx + (cellWidth - iconWidth) * Math.random();
            const y = cellHeight * yidx + (cellHeight - iconHeight) * Math.random();
            icons.push({ x: x, y: y, iconClass: cls });
        }
    }
</script>

<div class="background" style="height: {containerHeight}px">
    {#each icons as icon}
        <div
            class="icon {icon.iconClass}"
            style="top: {icon.x}px; left: {icon.y}px"
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
