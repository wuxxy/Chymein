<script lang="ts">
    import "../app.css";

    import Layout from "$lib/components/layout/Layout.svelte";
    import SetupForm from "$lib/components/forms/SetupForm.svelte";

    let { children, data } = $props();
    import { onMount } from 'svelte';
    import {user} from "$lib/stores/auth";

    onMount(() => {
        user.fetch(); // will populate the store
    });
</script>

{#if data.status.is_setup}
    {@render children()}
{:else}
    <SetupForm status={data.status} />
{/if}

<style>
    @import url('https://fonts.googleapis.com/css2?family=Noto+Sans:ital,wght@0,100..900;1,100..900&display=swap');
</style>