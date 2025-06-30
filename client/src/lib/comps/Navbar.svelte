<script>
    import SimpleButton from "$lib/comps/UI/SimpleButton.svelte";

    let buttons = [
        {
            name: "Forums",
            link: "/forums"
        },
        {
            
        }
    ]
</script>
<div class="grid grid-cols-3 w-full items-center px-6 py-3 rounded-lg">
    <!-- Left: Nav Links -->
    <nav class="flex justify-center items-center space-x-4">
        <SimpleButton href="/">Home</SimpleButton>
        <!-- Add more if needed -->
    </nav>

    <!-- Center: Logo / Brand -->
    <div class="text-xl font-bold text-gray-900 tracking-tight text-center">
        Chymein
    </div>

    <!-- Right: User Area -->
    <div class="flex justify-center items-center gap-4">
        <a href="/login" class="text-sm text-gray-700 hover:text-gray-900 transition">Login</a>
        <a href="/register" class="bg-gray-800 text-white text-sm px-3 py-1.5 rounded hover:bg-gray-900 transition">
            Sign Up
        </a>
    </div>
</div>
