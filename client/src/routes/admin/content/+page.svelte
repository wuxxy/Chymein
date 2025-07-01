<script lang="ts">
    import { axiosInstance } from "$lib/utils/axios";

    let containers: any[] = $state([]);

    const formatDate = (date: string) =>
        new Date(date).toLocaleString(undefined, {
            dateStyle: 'short',
            timeStyle: 'short'
        });

    $effect(async () => {
        try {
            const res = await axiosInstance.get('/Admin/containers');
            containers = res.data;
        } catch (err) {
            error = err?.response?.data?.message ?? "Couldn't load containers";
        }
    });
</script>

<table class="w-full text-sm border-separate border-spacing-y-1">
    <thead class="bg-gray-50 text-left">
    <tr>
        <th class="px-4 py-2">Name</th>
        <th class="px-4 py-2">Description</th>
        <th class="px-4 py-2">Type</th>
        <th class="px-4 py-2">Sort</th>
        <th class="px-4 py-2">Created</th>
        <th class="px-4 py-2">Updated</th>
    </tr>
    </thead>
    <tbody>
    {#each containers as container (container.ID)}
        <tr class="bg-white hover:bg-gray-50 rounded">
            <td class="px-4 py-2 font-medium">{container.Name}</td>
            <td class="px-4 py-2 text-gray-700">{container.Description || "—"}</td>
            <td class="px-4 py-2">{container.Type}</td>
            <td class="px-4 py-2">{container.SortOrder}</td>
            <td class="px-4 py-2 text-gray-500">{formatDate(container.CreatedAt)}</td>
            <td class="px-4 py-2 text-gray-500">{formatDate(container.UpdatedAt)}</td>
        </tr>
    {/each}
    {#if containers.length === 0}
        <p class="text-center text-gray-400 mt-4">No containers found.</p>
    {/if}
    </tbody>


</table>
