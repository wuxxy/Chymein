<script>
    import Navbar from "$lib/comps/Navbar.svelte";

    let {children} = $props()
</script>

<div class="flex flex-col ">
    <main class="w-full md:w-[80%] mx-auto">
        <Navbar />
        <div class="bg-gray-200 rounded-md p-2">
            {@render children()}
        </div>
    </main>

</div>