<script lang="ts">
    import "../app.css";

    import Layout from "$lib/comps/Layout.svelte";
    import SetupForm from "$lib/comps/SetupForm.svelte";
    import NotConnected from "$lib/comps/NotConnected.svelte";

    let { children, data } = $props();
</script>

{#if data.status.is_setup || !data.status.connected}
    <Layout {children} />
{:else}
    <SetupForm status={data.status} />
{/if}
