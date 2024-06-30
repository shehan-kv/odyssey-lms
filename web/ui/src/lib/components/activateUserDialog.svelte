<script>
	import * as Dialog from '$lib/components/ui/dialog';
	import { createEventDispatcher } from 'svelte';
	import Button from './ui/button/button.svelte';
	import { toast } from 'svelte-sonner';

	let dispatch = createEventDispatcher();

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

	const activateUser = () => {
		if (!user) return;
		loading = true;
		fetch(`/api/user/activate/${user.id}`, { method: 'POST' })
			.then((response) => {
				if (response.status != 200) {
					toast.error('Could not activate account');
				} else {
					toast.success('Account activated');
					dispatch('activate');
				}
			})
			.catch(() => {
				toast.error('Could not activate account');
			});
		loading = false;
		open = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => (open = !open)}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Activate Account ?</Dialog.Title>
			<Dialog.Description>
				This will Activate the account <span class="font-semibold"
					>({user?.firstName} {user?.lastName})</span
				> and the user will be able to sign in. Are you sure you want to activate ?
			</Dialog.Description>
		</Dialog.Header>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => (open = false)}>Cancel</Button>
			<Button disabled={loading} on:click={activateUser}>
				{loading ? 'Activating...' : 'Activate'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
