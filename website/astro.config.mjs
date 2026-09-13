// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

const guide = (lang, ja) => ({
	label: lang,
	translations: { ja },
	items: [
		{ label: 'Basics', translations: { ja: '基本' }, slug: `guides/${lang.toLowerCase().replace('.', '')}/basics` },
		{ label: 'Testing', translations: { ja: 'テストへの組み込み' }, slug: `guides/${lang.toLowerCase().replace('.', '')}/testing` },
	],
});

export default defineConfig({
	site: 'https://shibukawa.github.io',
	base: '/pgmem',
	integrations: [
		starlight({
			title: 'pgmem',
			description: 'Real PostgreSQL 18, in memory, for tests. No Docker.',
			defaultLocale: 'root',
			locales: {
				root: { label: 'English', lang: 'en' },
				ja: { label: '日本語', lang: 'ja' },
			},
			social: [{ icon: 'github', label: 'GitHub', href: 'https://github.com/shibukawa/pgmem' }],
			sidebar: [
				{
					label: 'Overview',
					translations: { ja: '概要' },
					items: [
						{ label: 'Quickstart', translations: { ja: 'クイックスタート' }, slug: 'quickstart' },
						{ label: 'Why pgmem', translations: { ja: 'なぜ pgmem か' }, slug: 'why' },
						{ label: 'Benchmarks', translations: { ja: 'ベンチマーク' }, slug: 'benchmarks' },
						{ label: 'Architecture', translations: { ja: 'アーキテクチャ' }, slug: 'architecture' },
						{ label: 'Extensions', translations: { ja: '拡張機能' }, slug: 'extensions' },
						{ label: 'Limits', translations: { ja: '制限事項' }, slug: 'limits' },
						{ label: 'Versioning', translations: { ja: 'バージョン番号' }, slug: 'versioning' },
					],
				},
				guide('Go', 'Go'),
				guide('Python', 'Python'),
				guide('Java', 'Java'),
				guide('Node.js', 'Node.js'),
			],
		}),
	],
});
