<script>
	// @ts-nocheck

	import { afterNavigate, beforeNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import CourseSectionPane from '$lib/components/courseSectionPane.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { BookOpenText, CheckCheck } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';

	let loading = true;
	let fetchError = false;
	let contentPaneOpen = false;

	/**
	 * @type {{ id: number, title: string, isComplete: boolean, content: {id:number, type:string, data: any[]}[] }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetch(`/api/course/${$page.params.id}/enroll/section/${$page.params.section}`)
			.then((response) => {
				if (!response.ok) {
					fetchError = true;
					return;
				} else {
					return response.json();
				}
			})
			.then(
				(
					/** @type {{
					 * id: number,
					 * title: string,
					 * isComplete: boolean,
					 * content: string }} */
					parsedData
				) => {
					data = {
						id: parsedData.id,
						title: parsedData.title,
						isComplete: parsedData.isComplete,
						content: JSON.parse(parsedData.content)
					};
				}
			)
			.catch(() => (fetchError = true));

		loading = false;
	};

	let completing = false;
	const markAsComplete = () => {
		completing = true;
		fetch(`/api/course/${$page.params.id}/enroll/section/${$page.params.section}/complete`, {
			method: 'POST'
		})
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not mark as completed');
					return;
				} else {
					data = { ...data, isComplete: true };
				}
			})
			.catch(() => toast.error('Could not mark as completed'));
		completing = false;
	};

	afterNavigate(() => {
		fetchData();
	});
</script>

<svelte:head>
	<title>{data && data.title ? data.title : 'Course Section'}</title>
</svelte:head>

<div
	class="relative flex flex-col gap-2 py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll grow"
>
	<div class="flex justify-end">
		<Button
			on:click={() => (contentPaneOpen = !contentPaneOpen)}
			class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
		>
			<BookOpenText size={20} />
		</Button>
	</div>

	{#if data && data.content.length > 0}
		<div class="mt-10 mx-auto w-full max-w-4xl">
			<p class="text-xl mb-10 font-semibold">{data.title}</p>
			{#each data.content as content}
				<!-- Headers -->
				{#if content.type == 'header' && content.data.level == 1}
					<h1 class="text-5xl my-4 font-semibold">{@html content.data.text}</h1>
				{:else if content.type == 'header' && content.data.level == 2}
					<h2 class="text-4xl my-4 font-semibold">{@html content.data.text}</h2>
				{:else if content.type == 'header' && content.data.level == 3}
					<h3 class="text-3xl my-4 font-semibold">{@html content.data.text}</h3>
				{:else if content.type == 'header' && content.data.level == 4}
					<h4 class="text-2xl my-4 font-semibold">{@html content.data.text}</h4>
				{:else if content.type == 'header' && content.data.level == 5}
					<h5 class="text-xl my-4 font-semibold">{@html content.data.text}</h5>
				{:else if content.type == 'header' && content.data.level == 6}
					<h6 class="text-lg my-4 font-semibold">{@html content.data.text}</h6>

					<!-- PARAGRAPH -->
				{:else if content.type == 'paragraph'}
					<p class="my-4">{@html content.data.text}</p>
				{/if}
			{/each}

			<div class="mt-10 flex justify-end">
				{#if data.isComplete}
					<div class="flex gap-2 items-center text-emerald-600 dark:text-emerald-400">
						<CheckCheck size={16} />
						<p>Already Completed</p>
					</div>
				{:else}
					<Button on:click={markAsComplete} disabled={completing}>
						<CheckCheck size={16} class="mr-2" />
						{completing ? 'Completing' : 'Mark As Complete'}
					</Button>
				{/if}
			</div>
		</div>
	{/if}
</div>

{#if contentPaneOpen}
	<CourseSectionPane on:contentPaneToggle={() => (contentPaneOpen = !contentPaneOpen)} />
{/if}
