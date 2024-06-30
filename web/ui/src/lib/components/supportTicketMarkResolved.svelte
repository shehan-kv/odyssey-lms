<script>
	import { toast } from 'svelte-sonner';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';

	export let id;

	let dispatch = createEventDispatcher();
	let marking = false;

	const markAsResolved = () => {
		marking = true;
		fetch(`/api/support-ticket/${id}/resolve`, { method: 'POST' })
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not mark as resolved');
					return;
				}

				toast.success('Marked as resolved');
				dispatch('success');
			})
			.catch(() => {
				toast.error('Could not mark as resolved');
			});

		marking = false;
	};
</script>

<Button disabled={marking} variant="secondary" on:click={markAsResolved} class="text-xs">
	{marking ? 'Marking...' : 'Mark As Resolved'}
</Button>
