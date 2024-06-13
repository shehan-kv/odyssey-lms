<script>
	import { cn } from '$lib/utils.js';
	let className = undefined;
	export { className as class };
</script>

<td
	class={cn('px-4 py-3 align-middle [&:has([role=checkbox])]:pr-0', className)}
	{...$$restProps}
	on:click
	on:keydown
>
	<slot />
</td>
