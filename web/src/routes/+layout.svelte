<script lang="ts">
    import { onMount } from "svelte";
    import "../app.css";
    import AppBar from "$components/AppBar.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;

    onMount(() => {
        if (
            localStorage.theme === "dark" ||
            (!("theme" in localStorage) &&
                window.matchMedia("(prefers-color-scheme: dark)").matches)
        ) {
            document.documentElement.classList.add("dark");
        } else {
            document.documentElement.classList.remove("dark");
        }
    });
</script>

<AppBar user={data.user} />

<slot />
