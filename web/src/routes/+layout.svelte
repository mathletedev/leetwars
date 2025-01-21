<script lang="ts">
    import "../app.css";
    import AppBar from "$components/AppBar.svelte";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<AppBar user={data.user} />

<div class="h-[calc(100vh-64px)]">
    <slot />
</div>
