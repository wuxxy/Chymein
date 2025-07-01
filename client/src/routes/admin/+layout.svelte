<script lang="ts">
    import { user } from "$lib/stores/auth";
    let { children } = $props();
</script>

<div class="flex flex-col min-h-screen bg-gray-50 text-gray-900 font-sans">
    <!-- Top Nav -->
    <header class="bg-white shadow-sm border-b border-gray-200 px-6 py-4">
        <div class="flex items-center justify-between">
            <div class="text-xl font-semibold tracking-tight">
                Admin Panel
            </div>
            <div class="text-sm text-gray-600">
                Logged in as: <span class="font-medium text-gray-800">{$user?.Username}</span>
            </div>
        </div>
        <nav class="mt-3">
            <ul class="flex gap-4 text-sm">
                <li>
                    <a href="/" class="text-gray-600 hover:text-black transition-colors">Home</a>
                </li>
                <li>
                    <a href="/admin/users" class="text-gray-600 hover:text-black transition-colors">Users</a>
                </li>
                <li>
                    <a href="/admin/styles" class="text-gray-600 hover:text-black transition-colors">Styles</a>
                </li>
            </ul>
        </nav>
    </header>

    <!-- Page Content -->
    <main class="flex-1 p-6">
        {@render children()}
    </main>
</div>
