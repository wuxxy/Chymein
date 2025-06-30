<script>
    let {href, active, children} = $props()
</script>

<button class={`${active ? "bg-gray-300" : ""} p-2 rounded-md hover:bg-gray-200 transition-`}>
    {@render children()}
</button>