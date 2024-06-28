<script>
	import { ChevronDown, X } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher();
	export let section;

	let expanded = false;
</script>

<div class="py-4">
	<div role="button" class="flex justify-between">
		<div class="flex iems-center gap-2">
			<Button variant="ghost" class="p-1 h-fit w-5" on:click={() => (expanded = !expanded)}>
				<ChevronDown size={16} />
			</Button>
			<p class="font-semibold text-base">
				{section.title}
			</p>
		</div>
		<Button variant="ghost" class="p-1 h-fit w-5" on:click={() => dispatch('delete')}>
			<X size={16} />
		</Button>
	</div>

	{#if expanded}
		<div>
			{#each section.content as content}
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
		</div>
	{/if}
</div>
