<script>
	import ImagePlaceholder from './imagePlaceholder.svelte';
	import LecternIcon from './lecternIcon.svelte';

	/**
	 * @type {{
	 * 	id: number,
	 * 	image: string,
	 * 	name: string,
	 * 	code: string,
	 * 	category:string,
	 * 	description:string,
	 * }}
	 */
	export let course;
</script>

<div class="relative max-w-md">
	<div class="space-y-2">
		<div class="h-60 rounded overflow-hidden">
			{#if course && course.image}
				<img
					src={`/uploads/${course.image}`}
					alt={course.name}
					class="w-full h-full object-cover"
				/>
			{:else}
				<ImagePlaceholder />
			{/if}
		</div>
		<div class="space-y-1">
			<a
				href={`/enrolled-courses/${course.id}`}
				class="inline-block font-semibold after:content-[' '] after:absolute after:top-0 after:left-0 after:inset-0"
			>
				{course.name}
			</a>
			<p class="text-neutral-500 text-sm">{course.code} | <span>{course.category}</span></p>
		</div>
		<p class="text-sm line-clamp-4">
			{course.description}
		</p>
	</div>
</div>
