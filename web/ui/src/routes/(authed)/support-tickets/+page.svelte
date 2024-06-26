<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import Input from '$lib/components/ui/input/input.svelte';
	import { CircleAlert, CircleCheck, Info, Search, Ticket } from 'lucide-svelte';
	import * as Select from '$lib/components/ui/select';
	import * as Table from '$lib/components/ui/table';
	import * as Pagination from '$lib/components/ui/pagination';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';

	let loading = true;
	let fetchError = false;

	let query = new URLSearchParams($page.url.searchParams.toString());

	let types = [
		{ value: '', label: 'All Types' },
		{ value: 'system', label: 'System' },
		{ value: 'course', label: 'Course' },
		{ value: 'user', label: 'User' }
	];

	let status = [
		{ value: '', label: 'All Status' },
		{ value: 'resolved', label: 'Resolved' },
		{ value: 'unresolved', label: 'Unresolved' },
	];

	let searchTextTemp = query.get('search') || ''; // search text variable to bind to the input
	let searchText = query.get('search') || '';
	let initialType = query.get('type') || '';
	let initialStatus = query.get('status') || '';
	let currentType = {
		value: types.filter((type) => type.value == initialType)[0]?.['value'] || types[0].value,
		label: types.filter((type) => type.value == initialType)[0]?.['label'] || types[0].label
	};
	let currentStatus = {
		value: status.filter((state) => state.value == initialStatus)[0]?.['value'] || status[0].value,
		label: status.filter((state) => state.value == initialStatus)[0]?.['label'] || status[0].label
	};

	let perPage = 30;
	let pageNum = Number(query.get('page')) || 1;

	$: {
		query.set('search', searchText);
		if (searchText == '') query.delete('search');

		query.set('type', currentType.value);
		if (currentType.value == '') query.delete('type');

		query.set('status', currentStatus.value);
		if (currentStatus.value == '') query.delete('status');

		query.set('page', pageNum.toString());
		if (pageNum == 1) query.delete('page');

		goto(`?${query.toString()}`);
		fetchData();
	}

	/**
	 * @type {{
	 * 	totalCount: number,
	 * 	tickets: {
	 *  id: number,
	 * 	createdAt: number,
	 * 	type: string,
	 * 	subject: string,
	 * 	user: string,
	 * 	status: string,
	 * }[]}}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/support-ticket?${query.toString()}`)
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

<svelte:head>
	<title>Support Tickets</title>
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
			<p>Support Tickets</p>
		</div>
		<div class="flex items-center gap-5">
			<div class="relative z-0">
				<Input
					type="text"
					placeholder="search"
					class="h-9 text-xs rounded w-60 pl-10"
					bind:value={searchTextTemp}
					on:keydown={(e) => {
						if (e.key == 'Enter') {
							searchText = searchTextTemp;
						}
					}}
				/>
				<Search class="absolute left-2 top-[50%] -translate-y-[50%] text-neutral-500" size="20" />
			</div>

			<Select.Root bind:selected={currentType}>
				<Select.Trigger class="w-[180px] text-xs rounded h-9">
					<Select.Value placeholder="All Types" />
				</Select.Trigger>
				<Select.Content>
					<Select.Item class="pr-2" value="">All Types</Select.Item>
					<Select.Item class="pr-2" value="system">System</Select.Item>
					<Select.Item class="pr-2" value="course">Course</Select.Item>
					<Select.Item class="pr-2" value="user">User</Select.Item>
				</Select.Content>
			</Select.Root>

			<Select.Root bind:selected={currentStatus}>
				<Select.Trigger class="w-[180px] text-xs rounded h-9">
					<Select.Value placeholder="All Status" />
				</Select.Trigger>
				<Select.Content>
					<Select.Item class="pr-2" value="">All Status</Select.Item>
					<Select.Item class="pr-2" value="resolved">Resolved</Select.Item>
					<Select.Item class="pr-2" value="unresolved">Unresolved</Select.Item>
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	{#if !fetchError && (!data || data.tickets.length == 0)}
		<div class="flex grow justify-center items-center">
			<p class="text-sm text-neutral-400 dark:text-neutral-600">No Tickets Found</p>
		</div>
	{/if}

	{#if data && data.tickets.length > 0}
		<Table.Root class="text-xs mt-10">
			<Table.Header>
				<Table.Row class="hover:bg-transparent">
					<Table.Head class="font-semibold">Timestamp</Table.Head>
					<Table.Head class="font-semibold">Type</Table.Head>
					<Table.Head class="font-semibold">Subject</Table.Head>
					<Table.Head class="font-semibold">User</Table.Head>
					<Table.Head class="font-semibold text-right">Severity</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each data.tickets as ticket}
					<Table.Row on:click={()=> goto(`/support-tickets/${ticket.id}`)}>
						<Table.Cell>
							{ticket.createdAt ? new Date(ticket.createdAt).toLocaleString() : 'Not Found'}
						</Table.Cell>
						<Table.Cell class="capitalize">{ticket.type}</Table.Cell>
						<Table.Cell>{ticket.subject}</Table.Cell>
						<Table.Cell>{ticket.user}</Table.Cell>
						<Table.Cell>
							{#if ticket.status == 'resolved'}
								<span class="flex justify-end items-center gap-2">
									Resolved <span><CircleCheck size={16} /></span>
								</span>
							{:else if ticket.status == 'unresolved'}
								<span class="flex justify-end items-center gap-2">
									Unresolved <span><CircleAlert size={16} /></span>
								</span>
							{/if}
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	{/if}

	{#if data && data.tickets.length > 0}
		<Pagination.Root
			count={data.totalCount}
			{perPage}
			page={pageNum}
			let:pages
			let:currentPage
			class="w-fit mr-0 mt-10"
		>
			<Pagination.Content>
				<Pagination.Item>
					<Pagination.PrevButton
						on:click={() => {
							pageNum = currentPage - 1;
						}}
					/>
				</Pagination.Item>
				{#each pages as page (page.key)}
					{#if page.type === 'ellipsis'}
						<Pagination.Item>
							<Pagination.Ellipsis />
						</Pagination.Item>
					{:else}
						<Pagination.Item isVisible={currentPage == page.value}>
							<Pagination.Link
								{page}
								isActive={currentPage == page.value}
								on:click={() => {
									pageNum = page.value;
								}}
							>
								{page.value}
							</Pagination.Link>
						</Pagination.Item>
					{/if}
				{/each}
				<Pagination.Item>
					<Pagination.NextButton
						on:click={(e) => {
							pageNum = currentPage + 1;
						}}
					/>
				</Pagination.Item>
			</Pagination.Content>
		</Pagination.Root>
	{/if}
</div>
