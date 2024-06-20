<script>
	import {
		Activity,
		Bell,
		GraduationCap,
		Info,
		LibrarySquare,
		SendHorizonal,
		Settings,
		Ticket,
		Users
	} from 'lucide-svelte';
	import LecternIcon from './lecternIcon.svelte';
	import SidebarLink from './sidebarLink.svelte';
</script>

<div class="hidden xl:block bg-white dark:bg-neutral-950 rounded py-4 w-60">
	<div class="flex flex-col">
		<SidebarLink to="/"><Activity size={20} /> System Console</SidebarLink>
		<SidebarLink to="/instructor-console"><LecternIcon size={20} /> Instructor Console</SidebarLink>
		<SidebarLink to="/enrolled-courses"><GraduationCap size={20} />Enrolled Courses</SidebarLink>
	</div>

	<div class="flex flex-col mt-5">
		<SidebarLink to="/course-library"><LibrarySquare size={20} />Course Library</SidebarLink>
		<SidebarLink to="/iam"><Users size={20} /> IAM</SidebarLink>
		<SidebarLink to="/events"><Info size={20} /> Events</SidebarLink>
		<SidebarLink to="/notifications"><Bell size={20} />Notifications</SidebarLink>
		<SidebarLink to="/sent-notifications"><SendHorizonal size={20} />Sent Notifications</SidebarLink
		>
		<SidebarLink to="/support-tickets"><Ticket size={20} />Support Tickets</SidebarLink>
		<SidebarLink to="/settings"><Settings size={20} />Settings</SidebarLink>
	</div>
</div>
