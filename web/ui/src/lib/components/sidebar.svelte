<script>
	import {
		Activity,
		GraduationCap,
		Info,
		LibrarySquare,
		Plus,
		Settings,
		Ticket,
		Users
	} from 'lucide-svelte';
	import SidebarLink from './sidebarLink.svelte';
	import { userRole } from '../../stores/store';
</script>

<div class="hidden xl:block bg-white dark:bg-neutral-950 rounded py-4 w-60">
	<div class="flex flex-col">
		{#if $userRole == 'administrator'}
			<SidebarLink to="/"><Activity size={20} /> System Console</SidebarLink>
			<SidebarLink to="/create-course"><Plus size={20} />Create Course</SidebarLink>
		{/if}
		<SidebarLink to="/enrolled-courses"><GraduationCap size={20} />Enrolled Courses</SidebarLink>
		<SidebarLink to="/course-library"><LibrarySquare size={20} />Course Library</SidebarLink>
		{#if $userRole == 'administrator'}
			<SidebarLink to="/iam"><Users size={20} /> IAM</SidebarLink>
			<SidebarLink to="/events"><Info size={20} /> Events</SidebarLink>
			<SidebarLink to="/support-tickets"><Ticket size={20} />Support Tickets</SidebarLink>
		{/if}
		{#if $userRole == 'student'}
			<SidebarLink to="/my-support-tickets"><Ticket size={20} />My Support Tickets</SidebarLink>
		{/if}
		<SidebarLink to="/settings"><Settings size={20} />Settings</SidebarLink>
	</div>
</div>
