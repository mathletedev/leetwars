<script lang="ts">
    import { SERVER_URL } from "$lib/config";

    let username: string;

    const handleSubmit = async () => {
        if (!username) {
            return;
        }

        let res = await fetch(`${SERVER_URL}/api/username`, {
            credentials: "include",
            method: "POST",
            body: username,
        });

        if (res.status === 404) {
            alert("LeetCode username not found");
            return;
        }

        location.href = "/";
    };
</script>

<div class="flex h-full">
    <div class="m-auto flex w-72 flex-col items-center">
        <p class="text-xl">Enter your LeetCode username:</p>
        <input class="mt-4 w-full" name="username" bind:value={username} />
        <button class="mt-6 w-full" on:click={handleSubmit}>Let's go!</button>
    </div>
</div>
