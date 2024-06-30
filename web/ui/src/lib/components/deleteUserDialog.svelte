<script>
	import * as Dialog from '$lib/components/ui/dialog';
	import { toast } from 'svelte-sonner';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let open = false;
	let loading = false;

	/**
	 * @type {{
	 * 	id: number,
	 *  firstName: string,
	 *  lastName: string
	 * } | null}
	 */
	export let user = null;

	const deleteUser = () => {
		if (!user) return;
		loading = true;
		fetch(`/api/user/${user.id}`, { method: 'DELETE' })
			.then((response) => {
				if (response.status != 204) {
					toast.error('Could not delete account');
				} else {
					toast.success('Account deleted');
					dispatch('delete');
				}
			})
			.catch(() => {
				toast.error('Could not delete account');
			});
		loading = false;
		open = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => (open = !open)}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Delete Account ?</Dialog.Title>
			<Dialog.Description>
				This will delete the account
				<span class="font-semibold">({user?.firstName} {user?.lastName})</span>
				and remove all associated data with the account. Are you sure you want to delete ?
			</Dialog.Description>
		</Dialog.Header>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => (open = false)}>Cancel</Button>
			<Button variant="destructive" disabled={loading} on:click={deleteUser}>
				{loading ? 'Deleting...' : 'Delete'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
