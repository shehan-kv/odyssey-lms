<script>
	import * as Dialog from '$lib/components/ui/dialog';
	import { createEventDispatcher } from 'svelte';
	import Button from './ui/button/button.svelte';
	import { toast } from 'svelte-sonner';

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

	const deactivateUser = () => {
		if (!user) return;
		loading = true;
		fetch(`/api/user/deactivate/${user.id}`, { method: 'POST' })
			.then((response) => {
				if (response.status != 200) {
					toast.error('Could not deactivate account');
				} else {
					toast.success('Account deactivated');
					dispatch('deactivate');
				}
			})
			.catch(() => {
				toast.error('Could not deactivate account');
			});
		loading = false;
		open = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => (open = !open)}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Deactivate Account ?</Dialog.Title>
			<Dialog.Description>
				This will deactivate the account <span class="font-semibold"
					>({user?.firstName} {user?.lastName})</span
				> and the user will not be able to sign in. Are you sure you want to deactivate ?
			</Dialog.Description>
		</Dialog.Header>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => (open = false)}>Cancel</Button>
			<Button variant="destructive" disabled={loading} on:click={deactivateUser}>
				{loading ? 'Deactivating...' : 'Deactivate'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
