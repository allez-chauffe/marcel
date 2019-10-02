import svelte from 'rollup-plugin-svelte';
import resolve from 'rollup-plugin-node-resolve';
import pkg from './package.json';

const name = pkg.name
	.replace(/^(@\S+\/)?(svelte-)?(\S+)/, '$3')
	.replace(/^\w/, m => m.toUpperCase())
	.replace(/-\w/g, m => m[1].toUpperCase());

console.log(name)

export default {
	input: pkg.svelte,
	output: [
		{ file: pkg.module, format: 'es' },
		{ file: pkg.main, format: 'umd', name }
	],
	plugins: [
		svelte({
			customElement: true
		}),
		resolve()
	]
};
