<script>
	import * as Table from '$lib/components/ui/table';
	import { CircleAlert, CircleCheck, GraduationCap, List, Plus } from 'lucide-svelte';
	import CardLoader from './cardLoader.svelte';
	import CardErrorRetry from './cardErrorRetry.svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	totalCourses: number,
	 *  courses:{
	 * 		timestamp: number,
	 * 		code: string,
	 * 		name: string,
	 * 		category: string,
	 * 		enrolled: number,
	 * 		instructor: string[],
	 * 		status: string
	 * 	}[]
	 * }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/course/recent')
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
		<div class="flex gap-2 text-sm font-semibold items-center">
			<GraduationCap size={20} />
			<p>Recently Added Courses</p>
			{#if data && data.totalCourses}
				<p class="text-xs text-neutral-500 font-normal ml-4">{data.totalCourses} Courses Found</p>
			{/if}
		</div>

		{#if !fetchError && (!data || data.courses.length == 0)}
			<p class="text-sm text-center text-neutral-400 dark:text-neutral-600">
				No Recent Courses Found
			</p>
		{/if}

		{#if data && data.courses.length > 0}
			<Table.Root class="text-xs">
				<Table.Header>
					<Table.Row class="hover:bg-transparent">
						<Table.Head class="font-semibold">Timestamp</Table.Head>
						<Table.Head class="font-semibold">Code</Table.Head>
						<Table.Head class="font-semibold">Name</Table.Head>
						<Table.Head class="font-semibold">Category</Table.Head>
						<Table.Head class="font-semibold">Enrolled</Table.Head>
						<Table.Head class="font-semibold">Instructor(s)</Table.Head>
						<Table.Head class="font-semibold text-right">Status</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each data.courses as course}
						<Table.Row>
							<Table.Cell>{new Date(course.timestamp).toLocaleString()}</Table.Cell>
							<Table.Cell>{course.code}</Table.Cell>
							<Table.Cell>{course.name}</Table.Cell>
							<Table.Cell>{course.category}</Table.Cell>
							<Table.Cell>{course.enrolled} Enrolled</Table.Cell>
							<Table.Cell>{course.instructor.join(', ')}</Table.Cell>
							<Table.Cell>
								<span class="flex justify-end items-center gap-2">
									{#if course.status == 'active'}
										Active <span><CircleCheck size={16} /></span>
									{:else if course.status == 'deactivated'}
										Deactivated <span><CircleAlert size={16} /></span>
									{/if}
								</span></Table.Cell
							>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		{/if}
	</div>

	<div class="flex gap-8 mt-10">
		<a href="/">
			<span class="flex text-xs items-center gap-2 font-semibold"
				><Plus size={20} /> Create A Course</span
			>
		</a>
		<a href="/">
			<span class="flex text-xs items-center gap-2 font-semibold"
				><List size={20} /> View All Courses</span
			>
		</a>
	</div>
</div>
