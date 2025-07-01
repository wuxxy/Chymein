<script>
    import Navbar from "$lib/components/layout/Navbar.svelte";

    let {children} = $props()
</script>

<div class="flex flex-col bg-gray-200">
    <main class="min-h-screen w-full md:w-[80%] mx-auto flex flex-col ">
        <Navbar />
        {@render children()}
    </main>

</div>