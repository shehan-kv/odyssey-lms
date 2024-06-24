<script>
	import { Bell, HelpCircle, Menu, MessageSquare, User } from 'lucide-svelte';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import Logo from './logo.svelte';
	import ThemeToggle from './themeToggle.svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';
	import { goto } from '$app/navigation';

	let dispatch = createEventDispatcher();
</script>

<header class="py-2 px-4 flex justify-between items-center bg-white dark:bg-neutral-950 rounded">
	<div class="flex items-center gap-4">
		<Button
			on:click={() => dispatch('mobileMenuToggle')}
			class="xl:hidden p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
		>
			<Menu />
		</Button>
		<Logo size={32} />
	</div>

	<div class="flex gap-6">
		<ThemeToggle
			class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
		/>
		<DropdownMenu.Root>
			<DropdownMenu.Trigger><HelpCircle size={20} /></DropdownMenu.Trigger>
			<DropdownMenu.Content>
				<DropdownMenu.Item class="text-xs" on:click={() => goto('/my-support-tickets')}>
					My Support Tickets
				</DropdownMenu.Item>
				<DropdownMenu.Item class="text-xs" on:click={() => dispatch('newTicket')}>
					New Support Ticket
				</DropdownMenu.Item>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
		<Button
			class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
		>
			<Bell />
		</Button>
		<Button
			class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100  w-5 hover:bg-transparent"
		>
			<MessageSquare />
		</Button>
		<Button
			class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
		>
			<User />
		</Button>
	</div>
</header>
