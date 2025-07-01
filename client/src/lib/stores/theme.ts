import { derived, writable } from 'svelte/store';
import { browser } from '$app/environment';
import { user } from './auth';
import { axiosInstance } from '$lib/utils/axios';

type ThemeStyles = Record<string, string>;

const themeName = writable<string>('default');
const styleMap = writable<ThemeStyles>({});

// Read from localStorage (only in browser)
if (browser) {
    const stored = localStorage.getItem('theme');
    if (stored) themeName.set(stored);
}

// Update from user store if available
const themeWatcher = derived(user, ($user) => {
    if ($user?.Theme) {
        themeName.set($user.Theme);
        if (browser) localStorage.setItem('theme', $user.Theme);
    }
});

// Fetch and store style config for the theme
async function fetchThemeStyles(name: string) {
    try {
        const res = await axiosInstance.get(`/api/theme/${name}`);
        styleMap.set(res.data);
    } catch (err) {
        console.error(`Failed to fetch theme: ${name}`, err);
        styleMap.set({});
    }
}

// Automatically watch for theme changes and fetch styles
themeName.subscribe((name) => {
    if (browser) localStorage.setItem('theme', name);

    void fetchThemeStyles(name);
});

export const theme = {
    name: themeName,
    styles: styleMap,
    set: (name: string) => themeName.set(name)
};
