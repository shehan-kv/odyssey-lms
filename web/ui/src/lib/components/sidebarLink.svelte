<script>
	export let to = '/';
</script>

<a href={to} on:click>
	<span
		class="flex gap-4 py-2 px-4 text-neutral-700 hover:text-neutral-950
		dark:text-neutral-300 dark:hover:text-neutral-100 text-sm"
	>
		<slot></slot>
	</span>
</a>
