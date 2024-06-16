<script>
	import { Users2 } from 'lucide-svelte';
	import ImagePlaceholder from './imagePlaceholder.svelte';

	/**
	 * @type {{
	 * 	img: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
	 *  category: string,
	 * 	students: number
	 * }}
	 */
	export let course;
</script>

<div class="max-w-md space-y-2">
	<div class="h-60 rounded overflow-hidden">
		{#if course && course.img}
			<img src={course.img} alt={course.name} />
		{:else}
			<ImagePlaceholder />
		{/if}
	</div>
	<p class="font-semibold">{course.name}</p>
	<p class="text-neutral-500 text-sm">{course.code} | <span>{course.category}</span></p>
	<p class="text-sm text-ellipsis overflow-hidden h-20">
		{course.description}
	</p>
	<p>
		<span class="flex items-center gap-2 text-sm"><Users2 size={16} />{course.students}</span>
	</p>
</div>
