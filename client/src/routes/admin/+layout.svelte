<script lang="ts">
    import { user } from "$lib/stores/auth";
    import {Button, Toolbar} from "bits-ui";
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
        <div class="flex flex-row">
            <Toolbar.Root>
                <Toolbar.Group class="flex gap-1 px-2 py-1">
                    <Toolbar.GroupItem>
                        <Button.Root class="data-[active=true]:bg-gray-100" data-active="true">Home</Button.Root>
                    </Toolbar.GroupItem>
                    <Toolbar.GroupItem>
                        <Button.Root href="/admin/users">Users</Button.Root>
                    </Toolbar.GroupItem>
                    <Toolbar.GroupItem>
                        <Button.Root href="/admin/content">Content</Button.Root>
                    </Toolbar.GroupItem>
                    <Toolbar.GroupItem>
                        <Button.Root>Styling</Button.Root>
                    </Toolbar.GroupItem>
                    <Toolbar.GroupItem>
                        <Button.Root>Plugins</Button.Root>
                    </Toolbar.GroupItem>
                </Toolbar.Group>
            </Toolbar.Root>

        </div>
    </header>

    <!-- Page Content -->
    <main class="flex-1 p-6">
        {@render children()}
    </main>
</div>
