<script>
	import { AlertCircle } from 'lucide-svelte';
</script>

<svelte:head>
	<title>Page Not Found</title>
</svelte:head>

<div class="flex flex-col gap-2 justify-center items-center h-screen">
	<AlertCircle size={30} />
	<p class="text-xl">Page Not Found</p>
	<a href="/" class="text-sm underline">Go Back</a>
</div>
