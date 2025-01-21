<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import type { User } from "$types/user";
    import { Icon, MagnifyingGlass } from "svelte-hero-icons";
    import AppBarMenu from "./AppBarMenu.svelte";

    export let user: User | null;

    let search = "";
    let showMenu = false;
    let menuAnchor: HTMLElement | null;

    const toggleMenu = () => {
        showMenu = !showMenu;
    };

    const closeMenu = () => {
        showMenu = false;
    };

    const signIn = () => {
        window.location.href = `${SERVER_URL}/auth/google`;
    };
</script>

<nav class="border-lavender flex h-16 items-center border-b-2 pr-8">
    <a class="text-lavender w-80 text-center text-2xl" href="/">LeetWars</a>
    <div class="relative ml-8 w-96">
        <input
            class="bg-surface0 w-full pl-10"
            placeholder="Search"
            bind:value={search}
        />
        <Icon
            class="text-overlay2 absolute left-0 top-0 ml-2 w-6"
            src={MagnifyingGlass}
        />
    </div>
    <div class="grow"></div>
    {#if user}
        <button
            class="h-10 w-10 rounded-full p-0 hover:shadow"
            on:click={toggleMenu}
            bind:this={menuAnchor}
        >
            <img
                class="rounded-full"
                width="100%"
                height="100%"
                src="https://assets.leetcode.com/users/mathletedev/avatar_1714888331.png"
                alt="profile_picture"
            />
        </button>
        {#if showMenu && menuAnchor}
            <AppBarMenu
                x={menuAnchor.offsetLeft + menuAnchor.offsetWidth / 2}
                close={closeMenu}
            >
                <a href={`${SERVER_URL}/signout/google`}>Sign out</a>
            </AppBarMenu>
        {/if}
    {:else}
        <button
            class="bg-blue h-10 rounded-full px-4 text-base hover:shadow"
            on:click={signIn}
        >
            Sign in
        </button>
    {/if}
</nav>
