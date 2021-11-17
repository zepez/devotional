<template>
	<div class="container">
		<ul class="-my-3 divide-y divide-gray-200">
			<li v-for="devotional in devotionals" :key="devotional.name" class="py-3 grid grid-cols-6 gap-4">
				<img :src="devotional.image" :alt="devotional.name" class="col-span-2 rounded-md"/>
				<div class="col-span-4">
					<h4 class="font-semibold text-2xl mb-2">{{ devotional.name }}</h4>
					<p>{{ devotional.plain_text.substring(0, 150) + "..." }}</p>
				</div>
			</li>
		</ul>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import api from "../config/axios";


interface Devotional {
	_id: string;
	source: string;
	target_date: string;
	name: string;
	image: string;
	html: string;
	plain_text: string;
	created_at: string;
	updated_at: string;
}

const devotionals = await api.get("/devotionals/0")
	.then(res => res.data)
	.catch(e => console.log(e));


export default defineComponent({
	data() {
		return {
			devotionals: devotionals as Devotional[]
		};
	}
});
</script>
