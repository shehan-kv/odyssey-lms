<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { CircleAlert, CircleCheck, Ellipsis, Search, UserPlus, Users } from 'lucide-svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import * as Pagination from '$lib/components/ui/pagination';
	import * as Select from '$lib/components/ui/select';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Table from '$lib/components/ui/table';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import DeactivateUserDialog from '$lib/components/deactivateUserDialog.svelte';
	import DeleteUserDialog from '$lib/components/deleteUserDialog.svelte';
	import ActivateUserDialog from '$lib/components/activateUserDialog.svelte';
	import { Toaster } from '$lib/components/ui/sonner';
	import CardLoader from '$lib/components/cardLoader.svelte';
	import CardErrorRetry from '$lib/components/cardErrorRetry.svelte';

	let loading = true;
	let fetchError = false;

	let query = new URLSearchParams($page.url.searchParams.toString());

	let roles = [
		{ value: '', label: 'All Roles' },
		{ value: 'administrator', label: 'Administrator' },
		{ value: 'student', label: 'Student' }
	];

	let searchTextTemp = query.get('search') || ''; // search text variable to bind to the input
	let searchText = query.get('search') || '';
	let initialRole = query.get('role') || '';
	let currentRole = {
		value: roles.filter((role) => role.value == initialRole)[0]?.['value'] || roles[0].value,
		label: roles.filter((role) => role.value == initialRole)[0]?.['label'] || roles[0].label
	};

	let perPage = 30;
	let pageNum = Number(query.get('page')) || 1;

	$: {
		query.set('search', searchText);
		if (searchText == '') query.delete('search');

		query.set('role', currentRole.value);
		if (currentRole.value == '') query.delete('role');

		query.set('page', pageNum.toString());
		if (pageNum == 1) query.delete('page');

		goto(`?${query.toString()}`);
		fetchData();
	}

	/**
	 * @type {{
	 * 	totalCount: number,
	 * 	users: {
	 * 	id: number;
	 * 	createdAt: string,
	 * 	lastLogin: string,
	 * 	firstName: string,
	 * 	lastName:string,
	 *  email: string,
	 * 	isActive: boolean,
	 *  role: string
	 * }[]}}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch(`/api/user?${query.toString()}`)
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

	/**
	 * @type {{
	 *  open: boolean,
	 *  user: {
	 * 		id: number,
	 * 		firstName: string,
	 * 		lastName: string
	 * 	} | null
	 * }}
	 */
	let deactivateUser = {
		open: false,
		user: null
	};

	/**
	 * @type {{
	 *  open: boolean,
	 *  user: {
	 * 		id: number,
	 * 		firstName: string,
	 * 		lastName: string
	 * 	} | null
	 * }}
	 */
	let activateUser = {
		open: false,
		user: null
	};

	/**
	 * @type {{
	 *  open: boolean,
	 *  user: {
	 * 		id: number,
	 * 		firstName: string,
	 * 		lastName: string
	 * 	} | null
	 * }}
	 */
	let deleteUser = {
		open: false,
		user: null
	};
</script>

<svelte:head>
	<title>IAM</title>
</svelte:head>

<div
	class="relative flex flex-col py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-scroll w-full"
>
	{#if loading}
		<CardLoader />
	{/if}
	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="flex justify-between items-center">
		<div class="flex gap-2 text-sm font-semibold">
			<Users size={20} />
			<p>IAM</p>
		</div>
		<div class="flex items-center gap-5">
			<Button
				class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-fit hover:bg-transparent text-xs"
			>
				<UserPlus size={20} class="mr-2" />
				Create User
			</Button>
			<div class="relative z-0">
				<Input
					type="text"
					placeholder="search"
					class="h-9 text-xs rounded w-60 pl-10"
					bind:value={searchTextTemp}
					on:keydown={(e) => {
						if (e.key == 'Enter') {
							searchText = searchTextTemp;
						}
					}}
				/>
				<Search class="absolute left-2 top-[50%] -translate-y-[50%] text-neutral-500" size="20" />
			</div>

			<Select.Root bind:selected={currentRole}>
				<Select.Trigger class="w-[180px] text-xs rounded h-9">
					<Select.Value placeholder="All Roles" />
				</Select.Trigger>
				<Select.Content>
					<Select.Item class="pr-2" value="">All Roles</Select.Item>
					<Select.Item class="pr-2" value="administrator">Administrator</Select.Item>
					<Select.Item class="pr-2" value="student">Student</Select.Item>
				</Select.Content>
			</Select.Root>
		</div>
	</div>

	{#if !fetchError && (!data || data.users.length == 0)}
		<div class="flex grow justify-center items-center">
			<p class="text-sm text-neutral-400 dark:text-neutral-600">No Users Found</p>
		</div>
	{/if}

	{#if data && data.users.length > 0}
		<Table.Root class="text-xs mt-10">
			<Table.Header>
				<Table.Row class="hover:bg-transparent">
					<Table.Head class="font-semibold">Name</Table.Head>
					<Table.Head class="font-semibold">Email</Table.Head>
					<Table.Head class="font-semibold">Roles</Table.Head>
					<Table.Head class="font-semibold">Created Date</Table.Head>
					<Table.Head class="font-semibold">Last Login</Table.Head>
					<Table.Head class="font-semibold text-right">Status</Table.Head>
					<Table.Head class="font-semibold text-right">Actions</Table.Head>
				</Table.Row>
			</Table.Header>
			<Table.Body>
				{#each data.users as user}
					<Table.Row>
						<Table.Cell>{user.firstName} {user.lastName}</Table.Cell>
						<Table.Cell>{user.email}</Table.Cell>
						<Table.Cell class="capitalize">{user.role}</Table.Cell>
						<Table.Cell>
							{user.createdAt ? new Date(user.createdAt).toLocaleString() : 'Not Found'}
						</Table.Cell>
						<Table.Cell>
							{user.lastLogin ? new Date(user.lastLogin).toLocaleString() : 'Not Found'}
						</Table.Cell>
						<Table.Cell>
							{#if user.isActive}
								<span class="flex justify-end items-center gap-2">
									Active <span><CircleCheck size={16} /></span>
								</span>
							{:else}
								<span class="flex justify-end items-center gap-2">
									Deactivated <span><CircleAlert size={16} /></span>
								</span>
							{/if}
						</Table.Cell>
						<Table.Cell class="text-right">
							<DropdownMenu.Root>
								<DropdownMenu.Trigger><Ellipsis size={20} /></DropdownMenu.Trigger>
								<DropdownMenu.Content>
									<DropdownMenu.Item class="text-xs">Profile</DropdownMenu.Item>
									{#if user.isActive}
										<DropdownMenu.Item
											class="text-xs"
											on:click={() => {
												deactivateUser.user = {
													id: user.id,
													firstName: user.firstName,
													lastName: user.lastName
												};
												deactivateUser.open = true;
											}}
										>
											Deactivate
										</DropdownMenu.Item>
									{:else}
										<DropdownMenu.Item
											class="text-xs"
											on:click={() => {
												activateUser.user = {
													id: user.id,
													firstName: user.firstName,
													lastName: user.lastName
												};
												activateUser.open = true;
											}}
										>
											Activate
										</DropdownMenu.Item>
									{/if}
									<DropdownMenu.Item
										class="text-xs"
										on:click={() => {
											deleteUser.user = {
												id: user.id,
												firstName: user.firstName,
												lastName: user.lastName
											};
											deleteUser.open = true;
										}}
									>
										Delete User
									</DropdownMenu.Item>
								</DropdownMenu.Content>
							</DropdownMenu.Root>
						</Table.Cell>
					</Table.Row>
				{/each}
			</Table.Body>
		</Table.Root>
	{/if}

	{#if data && data.users.length > 0}
		<Pagination.Root
			count={data.totalCount}
			{perPage}
			page={pageNum}
			let:pages
			let:currentPage
			class="w-fit mr-0 mt-10"
		>
			<Pagination.Content>
				<Pagination.Item>
					<Pagination.PrevButton
						on:click={() => {
							pageNum = currentPage - 1;
						}}
					/>
				</Pagination.Item>
				{#each pages as page (page.key)}
					{#if page.type === 'ellipsis'}
						<Pagination.Item>
							<Pagination.Ellipsis />
						</Pagination.Item>
					{:else}
						<Pagination.Item isVisible={currentPage == page.value}>
							<Pagination.Link
								{page}
								isActive={currentPage == page.value}
								on:click={() => {
									pageNum = page.value;
								}}
							>
								{page.value}
							</Pagination.Link>
						</Pagination.Item>
					{/if}
				{/each}
				<Pagination.Item>
					<Pagination.NextButton
						on:click={(e) => {
							pageNum = currentPage + 1;
						}}
					/>
				</Pagination.Item>
			</Pagination.Content>
		</Pagination.Root>
	{/if}
</div>

<DeactivateUserDialog
	open={deactivateUser.open}
	user={deactivateUser.user}
	on:deactivate={fetchData}
/>
<DeleteUserDialog open={deleteUser.open} user={deleteUser.user} on:delete={fetchData} />
<ActivateUserDialog open={activateUser.open} user={activateUser.user} on:activate={fetchData} />

<Toaster />
