<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import Input from '$lib/components/ui/input/input.svelte';
	import { Search, SendHorizonal } from 'lucide-svelte';
	import * as Pagination from '$lib/components/ui/pagination';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';

	let loading = true;
	let fetchError = false;

	let query = new URLSearchParams($page.url.searchParams.toString());

	let searchTextTemp = query.get('search') || ''; // search text variable to bind to the input
	let searchText = query.get('search') || '';

	let perPage = 30;
	let pageNum = Number(query.get('page')) || 1;

	$: {
		query.set('search', searchText);
		if (searchText == '') query.delete('search');

		query.set('page', pageNum.toString());
		if (pageNum == 1) query.delete('page');

		goto(`?${query.toString()}`);
		fetchData();
	}

	/**
	 * @type {{
	 * 	totalCount: number,
	 * 	notifications: {
	 * 	timestamp: number,
	 * 	subject: string,
	 * 	description: string,
	 * }[]}}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/notification/sent?${query.toString()}`)
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
	<title>Sent Notifications</title>
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
			<SendHorizonal size={20} />
			<p>Sent Notifications</p>
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
		</div>
	</div>

	{#if !fetchError && (!data || data.notifications.length == 0)}
		<div class="flex grow justify-center items-center">
			<p class="text-sm text-neutral-400 dark:text-neutral-600">No Notifications Found</p>
		</div>
	{/if}

	{#if data && data.notifications.length > 0}
		<div class="text-sm mt-10">
			{#each data.notifications as notification}
				<div class="px-2 py-6 border-b space-y-2 hover:bg-muted/50">
					<div class="flex justify-between">
						<p class="font-semibold">{notification.subject}</p>
						<p>{new Date(notification.timestamp).toLocaleString()}</p>
					</div>
					<p class="text-neutral-700 dark:text-neutral-300 max-w-2xl">
						{notification.description}
					</p>
				</div>
			{/each}
		</div>
	{/if}

	{#if data && data.notifications.length > 0}
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
