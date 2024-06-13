<script>
	import { CircleAlert, Info } from 'lucide-svelte';
	import * as Table from '$lib/components/ui/table';
	import CardLoader from './cardLoader.svelte';
	import CardErrorRetry from './cardErrorRetry.svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	timestamp: number,
	 *  description: string,
	 *  severity: string,
	 * }[]}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/event/recent')
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

<div class="relative py-6 px-4 grow bg-white dark:bg-neutral-950 rounded">
	{#if loading}
		<CardLoader />
	{/if}

	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="space-y-4">
		<div class="flex gap-2 text-sm font-semibold">
			<Info size={20} />
			<p>Recent Events</p>
		</div>

		{#if !fetchError && (!data || data.length == 0)}
			<p class="text-sm text-center text-neutral-400 dark:text-neutral-600">
				No Recent Events Found
			</p>
		{/if}

		{#if data && data.length > 0}
			<div class="max-h-64 overflow-scroll">
				<Table.Root class="text-xs">
					<Table.Header>
						<Table.Row class="hover:bg-transparent">
							<Table.Head class="font-semibold">Timestamp</Table.Head>
							<Table.Head class="font-semibold">Description</Table.Head>
							<Table.Head class="font-semibold text-right">Severity</Table.Head>
						</Table.Row>
					</Table.Header>
					<Table.Body>
						{#each data as event}
							<Table.Row>
								<Table.Cell>{new Date(event.timestamp).toLocaleString()}</Table.Cell>
								<Table.Cell>{event.description}</Table.Cell>
								<Table.Cell>
									<span class="flex justify-end items-center gap-2">
										{#if event.severity == 'critical'}
											Critical <span class="text-red-600"><CircleAlert size={16} /></span>
										{:else if event.severity == 'info'}
											Information <span><Info size={16} /></span>
										{/if}
									</span>
								</Table.Cell>
							</Table.Row>
						{/each}
					</Table.Body>
				</Table.Root>
			</div>
		{/if}
	</div>
</div>
