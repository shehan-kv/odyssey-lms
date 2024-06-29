import { writable } from 'svelte/store';

export const isDark = writable(false);
export const userRole = writable("student")
