<script lang="ts">
    import {axiosInstance} from "$lib/utils/axios";

    let { users }: { users: any[] } = $props();

    const formatDate = (date: string) =>
        new Date(date).toLocaleString(undefined, {
            dateStyle: 'short',
            timeStyle: 'short'
        });
    $effect(async () => {
        try {
            const res = await axiosInstance.get('/Admin/users');
            users = res.data;
        } catch (err) {
            error = err?.response?.data?.message ?? "Couldn't load users";
        }
    });
</script>

<table class="w-full border-separate border-spacing-y-2 text-sm">
    <thead class="bg-gray-100 text-left">
    <tr>
        <th class="px-4 py-2">Username</th>
        <th class="px-4 py-2">Email</th>
        <th class="px-4 py-2">Admin</th>
        <th class="px-4 py-2">Verified</th>
        <th class="px-4 py-2">Last Active</th>
        <th class="px-4 py-2">Sessions</th>
    </tr>
    </thead>
    <tbody>
    {#each users as user (user.ID)}
        <tr class="bg-white hover:bg-gray-50 rounded shadow-sm">
            <td class="px-4 py-2 font-medium">{user.Username}</td>
            <td class="px-4 py-2">{user.Email}</td>
            <td class="px-4 py-2">{user.Admin ? "✅" : "❌"}</td>
            <td class="px-4 py-2">{user.IsVerified ? "✅" : "❌"}</td>
            <td class="px-4 py-2">{formatDate(user.LastActive)}</td>
            <td class="px-4 py-2">
                <details>
                    <summary class="cursor-pointer text-blue-500 underline">View ({user.Sessions.length})</summary>
                    <ul class="mt-2 text-xs pl-4 space-y-1 list-disc">
                        {#each user.Sessions as s}
                            <li>
                                <span class="font-semibold">Agent:</span> {s.UserAgent || "N/A"}<br />
                                <span class="font-semibold">Last Used:</span> {formatDate(s.LastUsedAt)}<br />
                                <span class="font-semibold">IP:</span> {s.IPAddress || "unknown"}
                            </li>
                        {/each}
                    </ul>
                </details>
            </td>
        </tr>
    {/each}
    </tbody>
</table>
