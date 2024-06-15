<script>
	import { Users2 } from 'lucide-svelte';
	import ImagePlaceholder from './imagePlaceholder.svelte';

	/**
	 * @type {{
	 * 	img: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
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
	<p class="text-neutral-500">{course.code}</p>
	<p class="text-sm text-ellipsis overflow-hidden h-20">
		{course.description}
	</p>
	<p><span><Users2 size={20} class="inline mr-2" /></span>{course.students}</p>
</div>
