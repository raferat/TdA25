import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
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
