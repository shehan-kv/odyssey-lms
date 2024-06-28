<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import ImagePlaceholder from '$lib/components/imagePlaceholder.svelte';
	import { Input } from '$lib/components/ui/input/index.js';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import * as Select from '$lib/components/ui/select';
	import { Button } from '$lib/components/ui/button/index.js';
	import { toast } from 'svelte-sonner';
	import { Plus, X } from 'lucide-svelte';
	import CreateCategoryDialog from '$lib/components/createCategoryDialog.svelte';
	import AddCourseSection from '$lib/components/addCourseSection.svelte';
	import CourseSectionItem from '$lib/components/courseSectionItem.svelte';
	import { Title } from 'chart.js';

	let dispatch = createEventDispatcher();
	let courseData = { name: '', code: '', description: '' };

	/**
	 * @type {{
	 * 	title: string,
	 * 	content: any,
	 * }[]}}
	 */
	let sections = [];
	let createLoading = false;

	/**
	 * @type {{
	 * 	value: string,
	 * 	label: string,
	 * }[]}}
	 */
	let categories = [{ value: '', label: 'No Category' }];
	let currentCategory = categories[0];

	let categoryOpen = false;

	const fetchCategories = () => {
		fetch('/api/course/category')
			.then((response) => {
				if (!response.ok) {
					return;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				categories = parsedData;
			});
	};

	/**
	 * @type {string | undefined}
	 */
	let coverImage;

	/**
	 * @type {File | undefined}
	 */
	let imageFile;

	const onFileSelected = (/** @type {Event} */ event) => {
		// @ts-ignore
		imageFile = event.target?.files[0];
		if (imageFile) {
			coverImage = URL.createObjectURL(imageFile);
		}
	};

	const createCourse = () => {
		if (!currentCategory.value) return;

		let formattedSections = sections.map((section) => ({
			title: section.title,
			content: JSON.stringify(section.content)
		}));

		const formData = new FormData();
		formData.append('name', courseData.name);
		formData.append('code', courseData.code);
		formData.append('description', courseData.description);
		formData.append('category', currentCategory.value);
		formData.append('sections', JSON.stringify(formattedSections));
		if (imageFile) {
			formData.append('image', imageFile);
		}

		createLoading = true;
		fetch('/api/course', {
			method: 'POST',
			body: formData
		})
			.then((response) => {
				if (response.ok) {
					sections = [];
					courseData = { name: '', code: '', description: '' };
					imageFile = undefined;
					coverImage = undefined;
					toast.success('Course created');
				} else {
					toast.error('Could not create course');
				}
			})
			.catch(() => {
				toast.error('Could not create course');
			});

		createLoading = false;
	};

	onMount(() => {
		fetchCategories();
	});
</script>

<svelte:head>
	<title>Create Course</title>
</svelte:head>

<div
	class="relative flex flex-col py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll grow"
>
	<div class="flex gap-2 text-sm font-semibold">
		<Plus size={20} />
		<p>Create Course</p>
	</div>
	<div class="flex gap-20">
		<div class="w-full max-w-2xl">
			<div class="mt-10">
				<div class="rounded overflow-hidden">
					{#if coverImage}
						<img src={coverImage} alt="d" />
					{:else}
						<ImagePlaceholder />
					{/if}
					<Input
						type="file"
						name="image"
						class="mt-2"
						on:change={(/** @type {Event} */ e) => onFileSelected(e)}
					/>
				</div>
			</div>

			<div class="mt-10 space-y-4 text-sm">
				<div class="space-y-1.5">
					<p>Name</p>
					<Input type="text" name="name" bind:value={courseData.name} />
				</div>
				<div class="space-y-1.5">
					<p>Code</p>
					<Input type="text" name="code" bind:value={courseData.code} />
				</div>
				<div class="space-y-1.5">
					<p>Description</p>
					<Textarea class="h-48" bind:value={courseData.description} />
				</div>
				<div>
					<p>Category</p>
					<div class="flex items-center gap-2 mt-1.5">
						<Select.Root bind:selected={currentCategory}>
							<Select.Trigger class="w-[180px] text-xs rounded">
								<Select.Value placeholder="All Categories" />
							</Select.Trigger>
							<Select.Content>
								{#each categories as category (category)}
									<Select.Item class="pr-2" value={category.value}>{category.label}</Select.Item>
								{/each}
							</Select.Content>
						</Select.Root>
						<Button variant="link" on:click={() => (categoryOpen = !categoryOpen)}>
							Create Category
						</Button>
					</div>
				</div>
			</div>
			<Button disabled={createLoading} class="mt-6 w-full" on:click={createCourse}>
				{createLoading ? 'Creating...' : 'Create Course'}
			</Button>
		</div>

		<div class="mt-10 text-sm w-full max-w-2xl">
			<div>
				{#each sections as section, index (index)}
					<CourseSectionItem
						{section}
						on:delete={() => {
							sections = sections.filter((e, i) => i != index);
						}}
					/>
				{/each}
			</div>
			<AddCourseSection
				on:add={(e) => {
					sections = [...sections, e.detail];
				}}
			/>
		</div>
	</div>
</div>

<CreateCategoryDialog
	open={categoryOpen}
	on:created={() => {
		categoryOpen = false;
		fetchCategories();
	}}
/>
