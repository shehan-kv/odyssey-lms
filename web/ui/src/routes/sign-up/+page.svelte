<script>
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import Logo from '$lib/components/logo.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import Checkbox from '$lib/components/ui/checkbox/checkbox.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { MoveRight } from 'lucide-svelte';
	import ThemeToggle from '$lib/components/themeToggle.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import { toast } from 'svelte-sonner';
	import FullScreenLoader from '$lib/components/fullScreenLoader.svelte';

	const formState = { firstName: '', lastName: '', email: '', password: '', confirmPassword: '' };

	const formSubmit = () => {
		fetch('/api/auth/sign-up', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(formState)
		})
			.then((response) => {
				if (!response.ok) {
					toast.error('Sign up Failed');
					return;
				}

				toast.success('Signed Up');
				goto('/sign-in');
			})
			.catch(() => {
				toast.error('Sign up Failed');
			});
	};

	let loading = true;

	onMount(() => {
		fetch('/api/auth/is-signed-in')
			.then((response) => {
				if (response.status == 200) {
					goto('/');
				} else {
					loading = false;
				}
			})
			.catch(() => {
				loading = false;
			});
	});
</script>

<svelte:head>
	<title>Sign Up</title>
</svelte:head>

{#if loading}
	<FullScreenLoader />
{/if}

<div
	class="-z-10 fixed h-screen w-full bg-neutral-100 dark:bg-neutral-900 bg-[radial-gradient(#ddd_1px,transparent_1px)] dark:bg-[radial-gradient(#333_1px,transparent_1px)] [background-size:2rem_2rem]"
></div>

<ThemeToggle
	class="absolute top-5 right-4 h-14 rounded-full border bg-white hover:bg-neutral-100 text-neutral-900 dark:bg-neutral-950 dark:hover:bg-neutral-900 dark:text-neutral-100"
/>

<div class="min-h-screen px-4 sm:px-0 flex justify-center items-center">
	<div
		class="p-px my-24 rounded-xl bg-gradient-to-br from-emerald-300 via-transparent to-emerald-300 dark:from-emerald-700 dark:via-transparent dark:to-emerald-700 w-full max-w-lg"
	>
		<div
			class="bg-white dark:bg-neutral-950 dark:text-neutral-200 py-10 sm:py-16 px-6 sm:px-14 rounded-xl"
		>
			<div class="mb-4">
				<Logo size="48" />
			</div>
			<h1 class="text-xl mb-10">Sign Up</h1>

			<form on:submit|preventDefault={formSubmit}>
				<div class="flex gap-4">
					<div class="flex flex-col gap-1.5">
						<Label for="firstName" class="text-sm">First Name</Label>
						<Input
							id="firstName"
							type="text"
							class="rounded"
							required
							bind:value={formState.firstName}
						/>
					</div>

					<div class="flex flex-col gap-1.5">
						<Label for="lastName" class="text-sm">Last Name</Label>
						<Input
							id="lastName"
							type="text"
							class="rounded"
							required
							bind:value={formState.lastName}
						/>
					</div>
				</div>

				<div class="flex flex-col gap-1.5 mt-6">
					<Label for="email" class="text-sm">Email</Label>
					<Input id="email" type="email" class="rounded" required bind:value={formState.email} />
				</div>

				<div class="flex gap-4">
					<div class="flex flex-col gap-1.5 mt-6">
						<Label for="password" class="text-sm">Password</Label>
						<Input
							id="password"
							type="password"
							class="rounded"
							required
							bind:value={formState.password}
						/>
					</div>

					<div class="flex flex-col gap-1.5 mt-6">
						<Label for="confirmPassword" class="text-sm">Confirm Password</Label>
						<Input
							id="confirmPassword"
							type="password"
							class="rounded"
							required
							bind:value={formState.confirmPassword}
						/>
					</div>
				</div>

				<Button
					type="submit"
					class="w-full rounded font-bold bg-emerald-600 hover:bg-emerald-700 text-emerald-50 dark:bg-emerald-300 dark:hover:bg-emerald-400
					dark:text-emerald-950
					mt-7"
				>
					Sign Up <MoveRight class="ml-4 dark:text-emerald-950" />
				</Button>
			</form>
			<div class="mt-6">
				<p class="text-sm">
					Already have an account ?
					<a href="/sign-in" class="text-emerald-600 dark:text-emerald-300"> Sign In </a>
				</p>
			</div>
		</div>
	</div>
</div>

<Toaster />
