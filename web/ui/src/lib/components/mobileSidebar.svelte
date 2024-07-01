<script>
	import {
		Activity,
		GraduationCap,
		Info,
		LibrarySquare,
		Plus,
		Settings,
		Ticket,
		Users,
		X
	} from 'lucide-svelte';
	import SidebarLink from './sidebarLink.svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';
	import { userRole } from '../../stores/store';

	let dispatch = createEventDispatcher();
</script>

<div class="z-50 fixed inset-0">
	<div
		tabindex="-1"
		role="button"
		on:keydown={() => dispatch('mobileMenuToggle')}
		on:click={() => dispatch('mobileMenuToggle')}
		class="w-full h-full bg-neutral-100/70 dark:bg-neutral-900/70 backdrop-blur-sm"
	></div>
	<div
		class="fixed left-0 inset-y-0 bg-white dark:bg-neutral-950 rounded py-4 w-60 m-2 overflow-scroll"
	>
		<div class="flex justify-end px-4 mb-10">
			<Button
				on:click={() => dispatch('mobileMenuToggle')}
				class="p-0 bg-transparent h-fit text-neutral-950 dark:text-neutral-100 w-5 hover:bg-transparent"
			>
				<X />
			</Button>
		</div>
		<div class="flex flex-col">
			{#if $userRole == 'administrator'}
				<SidebarLink to="/" on:click={() => dispatch('mobileMenuToggle')}>
					<Activity size={20} /> System Console
				</SidebarLink>
				<SidebarLink to="/create-course" on:click={() => dispatch('mobileMenuToggle')}>
					<Plus size={20} />Create Course
				</SidebarLink>
			{/if}
			<SidebarLink to="/enrolled-courses" on:click={() => dispatch('mobileMenuToggle')}>
				<GraduationCap size={20} />Enrolled Courses
			</SidebarLink>
			<SidebarLink to="/course-library" on:click={() => dispatch('mobileMenuToggle')}>
				<LibrarySquare size={20} />Course Library
			</SidebarLink>
			{#if $userRole == 'administrator'}
				<SidebarLink to="/iam" on:click={() => dispatch('mobileMenuToggle')}>
					<Users size={20} /> IAM
				</SidebarLink>
				<SidebarLink to="/events" on:click={() => dispatch('mobileMenuToggle')}>
					<Info size={20} /> Events
				</SidebarLink>
				<SidebarLink to="/support-tickets" on:click={() => dispatch('mobileMenuToggle')}>
					<Ticket size={20} />Support Tickets
				</SidebarLink>
			{/if}
			{#if $userRole == 'student'}
				<SidebarLink to="/my-support-tickets" on:click={() => dispatch('mobileMenuToggle')}>
					<Ticket size={20} />My Support Tickets
				</SidebarLink>
			{/if}
			<SidebarLink to="/settings" on:click={() => dispatch('mobileMenuToggle')}>
				<Settings size={20} />Settings
			</SidebarLink>
		</div>
	</div>
</div>
