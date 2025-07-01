<script>
    import {user} from "$lib/stores/auth.js"
    import Icon from "@iconify/svelte";
    let buttons = [
        {
            name: "Forums",
            link: "/content"
        },
        {
            
        }
    ]
</script>
<div class="grid grid-cols-3 w-full items-center px-6 py-3 transition-all duration-75 h-[25vh]">
    <!-- Left: Nav Links -->

    <!-- Center: Logo / Brand -->
    <div class="text-xl font-bold text-gray-900 tracking-tight text-center">
        chymein.
    </div>
    <div>

    </div>
    <!-- Right: User Area -->
    <div class="flex justify-center items-center gap-4 uppercase font-black ">
        {#if $user}
            <p>Welcome, {$user.Username}</p>
        {:else}
            <a href="/login" class="text-sm text-gray-700 hover:text-gray-900 transition">Login</a>
            <a href="/register" class="bg-gray-800 text-white text-sm px-3 py-1.5 rounded hover:bg-gray-900 transition">
                Sign Up
            </a>
        {/if}

        <button class="bg-gray-200 p-2 rounded-md">
            <Icon icon="lucide:search" />
        </button>
    </div>
</div>
<div class="grid grid-cols-3 border-b-2">
    <div>Home > Forums</div>
    <div></div>
    <div class="">
        <nav class="flex flex-row gap-2 items-center justify-center">
            <button class="px-2 p-3 hover:bg-gray-300 ease-in-out duration-150">Home</button>
            <button class="px-2 p-3 hover:bg-gray-300 ease-in-out duration-150">Home</button>
        </nav>
    </div>
</div>