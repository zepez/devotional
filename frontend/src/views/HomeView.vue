<template>
	<div class="container">
		<ul class="-my-3 divide-y divide-gray-200">
			<li v-for="devotional in devotionals" :key="devotional.name">
				<router-link :to="'/d/' + devotional.id" class="py-3 grid grid-cols-6 gap-4">
					<img :src="devotional.image" :alt="devotional.name" class="col-span-6 md:col-span-2 rounded-md"/>
					<div class="col-span-6 md:col-span-4">
						<h4 class="font-semibold text-2xl">{{ devotional.name }}</h4>
						<p class="py-1 text-sm">{{ formatTargetDate(devotional.target_date) }}</p>
						<p>{{ devotional.plain_text.substring(0, 200) + "..." }}</p>
					</div>
				</router-link>
			</li>
		</ul>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import api from "../config/axios";
import formatTargetDate from "../utils/formatTargetDate";


const devotionals = await api.get("/devotionals/0")
	.then(res => res.data)
	.catch(e => console.log(e));


export default defineComponent({
	data() {
		return {
			devotionals: devotionals as Devotional[],
			formatTargetDate
		};
	}
});
</script>
