<script>
	import UserRegistrationsCard from '$lib/components/userRegistrationsCard.svelte';
	import ServerInformationCard from '$lib/components/serverInformationCard.svelte';
	import RecentEventsCard from '$lib/components/recentEventsCard.svelte';
	import RecentCoursesCard from '$lib/components/recentCoursesCard.svelte';
	import SupportTicketsCard from '$lib/components/supportTicketsCard.svelte';
</script>

<svelte:head>
	<title>System Console</title>
</svelte:head>

<div class="grow space-y-2 overflow-scroll">
	<div class="grid grid-cols-1 lg:grid-cols-2 gap-2 w-full">
		<UserRegistrationsCard />
		<div class="flex flex-col gap-2">
			<ServerInformationCard />
			<RecentEventsCard />
		</div>
	</div>

	<RecentCoursesCard />
	<SupportTicketsCard />
</div>
