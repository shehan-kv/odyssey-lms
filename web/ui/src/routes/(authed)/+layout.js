import { redirect } from '@sveltejs/kit';
import { userRole } from '../../stores/store';

export async function load() {
	const response = await fetch('/api/auth/is-signed-in');
	if (!response.ok) {
		redirect(302, '/sign-in');
	} else {
		const userResponse = await fetch('/api/user/self');
		if (userResponse.ok) {
			const user = await userResponse.json();
			userRole.set(user.role);
		}
	}
}
