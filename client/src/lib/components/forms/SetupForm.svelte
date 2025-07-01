<script>
    import { axiosInstance } from '$lib/utils/axios.ts';

    let props = $props();
    const status = props.status;

    let username = $state('');
    let email = $state('');
    let password = $state('');
    let confirmPassword = $state('');
    let submitting = $state(false);
    let errorMsg = $state('');
    let successMsg = $state('');

    async function createSuperadmin() {
        errorMsg = '';
        successMsg = '';
        submitting = true;

        if (password !== confirmPassword) {
            errorMsg = 'Passwords do not match.';
            submitting = false;
            return;
        }

        try {
            const res = await axiosInstance.post('/create_admin', {
                username,
                email,
                password
            });

            if (res.status === 201 || res.status === 200) {
                successMsg = 'Superadmin created successfully. Reloading...';
                setTimeout(() => location.reload(), 1500);
            } else {
                errorMsg = res.data?.error || 'Failed to create superadmin.';
            }
        } catch (err) {
            errorMsg = err.response?.data?.error || 'Server error occurred.';
        } finally {
            submitting = false;
        }
    }

    function refresh() {
        location.reload();
    }
</script>

<div class="min-h-screen bg-gray-200 flex items-center justify-center px-4">
    <div class="w-full max-w-xl bg-white rounded-xl p-8 shadow-xl space-y-8">
        <!-- Header -->
        <div class="space-y-2">
            <h1 class="text-3xl font-bold text-gray-900">Chymein Setup</h1>

            {#if status.database}
                <div class="flex items-center gap-2 text-green-700 bg-green-100 border border-green-300 rounded px-3 py-2 text-sm font-medium">
                    <div class="w-2 h-2 bg-green-600 rounded-full animate-pulse"></div>
                    <span>Database connected</span>
                </div>
            {/if}
        </div>

        <!-- Main Body -->
        {#if !status.database}
            <div class="space-y-4">
                <p class="text-gray-700 text-base leading-relaxed">
                    To continue setup, connect your database:
                </p>
                <ul class="list-disc pl-5 text-gray-700 space-y-1 text-sm">
                    <li>Edit <code class="bg-gray-100 px-1 py-0.5 rounded text-sm">config.json</code> in your server directory</li>
                    <li>Update the database credentials</li>
                    <li>Restart your server</li>
                    <li>Then refresh this page</li>
                </ul>
                <button
                        class="bg-gray-900 text-white px-4 py-2 rounded hover:bg-gray-800 transition w-fit"
                        on:click={refresh}
                >
                    Refresh Page
                </button>
            </div>
        {:else if !status.is_setup}
            <div class="space-y-4">
                <p class="text-gray-700 text-base">
                    Create a <span class="font-semibold text-gray-900">Superadmin</span> account to finish setup.
                </p>

                <!-- Feedback Message -->
                {#if errorMsg}
                    <div class="bg-red-100 border border-red-300 text-red-800 text-sm rounded px-4 py-2">
                        {errorMsg}
                    </div>
                {/if}

                {#if successMsg}
                    <div class="bg-green-100 border border-green-300 text-green-800 text-sm rounded px-4 py-2">
                        {successMsg}
                    </div>
                {/if}

                <!-- Form -->
                <div class="space-y-3 flex flex-col select-none">
                    <input
                            class="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-gray-800 disabled:bg-gray-100"
                            placeholder="Email"
                            type="email"
                            bind:value={email}
                            disabled={submitting}
                    />
                    <input
                            class="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-gray-800 disabled:bg-gray-100"
                            placeholder="Username"
                            bind:value={username}
                            disabled={submitting}
                    />
                    <input
                            class="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-gray-800 disabled:bg-gray-100"
                            placeholder="Password"
                            type="password"
                            bind:value={password}
                            disabled={submitting}
                    />
                    <input
                            class="w-full border border-gray-300 px-4 py-2 rounded focus:outline-none focus:ring-2 focus:ring-gray-800 disabled:bg-gray-100"
                            placeholder="Confirm Password"
                            type="password"
                            bind:value={confirmPassword}
                            disabled={submitting}
                    />
                    <hr class="text-gray-300" />
                    <button
                            class="w-full bg-gray-900 text-white py-2 rounded hover:bg-gray-800 transition disabled:opacity-50"
                            on:click={createSuperadmin}
                            disabled={submitting}
                    >
                        {submitting ? 'Creating...' : 'Create Superadmin'}
                    </button>
                </div>
            </div>
        {:else}
            <div class="bg-green-100 border border-green-300 rounded px-4 py-3 text-green-800 text-sm font-medium">
                Setup complete. You may now log in and start using the platform.
            </div>
        {/if}
    </div>
</div>
