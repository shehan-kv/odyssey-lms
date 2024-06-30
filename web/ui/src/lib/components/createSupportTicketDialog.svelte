<script>
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Select from '$lib/components/ui/select';
	import { createEventDispatcher } from 'svelte';
	import Button from './ui/button/button.svelte';
	import { toast } from 'svelte-sonner';

	import { Textarea } from './ui/textarea';
	import { Input } from './ui/input';

	const dispatch = createEventDispatcher();

	export let open = false;
	let loading = false;

	let ticketData = {
		subject: '',
		type: '',
		description: ''
	};

	$: isValid = ticketData.subject != '' && ticketData.type != '' && ticketData.description != '';

	const createTicket = () => {
		if (!isValid) {
			return;
		}
		loading = true;
		fetch('/api/support-ticket', { method: 'POST', body: JSON.stringify(ticketData) })
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not create ticket');
				} else {
					toast.success('Ticket created');
					dispatch('close');
				}
			})
			.catch(() => {
				toast.error('Could not create ticket');
			});
		loading = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => dispatch('close')}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create Support Ticket</Dialog.Title>
			<Dialog.Description>
				Fill out the information below to submit a support request
			</Dialog.Description>
		</Dialog.Header>

		<div class="space-y-2">
			<Input type="text" id="subject" placeholder="Subject" bind:value={ticketData.subject} />

			<Select.Root
				onSelectedChange={(e) => {
					ticketData.type = e?.value;
				}}
			>
				<Select.Trigger>
					<Select.Value placeholder="Select Ticket Type" />
				</Select.Trigger>
				<Select.Content>
					<Select.Item class="pr-2" value="system">System</Select.Item>
					<Select.Item class="pr-2" value="course">Course</Select.Item>
					<Select.Item class="pr-2" value="user">User</Select.Item>
				</Select.Content>
			</Select.Root>

			<Textarea placeholder="Describe your issue" bind:value={ticketData.description} />
		</div>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => dispatch('close')}>Cancel</Button
			>
			<Button disabled={loading || !isValid} on:click={createTicket}>
				{loading ? 'Creating...' : 'Create'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
