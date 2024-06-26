<script>
	import { Send } from 'lucide-svelte';
	import Button from './ui/button/button.svelte';
	import Textarea from './ui/textarea/textarea.svelte';
	import { toast } from 'svelte-sonner';
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher();

	export let ticketId;

	let message = '';
	let sending = false;

	const sendReply = () => {
		sending = true;
		if (message == '') {
			toast.error('Message is empty');
			sending = false;
			return;
		}

		fetch(`/api/support-ticket/${ticketId}/message`, {
			method: 'POST',
			body: JSON.stringify({ message: message })
		})
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not send message');
					return;
				}

				message = '';
				dispatch('success');
			})
			.catch(() => {
				toast.error('Could not send message');
			});

		sending = false;
	};
</script>

<div class="mt-20 max-w-4xl pl-10 py-10 border-t">
	<p class="text-sm font-semibold mb-2">Send A Reply</p>
	<Textarea bind:value={message} class="h-56" />
	<Button
		disabled={sending || message == ''}
		on:click={sendReply}
		class="w-full mt-4 py-6 font-semibold"
	>
		{sending ? 'Sending...' : 'Send'}
		<Send size={16} class="ml-2" />
	</Button>
</div>
