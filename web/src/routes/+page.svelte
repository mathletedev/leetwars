<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import { onMount } from "svelte";

    let content = "Loading...";

    onMount(async () => {
        const res = await fetch(`${SERVER_URL}/api/me`, {
            credentials: "include",
        });

        if (res.status === 401) {
            content = "Hello, world!";
            return;
        }

        const data = await res.json();

        content = `Hello, ${data.username}!`;
    });
</script>

<p>{content}</p>
