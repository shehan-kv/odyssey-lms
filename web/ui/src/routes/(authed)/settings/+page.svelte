<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { MoveRight, Settings } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';

	let loading = true;
	let fetchError = false;

	let nameData = {
		firstName: '',
		lastName: '',
		bio: ''
	};

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/user/self`)
			.then((response) => {
				if (!response.ok) {
					fetchError = true;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				nameData.firstName = parsedData.firstName;
				nameData.lastName = parsedData.lastName;
				nameData.bio = parsedData.bio;
			})
			.catch(() => {
				fetchError = true;
			});
		loading = false;
	};

	onMount(() => {
		fetchData();
	});

	let nameUpdating = false;
	function updateName() {
		nameUpdating = true;
		fetch(`/api/user/self`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				firstName: nameData.firstName,
				lastName: nameData.lastName,
				bio: nameData.bio
			})
		})
			.then((response) => {
				console.log('calling updatenmae');
				if (!response.ok) {
					nameUpdating = false;
					toast.error('Name could not be updated');
				} else {
					toast.success('Name updated successfully');
					fetchData();
				}
			})
			.catch(() => {
				nameUpdating = false;
				toast.error('Name could not be updated');
			});

		nameUpdating = false;
	}

	let passwordData = {
		currentPassword: '',
		newPassword: '',
		confirmNewPassword: ''
	};
	let passwordUpdating = false;
	function updatePassword() {
		passwordUpdating = true;

		if (passwordData.newPassword !== passwordData.confirmNewPassword) {
			passwordUpdating = false;
			toast.error("Passwords don't match");
			return;
		}

		fetch(`/api/user/self/password`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(passwordData)
		})
			.then((response) => {
				if (!response.ok) {
					toast.error('Password could not be updated');
				} else {
					toast.success('Password updated successfully');
					passwordData.currentPassword = '';
					passwordData.newPassword = '';
					passwordData.confirmNewPassword = '';
				}
			})
			.catch(() => {
				toast.error('Password could not be updated');
			});
		passwordUpdating = false;
	}
</script>

<svelte:head>
	<title>Settings</title>
</svelte:head>

<div
	class="relative flex flex-col py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll grow"
>
	<div class="flex gap-2 text-sm font-semibold">
		<Settings size={20} />
		<p>Settings</p>
	</div>

	<div class="text-sm mt-10 max-w-lg space-y-10">
		<div>
			<p class="font-bold">Update Name</p>

			<form on:submit|preventDefault={updateName} class="mt-3 space-y-4">
				<div class="space-y-1.5">
					<Label for="firstName">First Name</Label>
					<Input type="text" id="firstName" bind:value={nameData.firstName} required />
				</div>
				<div class="space-y-1.5">
					<Label for="lastName">Last Name</Label>
					<Input type="text" id="lastName" bind:value={nameData.lastName} required />
				</div>
				<div class="space-y-1.5">
					<Label for="bio">Bio</Label>
					<Textarea type="text" id="bio" bind:value={nameData.bio} />
				</div>
				<Button disabled={nameUpdating} type="submit" class="ml-auto">
					{nameUpdating ? 'Updating...' : 'Update'}
					<MoveRight class="ml-3" />
				</Button>
			</form>
		</div>
		<div>
			<p class="font-bold">Update Password</p>

			<form on:submit|preventDefault={updatePassword} class="mt-3 space-y-4">
				<div class="space-y-1.5">
					<Label for="currentPassword">Current Password</Label>
					<Input
						type="password"
						id="currentPassword"
						bind:value={passwordData.currentPassword}
						required
					/>
				</div>
				<div class="space-y-1.5">
					<Label for="newPassword">New Password</Label>
					<Input type="password" id="newPassword" bind:value={passwordData.newPassword} required />
				</div>
				<div class="space-y-1.5">
					<Label for="confirmNewPassword">Confirm New Password</Label>
					<Input
						type="password"
						id="confirmNewPassword"
						bind:value={passwordData.confirmNewPassword}
						required
					/>
				</div>
				<Button disabled={passwordUpdating} type="submit">
					{passwordUpdating ? 'Updating...' : 'Update'}
					<MoveRight class="ml-3" />
				</Button>
			</form>
		</div>
	</div>
</div>

<Toaster />
