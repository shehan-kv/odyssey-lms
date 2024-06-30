<script>
	import { CheckCircle2, X } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher, onMount } from 'svelte';
	import CardLoader from './cardLoader.svelte';
	import { page } from '$app/stores';
	import CardErrorRetry from './cardErrorRetry.svelte';

	let dispatch = createEventDispatcher();

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{id: number, title: string, isComplete: boolean}[]}
	 */
	let data = [];

	const fetchData = () => {
		loading = true;
		fetch(`/api/course/${$page.params.id}/enroll/section`)
			.then((response) => {
				if (!response.ok) {
					fetchError = true;
					return;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				data = parsedData;
			})
			.catch(() => (fetchError = true));

		loading = false;
	};

	onMount(() => {
		fetchData();
	});
</script>

<div class="z-50 fixed inset-0">
	<div
		tabindex="-1"
		role="button"
		on:keydown={() => dispatch('contentPaneToggle')}
		on:click={() => dispatch('contentPaneToggle')}
		class="w-full h-full bg-neutral-100/70 dark:bg-neutral-900/70 backdrop-blur-sm"
	></div>
	<div
		class="fixed right-0 inset-y-0 bg-white dark:bg-neutral-950 rounded py-4 w-full max-w-96 m-2 overflow-scroll"
	>
		{#if loading}
			<CardLoader />
		{/if}
		{#if fetchError}
			<CardErrorRetry on:retry={fetchData} />
		{/if}

		<div class="flex justify-end px-4 mb-10">
			<Button
				on:click={() => dispatch('contentPaneToggle')}
				class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent z-20"
			>
				<X />
			</Button>
		</div>

		{#if data && data.length > 0}
			<div class="px-4 text-sm">
				{#each data as section}
					<a
						on:click={() => dispatch('contentPaneToggle')}
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
</div>
