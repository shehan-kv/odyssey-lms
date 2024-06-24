<script>
	import CreateSupportTicketDialog from '$lib/components/createSupportTicketDialog.svelte';
	import Header from '$lib/components/header.svelte';
	import MobileSidebar from '$lib/components/mobileSidebar.svelte';
	import ScreenSizeWidget from '$lib/components/screenSizeWidget.svelte';
	import Sidebar from '$lib/components/sidebar.svelte';
	import { Toaster } from '$lib/components/ui/sonner';

	let mobileMenu = false;
	const handleMobileMenu = () => {
		mobileMenu = !mobileMenu;
	};

	let supportTicketDialog = false;
	const handleTicketDialog = () => {
		supportTicketDialog = !supportTicketDialog;
	};
</script>

{#if mobileMenu}
	<MobileSidebar on:mobileMenuToggle={handleMobileMenu} />
{/if}

<CreateSupportTicketDialog open={supportTicketDialog} on:close={handleTicketDialog} />

<div class="flex flex-col gap-y-2 h-screen bg-neutral-100 dark:bg-neutral-900 p-2">
	<Header on:mobileMenuToggle={handleMobileMenu} on:newTicket={handleTicketDialog} />
	<div class="flex grow gap-2 overflow-hidden">
		<Sidebar />

		<slot></slot>
	</div>
</div>

<Toaster />
