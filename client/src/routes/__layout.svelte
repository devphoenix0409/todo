<script context="module">
	import { apolloQueryClient } from '$lib/apollo-client';
	/**
	 * @type {import('@sveltejs/kit').Load}
	 */
	export async function load({ stuff }) {
		const client = apolloQueryClient();
		return {
			stuff: {
				...stuff,
				client
			}
		};
	}
</script>

<slot />
