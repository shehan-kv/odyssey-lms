<script>
	import { Users2 } from 'lucide-svelte';
	import ImagePlaceholder from './imagePlaceholder.svelte';

	/**
	 * @type {{
	 * 	image: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
	 *  category: string,
	 * 	enrolled: number
	 * }}
	 */
	export let course;
</script>

<div class="max-w-md space-y-2">
	<div class="h-60 rounded overflow-hidden">
		{#if course && course.image}
			<img src={`/uploads/${course.image}`} alt={course.name} class="object-cover" />
		{:else}
			<ImagePlaceholder />
		{/if}
	</div>
	<p class="font-semibold">{course.name}</p>
	<p class="text-neutral-500 text-sm">{course.code} | <span>{course.category}</span></p>
	<p class="text-sm text-ellipsis overflow-hidden h-20">
		{course.description}
	</p>
</div>
