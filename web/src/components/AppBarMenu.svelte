<script lang="ts">
    import { clickoutside } from "@svelte-put/clickoutside";

    export let x: number;
    export let close: () => void;

    let element: HTMLElement | null;
    let elementWidth: number | null;
    let screenWidth: number | null;
    let count = 0;

    let left: number;
    $: if (elementWidth && screenWidth) {
        // ensure element stays on screen
        left = Math.min(x, screenWidth - elementWidth / 2 - 12);
    }

    let clickOutside = () => {
        // do not close on the first click
        count = 1;
    };

    $: if (count === 1) {
        clickOutside = close;
    }
</script>

<svelte:window bind:innerWidth={screenWidth} />

<div
    class="absolute top-20 flex w-64 -translate-x-1/2 flex-col gap-2 rounded-lg border-2 p-2 dark:border-violet-900 dark:bg-neutral-900"
    style:left={`${left}px`}
    bind:this={element}
    bind:offsetWidth={elementWidth}
    use:clickoutside
    on:clickoutside={clickOutside}
>
    <slot />
</div>
