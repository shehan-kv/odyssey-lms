<script>
	import { page } from '$app/stores';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import { LibrarySquare, Search } from 'lucide-svelte';
	import * as Select from '$lib/components/ui/select';
	import { onMount } from 'svelte';
	import LibraryCourseCard from '$lib/components/libraryCourseCard.svelte';
	import { goto } from '$app/navigation';
	import * as Pagination from '$lib/components/ui/pagination';

	let loading = true;
	let fetchError = false;

	let query = new URLSearchParams($page.url.searchParams.toString());

	/**
	 * @type {{value: string, label: string}[]}
	 */
	let categories = [{ value: '', label: 'All Categories' }];

	let searchTextTemp = query.get('search') || ''; // search text variable to bind to the input
	let searchText = query.get('search') || '';
	let initialCategory = query.get('category') || '';
	let currentCategory = {
		value:
			categories.filter((category) => category.value == initialCategory)[0]?.['value'] ||
			categories[0].value,
		label:
			categories.filter((category) => category.value == initialCategory)[0]?.['label'] ||
			categories[0].label
	};

	let perPage = 30;
	let pageNum = Number(query.get('page')) || 1;

	$: {
		query.set('search', searchText);
		if (searchText == '') query.delete('search');

		query.set('category', currentCategory.value);
		if (currentCategory.value == '') query.delete('category');

		query.set('page', pageNum.toString());
		if (pageNum == 1) query.delete('page');

		goto(`?${query.toString()}`);
		fetchData();
	}

	const fetchCategories = () => {
		fetch('/api/course/category')
			.then((response) => {
				if (response.status != 200) {
					return;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				categories = [...categories, ...parsedData];
			});
	};

	/**
	 * @type {{
	 *	totalCount:number,
	 *  courses: {
	 * 	img: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
	 *  category: string,
	 * 	enrolled: number
	 * }[]}}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/course?${query.toString()}`)
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
		fetchCategories();
		fetchData();
	});
</script>

<svelte:head>
	<title>Course Library</title>
</svelte:head>

<div
	class="relative flex flex-col py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll w-full"
>
	{#if loading}
		<CardLoader />
	{/if}
	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="flex justify-between items-center">
		<div class="flex gap-2 text-sm font-semibold">
			<LibrarySquare size={20} />
			<p>Course Library</p>
		</div>
		<div class="flex gap-5">
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

			<Select.Root bind:selected={currentCategory}>
				<Select.Trigger class="w-[180px] text-xs rounded h-9">
					<Select.Value placeholder="All Categories" />
				</Select.Trigger>
				<Select.Content>
					{#each categories as category (category)}
						<Select.Item class="pr-2" value={category.value}>{category.label}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	{#if !fetchError && (!data || data.courses.length == 0)}
		<div class="flex grow justify-center items-center">
			<p class="text-sm text-neutral-400 dark:text-neutral-600">No Courses Found</p>
		</div>
	{/if}

	{#if data && data.courses.length > 0}
		<div class="flex flex-wrap gap-x-10 gap-y-16 mt-10">
			{#each data.courses as course}
				<LibraryCourseCard {course} />
			{/each}
		</div>
	{/if}

	{#if data && data.courses.length > 0}
		<Pagination.Root
			count={data.totalCount}
			{perPage}
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
						on:click={() => {
							pageNum = currentPage + 1;
						}}
					/>
				</Pagination.Item>
			</Pagination.Content>
		</Pagination.Root>
	{/if}
</div>
