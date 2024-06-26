<script>
	import { page } from '$app/stores';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import ReplySupportTicket from '$lib/components/replySupportTicket.svelte';
	import SupportTicketMarkResolved from '$lib/components/supportTicketMarkResolved.svelte';
	import { CalendarClock, CheckCheck, Ticket, User } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	ticket: {
	 * 		id: number,
	 * 		subject: string,
	 * 		user: string,
	 * 		description:string,
	 * 		type: string,
	 * 		createdAt: string,
	 * 		closedAt: string,
	 * 		status: string
	 * 	},
	 * 	messages: {
	 *  	id: number,
	 * 		createdAt: string,
	 * 		user: string,
	 * 		content: string,
	 * 	}[]
	 * }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/support-ticket/${$page.params.id}`)
			.then((response) => {
				if (!response.ok) {
					fetchError = true;
					loading = false;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				data = parsedData;
				loading = false;
			})
			.catch(() => {
				fetchError = true;
				loading = false;
			});
	};

	onMount(() => {
		fetchData();
	});
</script>

<svelte:head>
	<title>Support Ticket Details</title>
</svelte:head>

<div
	class="relative flex flex-col py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll grow"
>
	{#if loading}
		<CardLoader />
	{/if}
	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="flex justify-between items-center">
		<div class="flex gap-2 text-sm font-semibold">
			<Ticket size={20} />
			<p>Support Ticket Details</p>
		</div>
	</div>

	{#if data}
		<div class="mt-10 max-w-4xl">
			<p class="text-lg font-semibold">{data.ticket.subject}</p>
			<div class="flex justify-between">
				<div class="flex gap-6">
					<div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 text-xs mt-2">
						<CalendarClock size={16} />
						{data.ticket.createdAt ? new Date(data.ticket.createdAt).toLocaleString() : 'Not Found'}
					</div>
					<div class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 text-xs mt-2">
						<User size={16} />
						{data.ticket.user}
					</div>
				</div>
				{#if data && data.ticket.status != 'resolved'}
					<SupportTicketMarkResolved id={$page.params.id} on:success={fetchData} />
				{/if}
			</div>
			<p class="mt-4">
				{data.ticket.description}
			</p>
		</div>
	{/if}

	{#if data && data.messages.length > 0}
		<div class="mt-20 max-w-4xl space-y-10 pl-10">
			{#each data.messages as message}
				<div>
					<div class="flex gap-6">
						<div
							class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 text-xs mt-2"
						>
							<CalendarClock size={16} />
							{message.createdAt ? new Date(message.createdAt).toLocaleString() : 'Not Found'}
						</div>
						<div
							class="flex items-center gap-2 text-neutral-600 dark:text-neutral-400 text-xs mt-2"
						>
							<User size={16} />
							{message.user}
						</div>
					</div>
					<div class="mt-2">
						<p>{message.content}</p>
					</div>
				</div>
			{/each}
		</div>
	{/if}

	{#if data && data.ticket.status == 'resolved'}
		<div
			class="flex gap-2 items-center mt-20 w-fit p-4 bg-emerald-100 text-emerald-950 dark:bg-emerald-950 dark:text-emerald-100 font-semibold text-sm rounded"
		>
			<CheckCheck size={20} />
			<p class="">
				Marked As Resolved: {data.ticket.closedAt
					? new Date(data.ticket.closedAt).toLocaleString()
					: 'Date Not Found'}
			</p>
		</div>
	{/if}

	{#if data && data.ticket.status != 'resolved'}
		<ReplySupportTicket ticketId={$page.params.id} on:success={fetchData} />
	{/if}
</div>
