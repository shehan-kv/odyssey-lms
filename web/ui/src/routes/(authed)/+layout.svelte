<script>
	import { goto } from '$app/navigation';
	import FullScreenLoader from '$lib/components/fullScreenLoader.svelte';
	import Header from '$lib/components/header.svelte';
	import MobileSidebar from '$lib/components/mobileSidebar.svelte';
	import ScreenSizeWidget from '$lib/components/screenSizeWidget.svelte';
	import Sidebar from '$lib/components/sidebar.svelte';
	import { onMount } from 'svelte';

	let mobileMenu = false;

	const handleMobileMenu = () => {
		mobileMenu = !mobileMenu;
	};

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

{#if loading}
	<FullScreenLoader />
{/if}

{#if mobileMenu}
	<MobileSidebar on:mobileMenuToggle={handleMobileMenu} />
{/if}

<div class="flex flex-col gap-y-2 h-screen bg-neutral-100 dark:bg-neutral-900 p-2">
	<Header on:mobileMenuToggle={handleMobileMenu} />
	<div class="flex grow gap-2 overflow-hidden">
		<Sidebar />

		<slot></slot>
	</div>
</div>
