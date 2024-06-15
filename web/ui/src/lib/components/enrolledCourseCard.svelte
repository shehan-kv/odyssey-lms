<script>
	import ImagePlaceholder from './imagePlaceholder.svelte';
	import LecternIcon from './lecternIcon.svelte';

	/**
	 * @type {{
	 * 	img: string,
	 * 	name: string,
	 * 	code: string,
	 * 	description:string,
	 * 	instructor : string[],
	 * 	completed: number
	 * }}
	 */
	export let course;
</script>

<div class="max-w-md">
	<div class="space-y-2">
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
		<p>
			<span class="flex items-center gap-2 text-sm"
				><LecternIcon size={16} />{course.instructor.join(', ')}</span
			>
		</p>
	</div>
	<div class="flex items-center gap-8 mt-4">
		<div class="h-1.5 rounded w-full bg-neutral-200 dark:bg-neutral-900 overflow-hidden">
			<div
				class="h-1.5 rounded bg-neutral-400 dark:bg-neutral-600"
				style={`width: ${course.completed}%;`}
			></div>
		</div>
		<p class="text-xs">{course.completed}%</p>
	</div>
</div>
