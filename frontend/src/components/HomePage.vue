<template>
	<div v-for="devotional in devotionals" :key="devotional.name">
		<h1>{{ devotional.name }}</h1>
		<p>{{ devotional.plain_text }}</p>
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
