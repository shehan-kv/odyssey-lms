<script>
	import { page } from '$app/stores';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import ImagePlaceholder from '$lib/components/imagePlaceholder.svelte';
	import { CheckCircle2 } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 *  id: number,
	 * 	image: string,
	 * 	name: string,
	 *  code: string,
	 *  category: string,
	 *  description: string,
	 *  sections: { id: number, title: string, isComplete: boolean }[]
	 * }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/course/${$page.params.id}/enroll`)
			.then((response) => {
				if (!response.ok) {
					fetchError = true;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				data = parsedData;
			})
			.catch(() => {
				fetchError = true;
			});
		loading = false;
	};

	onMount(() => {
		fetchData();
	});
</script>

<svelte:head>
	<title>{data ? data.name : 'Course'}</title>
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

	{#if data}
		<div class="mx-auto max-w-4xl">
			<div class="w-full h-96 overflow-hidden rounded">
				{#if data.image}
					<img src={`/uploads/${data.image}`} alt={data.name} class="w-full h-full object-cover" />
				{:else}
					<ImagePlaceholder />
				{/if}
			</div>
			<div class="mt-4">
				<div class="flex justify-between">
					<div>
						<p class="text:lg 2xl:text-2xl font-semibold">
							{data.name}
						</p>
						<p class="text-sm mt-2 text-neutral-500">{data.code} | {data.category}</p>
					</div>
				</div>
				<p class="mt-4">
					{data.description}
				</p>
			</div>

			{#if data && data.sections.length > 0}
				<div class="mt-10">
					<p class="font-semibold mb-4 text-neutral-500">Course Content</p>
					{#each data.sections as section (section)}
						<a
							href={`/enrolled-courses/${$page.params.id}/${section.id}`}
							class="block flex items-center justify-between py-3 px-4 border-b hover:bg-neutral-100 hover:dark:bg-neutral-900"
						>
							<span>{section.title}</span>
							{#if section.isComplete}
								<CheckCircle2 size={16} class="text-emerald-600 dark:text-emerald-400" />
							{/if}
						</a>
					{/each}
				</div>
			{/if}
		</div>
	{/if}
</div>
