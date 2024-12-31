
export function appearAfterDelay(node: HTMLElement, params: {delay: number} = {delay: 100}): {tick: (t: number) => void, duration: number } {
    const disp = window.getComputedStyle(node).display;
    return {
        duration: params.delay,
        tick: (t) => {
            if (t == 1)
                node.style.display = disp;
            else
                node.style.display = "none";
        },
    }
}