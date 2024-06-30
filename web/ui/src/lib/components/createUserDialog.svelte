<script>
	import { createEventDispatcher } from 'svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import * as Select from '$lib/components/ui/select';
	import Button from './ui/button/button.svelte';
	import { toast } from 'svelte-sonner';
	import Label from './ui/label/label.svelte';
	import Input from './ui/input/input.svelte';
	import Checkbox from './ui/checkbox/checkbox.svelte';

	let dispatch = createEventDispatcher();

	export let open = false;
	let loading = false;

	let user = {
		firstName: '',
		lastName: '',
		email: '',
		password: '',
		isActive: false
	};

	/**
	 * @type {{value: string, label: string}}
	 */
	let userRole = { value: 'administrator', label: 'Administrator' };

	const createUser = () => {
		if (!user) return;
		loading = true;
		fetch(`/api/user`, { method: 'POST', body: JSON.stringify({ ...user, role: userRole.value }) })
			.then((response) => {
				if (!response.ok) {
					toast.error('Could not create account');
				} else {
					toast.success('Account Created');
					user = { firstName: '', lastName: '', email: '', password: '', isActive: false };
					dispatch('create');
					dispatch('close');
				}
			})
			.catch(() => {
				toast.error('Could not create account');
			});
		loading = false;
	};
</script>

<Dialog.Root {open} onOpenChange={() => dispatch('close')}>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Create Account</Dialog.Title>
		</Dialog.Header>
		<div class="space-y-2">
			<div class="space-y-1.5">
				<Label for="firstName">First Name</Label>
				<Input type="text" id="firstName" bind:value={user.firstName} />
			</div>
			<div class="space-y-1.5">
				<Label for="lastName">Last Name</Label>
				<Input type="text" id="lastName" bind:value={user.lastName} />
			</div>
			<div class="space-y-1.5">
				<Label for="email">Email</Label>
				<Input type="email" id="email" bind:value={user.email} />
			</div>
			<div class="space-y-1.5">
				<Label for="password">Password</Label>
				<Input type="password" id="password" bind:value={user.password} />
			</div>
			<div class="flex items-center gap-2">
				<Checkbox id="isActive" bind:checked={user.isActive} class="rounded" />
				<Label for="isActive" class="text-sm">Activate</Label>
			</div>
			<div class="space-y-1.5">
				<Label>Role</Label>
				<Select.Root bind:selected={userRole}>
					<Select.Trigger class="w-[180px] rounded">
						<Select.Value placeholder="Administrator" />
					</Select.Trigger>
					<Select.Content>
						<Select.Item class="pr-2" value="administrator">Administrator</Select.Item>
						<Select.Item class="pr-2" value="student">Student</Select.Item>
					</Select.Content>
				</Select.Root>
			</div>
		</div>
		<div>
			<Button variant="outline" disabled={loading} on:click={() => dispatch('close')}>Cancel</Button
			>
			<Button disabled={loading} on:click={createUser}>
				{loading ? 'Creating...' : 'Create'}
			</Button>
		</div>
	</Dialog.Content>
</Dialog.Root>
