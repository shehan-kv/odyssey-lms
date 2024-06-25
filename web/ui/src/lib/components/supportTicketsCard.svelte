<script>
	import * as Table from '$lib/components/ui/table';
	import { CircleAlert, CircleCheck, Ticket } from 'lucide-svelte';
	import CardLoader from './cardLoader.svelte';
	import CardErrorRetry from './cardErrorRetry.svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * totalCount: number,
	 * tickets:{
	 * 	createdAt: number,
	 *  type: string,
	 *  subject: string,
	 *  user: string,
	 *  status: string,
	 * }[]}}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/support-ticket?limit=6')
			.then((response) => {
				if (response.status != 200) {
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

<div class="relative py-6 px-4 bg-white dark:bg-neutral-950 rounded min-h-40">
	{#if loading}
		<CardLoader />
	{/if}

	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="space-y-4">
		<div class="flex gap-2 text-sm font-semibold">
			<Ticket size={20} />
			<p>Support Tickets</p>
		</div>

		{#if !fetchError && (!data || data.tickets.length == 0)}
			<p class="text-sm text-center text-neutral-400 dark:text-neutral-600">
				No Recent Support Tickets Found
			</p>
		{/if}

		{#if data && data.tickets.length > 0}
			<Table.Root class="text-xs">
				<Table.Header>
					<Table.Row class="hover:bg-transparent">
						<Table.Head class="font-semibold">Timestamp</Table.Head>
						<Table.Head class="font-semibold">Category</Table.Head>
						<Table.Head class="font-semibold">Subject</Table.Head>
						<Table.Head class="font-semibold">Submitted By</Table.Head>
						<Table.Head class="font-semibold text-right">Status</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each data.tickets as ticket}
						<Table.Row>
							<Table.Cell>
								{ticket.createdAt ? new Date(ticket.createdAt).toLocaleString() : 'Not Found'}
							</Table.Cell>
							<Table.Cell class="capitalize">{ticket.type}</Table.Cell>
							<Table.Cell>{ticket.subject}</Table.Cell>
							<Table.Cell>{ticket.user}</Table.Cell>
							<Table.Cell>
								<span class="flex justify-end items-center gap-2">
									{#if ticket.status == 'unresolved'}
										Unresolved <span><CircleAlert size={16} /></span>
									{:else if ticket.status == 'resolved'}
										Resolved <span><CircleCheck size={16} /></span>
									{/if}
								</span>
							</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		{/if}
	</div>
</div>
