<script>
	import { Moon, Sun } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';
	import { isDark } from '../../stores/store';

	let className = '';
	export { className as class };

	if (document.documentElement.classList.contains('dark') || localStorage.theme == 'dark') {
		isDark.set(true);
	}

	const toggleTheme = () => {
		if ($isDark) {
			isDark.set(false);
			localStorage.theme = 'light';
			document.documentElement.classList.remove('dark');
		} else {
			isDark.set(true);
			localStorage.theme = 'dark';
			document.documentElement.classList.add('dark');
		}
	};
</script>

<Button on:click={toggleTheme} class={className}>
	{#if $isDark}
		<Sun />
	{:else}
		<Moon />
	{/if}
</Button>
