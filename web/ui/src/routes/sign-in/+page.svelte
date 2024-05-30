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
	import FullScreenLoader from '$lib/components/ui/fullScreenLoader.svelte';

	const formState = { email: '', password: '', remember_me: false };

	const formSubmit = () => {
		fetch('/api/auth/sign-in', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(formState)
		})
			.then((response) => {
				if (response.status !== 200) {
					toast.error('Sign in Failed');
					return;
				}

				toast.success('Signed In');
				goto('/');
			})
			.catch(() => {
				toast.error('Sign in Failed');
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
	<title>Sign In</title>
</svelte:head>

{#if loading}
	<FullScreenLoader />
{/if}

<div class="fixed top-10 left-10 text-neutral-900 dark:text-neutral-100">
	<p class="sm:hidden">xs</p>
	<p class="hidden sm:block md:hidden">sm</p>
	<p class="hidden md:block lg:hidden">md</p>
	<p class="hidden lg:block xl:hidden">lg</p>
	<p class="hidden xl:block 2xl:hidden">xl</p>
	<p class="hidden 2xl:block 3xl:hidden">2xl</p>
	<p class="hidden 3xl:block 4xl:hidden">3xl</p>
</div>

<div
	class="-z-10 fixed h-screen w-full bg-neutral-100 dark:bg-neutral-900 bg-[radial-gradient(#ddd_1px,transparent_1px)] dark:bg-[radial-gradient(#333_1px,transparent_1px)] [background-size:2rem_2rem]"
></div>

<ThemeToggle
	class="absolute top-5 right-10 h-14 rounded-full border bg-white hover:bg-neutral-100 text-neutral-900 dark:bg-neutral-950 dark:hover:bg-neutral-900 dark:text-neutral-100"
/>

<div class="min-h-screen px-4 sm:px-0 flex justify-center items-center">
	<div
		class="p-px my-24 rounded-xl bg-gradient-to-br from-emerald-300 via-transparent to-emerald-300 dark:from-emerald-700 dark:via-transparent dark:to-emerald-700 w-full max-w-md"
	>
		<div
			class="bg-white dark:bg-neutral-950 dark:text-neutral-200 py-10 sm:py-16 px-6 sm:px-14 rounded-xl"
		>
			<div class="mb-4">
				<Logo size="48" />
			</div>
			<h1 class="text-xl mb-10">Sign In</h1>

			<form on:submit|preventDefault={formSubmit}>
				<div class="flex flex-col gap-1.5">
					<Label for="email" class="text-sm">Email</Label>
					<Input id="email" type="email" class="rounded" required bind:value={formState.email} />
				</div>

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

				<div class="flex items-center gap-3 mt-6">
					<Checkbox
						id="remember_me"
						bind:checked={formState.remember_me}
						class="rounded data-[state=checked]:bg-emerald-600 data-[state=checked]:text-emerald-50 dark:data-[state=checked]:bg-emerald-300 dark:data-[state=checked]:text-emerald-950"
					/>
					<Label for="remember_me" class="text-sm">Remember Me</Label>
				</div>

				<Button
					type="submit"
					class="w-full rounded font-bold bg-emerald-600 hover:bg-emerald-700 text-emerald-50 dark:bg-emerald-300 dark:hover:bg-emerald-400
					dark:text-emerald-950
					mt-7"
				>
					Sign in <MoveRight class="ml-4 dark:text-emerald-950" />
				</Button>
			</form>
		</div>
	</div>
</div>

<Toaster />
