import { redirect } from '@sveltejs/kit';

export async function load() {
	const response = await fetch('/api/auth/is-signed-in');
	if (!response.ok) {
		redirect(302, '/sign-in');
	}
}
