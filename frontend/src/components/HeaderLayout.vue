<template>
	<div class="bg-white pt-8 pb-6">
		<h1 class="text-5xl text-center font-semibold mb-6">zepez / devotional demo</h1>
		<nav>
			<ul class="flex justify-center flex-wrap">
				<li v-for="link in navLinks" :key="link.name" class="mx-8 px-4 my-2 text-2xl border-2 border-gray-800">
					<a :href="link.href" :target="link.external ? '_blank' : ''">
						{{ link.name }}
					</a>
				</li>
			</ul>
		</nav>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";

const navLinks = [
	{ name: "GitHub", href: "https://github.com/zepez/devotional", external: true },
	{ name: "lhm.org", href: "https://lhm.org", external: true },
	{ name: "Portfolio", href: "https://zep.sh", external: true }
];

export default defineComponent({
	data() {
		return {
			navLinks
		};
	}
});
</script>
