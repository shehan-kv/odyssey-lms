<script>
	import { Cpu, HardDrive, MemoryStick, Server } from 'lucide-svelte';
	import CardErrorRetry from './cardErrorRetry.svelte';
	import CardLoader from './cardLoader.svelte';
	import { onMount } from 'svelte';

	let loading = true;
	let fetchError = false;

	/**
	 * @type {{
	 * 	cpu: number,
	 *  memory: number,
	 *  uploadsFree: number,
	 *  uploadsTotal: number,
	 * }}
	 */
	let data;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/system')
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

	const KB = 1024;
	const MB = KB * 1024;
	const GB = MB * 1024;
	const TB = GB * 1024;

	/**
	 * @param {number} bytes
	 */
	function formatBytes(bytes) {
		if (bytes === 0) return 0;

		if (bytes < KB) return bytes;
		if (bytes >= KB && bytes < MB) return (bytes / KB).toFixed(0);
		if (bytes >= MB && bytes < GB) return (bytes / MB).toFixed(0);
		if (bytes >= GB && bytes < TB) return (bytes / GB).toFixed(0);
		return (bytes / TB).toFixed(0);
	}

	/**
	 * @param {number} bytes
	 */
	function formatSuffix(bytes) {
		if (bytes === 0) return 'B';

		if (bytes < KB) return 'B';
		if (bytes >= KB && bytes < MB) return 'KB';
		if (bytes >= MB && bytes < GB) return 'MB';
		if (bytes >= GB && bytes < TB) return 'GB';
		return 'TB';
	}

	$: usedDisk = data
		? (((data.uploadsTotal - data.uploadsFree) / data.uploadsTotal) * 100).toFixed(0)
		: 0;

	onMount(() => {
		fetchData();
	});
</script>

<div class="relative py-6 px-4 rounded bg-white dark:bg-neutral-950 overflow-hidden h-fit">
	{#if loading}
		<CardLoader />
	{/if}

	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="space-y-10">
		<div class="flex gap-2 text-sm font-semibold">
			<Server size={20} />
			<p>Server Information</p>
		</div>

		<div class="flex flex-wrap gap-10 justify-between">
			<div class="space-y-2">
				<div class="flex gap-2 text-sm font-semibold text-neutral-600 dark:text-neutral-300">
					<Cpu size={20} />
					<p>CPU</p>
				</div>
				<p class="text-3xl">
					{#if data && data.cpu}
						{data.cpu}
					{:else}
						N/A
					{/if}
					<span class="text-xs">CORES</span>
				</p>
			</div>
			<div class="space-y-2">
				<div class="flex gap-2 text-sm font-semibold text-neutral-600 dark:text-neutral-300">
					<MemoryStick size={20} />
					<p>System Memory</p>
				</div>
				<p class="text-3xl">
					{#if data && data.memory}
						{formatBytes(data.memory)}
						<span class="text-xs">{formatSuffix(data.memory)}</span>
					{:else}
						N/A
					{/if}
				</p>
			</div>
			<div class="space-y-2">
				<div class="flex gap-2 text-sm font-semibold text-neutral-600 dark:text-neutral-300">
					<HardDrive size={20} />
					<p>Uploads Storage</p>
				</div>
				<p class="text-3xl">
					{#if data && data.uploadsFree}
						{formatBytes(data.uploadsFree)}
						<span class="text-xs">{formatSuffix(data.uploadsFree)} FREE</span>
					{:else}
						N/A
					{/if}
					/

					{#if data && data.uploadsTotal}
						{formatBytes(data.uploadsTotal)}
						<span class="text-xs">{formatSuffix(data.uploadsTotal)}</span>
					{:else}
						N/A
					{/if}
				</p>
				<div class="h-1.5 bg-emerald-100 dark:bg-emerald-950 rounded oveflow-hidden">
					<div
						class={`h-1.5 bg-emerald-500 rounded max-w-full`}
						style={`width: ${usedDisk}%`}
					></div>
				</div>
			</div>
		</div>
	</div>
</div>
