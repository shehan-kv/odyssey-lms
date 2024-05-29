<script>
	import { Moon, Sun } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';

	let className = '';
	export { className as class };

	let isDark = false;

	if (document.documentElement.classList.contains('dark') || localStorage.theme == 'dark') {
		isDark = true;
	}

	const toggleTheme = () => {
		if (isDark) {
			isDark = false;
			localStorage.theme = 'light';
			document.documentElement.classList.remove('dark');
		} else {
			isDark = true;
			localStorage.theme = 'dark';
			document.documentElement.classList.add('dark');
		}
	};
</script>

<Button on:click={toggleTheme} class={className}>
	{#if isDark}
		<Sun />
	{:else}
		<Moon />
	{/if}
</Button>
