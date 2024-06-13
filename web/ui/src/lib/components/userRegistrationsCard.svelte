<script>
	import { UserPlus, UserSearch, Users } from 'lucide-svelte';
	import { Chart } from 'chart.js/auto';

	import { isDark } from '../../stores/store';
	import CardLoader from './cardLoader.svelte';
	import CardErrorRetry from './cardErrorRetry.svelte';
	import { onMount } from 'svelte';

	/**
	 * @type {{
	 * 	allUsers: number,
	 *  students: number,
	 * 	instructors: number,
	 *  administrators: number,
	 *  stats: {month: string, count: number }[]
	 * }}
	 */
	let data;

	let loading = true;
	let fetchError = false;

	const fetchData = () => {
		loading = true;
		fetchError = false;

		fetch('/api/user/register-summary')
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
	 * @type {HTMLCanvasElement}
	 */
	let chart;

	/**
	 * @type {Chart}
	 */
	let existingChart;

	$: {
		if (chart && data) {
			let gradient = chart.getContext('2d')?.createLinearGradient(0, 0, 0, 400);

			gradient?.addColorStop(0, $isDark ? 'rgba(0, 253, 152, 0.3)' : 'rgba(16, 185, 129, 0.3)');
			gradient?.addColorStop(0.6, $isDark ? 'rgba(0, 253, 152, 0)' : 'rgba(16, 185, 129, 0)');

			if (existingChart) {
				existingChart.destroy();
			}

			/**
			 * @type {boolean}
			 */
			let delayed;

			existingChart = new Chart(chart, {
				type: 'line',
				options: {
					animation: {
						onComplete: () => {
							delayed = true;
						},
						delay: (context) => {
							let delay = 0;
							if (context.type === 'data' && context.mode === 'default' && !delayed) {
								delay = context.dataIndex * 50 + context.datasetIndex * 100;
							}
							return delay;
						}
					},
					responsive: true,
					backgroundColor: gradient,
					borderColor: $isDark ? '#00FD98' : '#10b981',
					aspectRatio: 0,
					plugins: {
						legend: {
							display: false
						}
					},
					elements: {
						line: {
							borderJoinStyle: 'round',
							tension: 0.3
						}
					},
					scales: {
						x: {
							grid: {
								color: 'rgba(0,0,0,0)'
							}
						},
						y: {
							grid: {
								color: 'rgba(0,0,0,0)'
							}
						}
					}
				},
				data: {
					labels: data.stats.map((row) => row.month),
					datasets: [
						{
							label: 'Registrations',
							data: data.stats.map((row) => row.count),
							fill: true
						}
					]
				}
			});
		}
	}
</script>

<div class="relative py-6 px-4 bg-white dark:bg-neutral-950 rounded overflow-hidden">
	{#if loading}
		<CardLoader />
	{/if}

	{#if fetchError}
		<CardErrorRetry on:retry={fetchData} />
	{/if}

	<div class="space-y-10">
		<div class="flex gap-2 text-sm font-semibold">
			<Users size={20} />
			<p>User Registrations</p>
		</div>

		<div class="w-full h-[280px]">
			<canvas bind:this={chart}></canvas>
		</div>

		<div class="flex flex-wrap gap-10 justify-between">
			<div class="space-y-1">
				<p class="text-xs font-semibold text-neutral-600 dark:text-neutral-300">All Users</p>
				<p class="text-3xl">
					{#if data && data.allUsers}
						{data.allUsers}
					{:else}
						N/A
					{/if}
				</p>
			</div>
			<div class="space-y-1">
				<p class="text-xs font-semibold text-neutral-600 dark:text-neutral-300">Students</p>
				<p class="text-3xl">
					{#if data && data.students}
						{data.students}
					{:else}
						N/A
					{/if}
				</p>
			</div>
			<div class="space-y-1">
				<p class="text-xs font-semibold text-neutral-600 dark:text-neutral-300">Instructors</p>
				<p class="text-3xl">
					{#if data && data.instructors}
						{data.instructors}
					{:else}
						N/A
					{/if}
				</p>
			</div>
			<div class="space-y-1">
				<p class="text-xs font-semibold text-neutral-600 dark:text-neutral-300">Administrators</p>
				<p class="text-3xl">
					{#if data && data.administrators}
						{data.administrators}
					{:else}
						N/A
					{/if}
				</p>
			</div>
		</div>

		<div class="flex gap-8">
			<a href="/">
				<span class="flex text-xs items-center gap-2 font-semibold"
					><UserPlus size={20} /> Create A User</span
				>
			</a>
			<a href="/">
				<span class="flex text-xs items-center gap-2 font-semibold"
					><UserSearch size={20} /> View All Users</span
				>
			</a>
		</div>
	</div>
</div>
