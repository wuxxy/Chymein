<script lang="ts">
    import {onDestroy, onMount} from 'svelte';


    import { axiosInstance } from '$lib/utils/axios';
    import {redirect} from "@sveltejs/kit";
    import {user} from "$lib/stores/auth";
    import type {Unsubscriber} from "svelte/store";
    import {goto} from "$app/navigation";

    let username = $state("");
    let password = $state("");
    let errorMsg = $state("");
    let loading = $state(false);
    let unsubscribe: Unsubscriber;

    onMount(() => {
        unsubscribe = user.subscribe((user) => {
            if (user) {
                // Already logged in, redirect
                goto('/');
            }
        });
    });

    // Clean up on destroy
    onDestroy(() => {
        unsubscribe?.();
    });
    async function login() {
        loading = true;
        errorMsg = '';

        try {
            const res = await axiosInstance.post('/Auth/login', { username, password });

            if (res.status === 200) {
                await user.fetch()
                await goto('/');
            }
        } catch (err) {
            if (err.response && err.response.data?.error) {
                errorMsg = err.response.data.error;
            } else {
                errorMsg = 'Unexpected server error occurred.';
            }
        } finally {
            loading = false;
        }

    }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-100 px-4">
    <div class="w-full max-w-md bg-white p-8 rounded-xl shadow-xl space-y-6 border border-gray-200">
        <h1 class="text-2xl font-bold text-gray-900 text-center">Welcome back</h1>
        <p class="text-sm text-gray-500 text-center">Enter your credentials to continue</p>

        {#if errorMsg}
            <div class="text-sm text-red-600 bg-red-100 p-3 rounded-md">{errorMsg}</div>
        {/if}

        <div class="space-y-4">
            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1" for="username">Username</label>
                <input
                        id="username"
                        type="text"
                        bind:value={username}
                        placeholder="johndoe"
                        class="w-full px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-800"
                />
            </div>

            <div>
                <label class="block text-sm font-medium text-gray-700 mb-1" for="password">Password</label>
                <input
                        id="password"
                        type="password"
                        bind:value={password}
                        placeholder="••••••••"
                        class="w-full px-4 py-2 rounded-md border border-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-800"
                />
            </div>

            <button
                    onclick={login}
                    class="w-full bg-gray-900 text-white py-2 rounded-md hover:bg-gray-800 transition text-sm font-medium"
                    disabled={loading}
            >
                {loading ? 'Logging in...' : 'Login'}
            </button>
        </div>
    </div>
</div>
