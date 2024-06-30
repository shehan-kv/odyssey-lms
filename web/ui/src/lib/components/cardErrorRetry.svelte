<script>
	import { RotateCcw, TriangleAlert } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher();
</script>

<div
	class="absolute z-20 inset-0 bg-white dark:bg-neutral-950
    top-0 left-0 text-neutral-600 dark:text-neutral-400"
>
	<div class="flex flex-col items-center justify-center h-full">
		<TriangleAlert />
		<p class="mt-2 text-sm">There was an error fetching data</p>
		<Button
			on:click={() => dispatch('retry')}
			class="p-0 bg-transparent h-fit text-neutral-600 dark:text-neutral-400 w-fit hover:bg-transparent space-x-2 mt-5"
		>
			<RotateCcw size={20} /> <span>Retry</span>
		</Button>
	</div>
</div>
