import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit(),
	],
	//remove after dev
	server: {
		proxy: {
			'/api': {
				target: 'http://localhost:4242',
				changeOrigin: true,
			}
		}
	}
});
