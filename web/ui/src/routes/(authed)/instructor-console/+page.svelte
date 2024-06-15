<script>
	import AssignedCourseCard from '$lib/components/assignedCourseCard.svelte';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import LecternIcon from '$lib/components/lecternIcon.svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	img: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
	 * 	students: number
	 * }[]}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/course/assigned')
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
	<title>Instructor Console</title>
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

	<div class="flex gap-2 text-sm font-semibold">
		<LecternIcon size={20} />
		<p>Assigned Courses</p>
	</div>

	{#if !fetchError && (!data || data.length == 0)}
		<div class="flex grow justify-center items-center">
			<p class="text-sm text-neutral-400 dark:text-neutral-600">No Assigned Courses Found</p>
		</div>
	{/if}

	{#if data && data.length > 0}
		<div class="flex flex-wrap gap-x-10 gap-y-16 mt-10">
			{#each data as course}
				<AssignedCourseCard {course} />
			{/each}
		</div>
	{/if}
</div>
