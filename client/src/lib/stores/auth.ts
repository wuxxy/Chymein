import { writable } from 'svelte/store';
import { axiosInstance } from '$lib/utils/axios';
import type {UserType} from "$lib/utils/types";


const { subscribe, set, update } = writable<UserType | null>(null);

async function fetchUser(): Promise<void> {
    try {
        const { data } = await axiosInstance.get<UserType>('/Auth/me');
        set(data);
    } catch {
        set(null); // Not logged in or session expired
    }
}

export const user = {
    subscribe,
    fetch: fetchUser,
    reset: () => set(null)
};
