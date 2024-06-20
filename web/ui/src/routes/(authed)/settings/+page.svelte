<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import { MoveRight, Settings } from 'lucide-svelte';
	import { onMount } from 'svelte';
	import { toast } from 'svelte-sonner';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	id: number,
	 *  firstName: string,
	 *  lastName: string
	 * }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/user/self`)
			.then((response) => {
				if (response.status != 200) {
					fetchError = true;
					loading = false;
				} else {
					return response.json();
				}
			})
			.then((parsedData) => {
				data = parsedData;
				loading = false;
			})
			.catch(() => {
				fetchError = true;
				loading = false;
			});
	};

	onMount(() => {
		fetchData();
	});

	let nameUpdating = false;
	function updateName() {
		nameUpdating = true;
		fetch(`/api/user/self/name`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ firstName: data.firstName, lastName: data.lastName })
		})
			.then((response) => {
				if (response.status != 200) {
					nameUpdating = false;
					toast.error('Name could not be updated');
				} else {
					toast.success('Name updated successfully');
					fetchData();
				}
			})
			.catch(() => {
				toast.error('Name could not be updated');
			});
	}

	let passwordData = {
		currentPassword: '',
		newPassword: '',
		confirmNewPassword: ''
	};
	let passwordUpdating = false;
	function updatePassword() {
		nameUpdating = true;
		fetch(`/api/user/self/password`, {
			method: 'PUT',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(passwordData)
		})
			.then((response) => {
				if (response.status != 200) {
					nameUpdating = false;
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
					<Input type="text" id="firstName" />
				</div>
				<div class="space-y-1.5">
					<Label for="lastName">Last Name</Label>
					<Input type="text" id="lastName" />
				</div>
				<Button disabled={nameUpdating} type="submit" class="ml-auto"
					>{nameUpdating ? 'Updating...' : 'Update'} <MoveRight class="ml-3" /></Button
				>
			</form>
		</div>
		<div>
			<p class="font-bold">Update Password</p>

			<form on:submit|preventDefault={updatePassword} class="mt-3 space-y-4">
				<div class="space-y-1.5">
					<Label for="currentPassword">Current Password</Label>
					<Input type="password" id="currentPassword" bind:value={passwordData.currentPassword} />
				</div>
				<div class="space-y-1.5">
					<Label for="newPassword">New Password</Label>
					<Input type="password" id="newPassword" bind:value={passwordData.newPassword} />
				</div>
				<div class="space-y-1.5">
					<Label for="confirmNewPassword">Confirm New Password</Label>
					<Input
						type="password"
						id="confirmNewPassword"
						bind:value={passwordData.confirmNewPassword}
					/>
				</div>
				<Button disabled={passwordUpdating} type="submit"
					>{passwordUpdating ? 'Updating...' : 'Update'} <MoveRight class="ml-3" /></Button
				>
			</form>
		</div>
	</div>
</div>

<Toaster />
