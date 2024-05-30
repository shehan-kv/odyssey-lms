<script>
	import { goto } from '$app/navigation';
	import FullScreenLoader from '$lib/components/ui/fullScreenLoader.svelte';
	import { onMount } from 'svelte';

	let loading = true;

	onMount(() => {
		fetch('/api/auth/is-signed-in')
			.then((response) => {
				if (response.status != 200) {
					goto('/sign-in');
				} else {
					loading = false;
				}
			})
			.catch(() => {
				goto('/sign-in');
			});
	});
</script>

<slot />

{#if loading}
	<FullScreenLoader />
{/if}
