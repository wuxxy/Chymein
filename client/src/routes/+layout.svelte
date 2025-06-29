<script lang="ts">

    import SetupForm from "$lib/comps/SetupForm.svelte";
    import "../app.css";
    let { children, data } = $props();
</script>

{#if data.status.is_setup}
    {@render children()}
{:else}
    <SetupForm status={data.status} />
{/if}