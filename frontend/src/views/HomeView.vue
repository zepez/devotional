<template>
	<div class="container">
		<ul class="-my-3 divide-y divide-gray-200">
			<li v-for="devotional in devotionals" :key="devotional.name" class="py-3">
				<h4 class="font-semibold text-xl">{{ devotional.name }}</h4>
				<p>{{ devotional.plain_text }}</p>
			</li>
		</ul>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";

interface Devotional {
	name: string;
	plain_text: string;
}

const devotionals = [
	{
		name: "testing1",
		plain_text: "testing plain text"
	},
	{
		name: "testing2",
		plain_text: "testing plain text"
	}
];

export default defineComponent({
	data() {
		return {
			devotionals: devotionals as Devotional[]
		};
	}
});
</script>
