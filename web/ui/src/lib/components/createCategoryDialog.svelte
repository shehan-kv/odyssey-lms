<script>
	import * as Dialog from '$lib/components/ui/dialog';
	import { createEventDispatcher } from 'svelte';
	import Button from './ui/button/button.svelte';
	import { toast } from 'svelte-sonner';
	import Label from './ui/label/label.svelte';
	import Input from './ui/input/input.svelte';

	let dispatch = createEventDispatcher();

	export let open = false;
	let loading = false;

	let name = '';

	const activateUser = () => {
		if (!name) return;
		loading = true;

		fetch(`/api/course/category`, { method: 'POST', body: JSON.stringify({ name: name }) })
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not create category');
				} else {
					name = '';
					toast.success('Category created');
					dispatch('created');
				}
			})
			.catch(() => {
				toast.error('Could not create category');
			});
		loading = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => (open = !open)}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create Category</Dialog.Title>
		</Dialog.Header>
		<div>
			<Input type="text" id="name" placeholder="Name" bind:value={name} />
		</div>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => (open = false)}>Cancel</Button>
			<Button disabled={loading || name == ''} on:click={activateUser}>
				{loading ? 'Creating...' : 'Create'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
