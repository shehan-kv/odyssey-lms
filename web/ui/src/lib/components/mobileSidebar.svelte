<script>
	import {
		Activity,
		Bell,
		GraduationCap,
		Info,
		LibrarySquare,
		Settings,
		Ticket,
		Users,
		X
	} from 'lucide-svelte';
	import SidebarLink from './sidebarLink.svelte';
	import LecternIcon from './lecternIcon.svelte';
	import Button from './ui/button/button.svelte';
	import { createEventDispatcher } from 'svelte';

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
			<SidebarLink to="/"><Activity size={20} /> System Console</SidebarLink>
			<SidebarLink to="/instructor-console"
				><LecternIcon size={20} /> Instructor Console</SidebarLink
			>
			<SidebarLink to="/enrolled-courses"><GraduationCap size={20} />Enrolled Courses</SidebarLink>
		</div>

		<div class="flex flex-col mt-5">
			<SidebarLink to="/course-library"><LibrarySquare size={20} />Course Library</SidebarLink>
			<SidebarLink to="/iam"><Users size={20} /> IAM</SidebarLink>
			<SidebarLink to="/events"><Info size={20} /> Events</SidebarLink>
			<SidebarLink><Bell size={20} />Notifications</SidebarLink>
			<SidebarLink to="/support-tickets"><Ticket size={20} />Support Tickets</SidebarLink>
			<SidebarLink><Settings size={20} />Settings</SidebarLink>
		</div>
	</div>
</div>
