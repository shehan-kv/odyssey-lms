<script>
	import Button from './ui/button/button.svelte';
	import Input from './ui/input/input.svelte';
	import EditorJS from '@editorjs/editorjs';
	import Header from '@editorjs/header';
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher();

	const editor = new EditorJS({
		tools: {
			header: Header
		}
	});

	let section = { title: '', content: [] };

	const addSection = () => {
		editor.save().then((outputData) => {
			// @ts-ignore
			section.content = outputData.blocks;
			dispatch('add', { ...section });
			section = { title: '', content: [] };
			editor.clear();
		});
	};
</script>

<div class="border-t space-y-2 pt-4">
	<p class="font-semibold">Add Section</p>
	<Input type="text" placeholder="Title" bind:value={section.title} />
	<div id="editorjs" class=" rounded p-2 border"></div>

	<Button variant="outline" on:click={addSection}>Add</Button>
</div>
