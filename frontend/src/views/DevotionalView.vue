<template>
	<div class="container grid grid-cols-3 gap-4 divide-y md:divide-y-0 md:divide-x divide-gray-200">
		<div class="col-span-3 md:col-span-2">
			<h1 class="mb-4 text-5xl font-serif">{{ devotional.name }}</h1>

			<div class="border-t border-gray-200 mb-4 pt-2 text-sm">
				<p>{{ formatTargetDate(devotional.target_date) }}</p>
			</div>

			<img class="mb-4 rounded-sm" :src="devotional.image" :alt="devotional.name" />
			<div v-html="devotional.html"></div>
		</div>
		<div class="col-span-3 md:col-span-1 py-4 md:pl-4 md:py-0">
			<ShareDevotional v-bind:devotional="devotional" />

			<LabelText text="More devotionals:" />
			<ul class="divide-y divide-gray-200 -my-3">
				<li class="py-3" v-for="oneMore in more" :key="oneMore.name">
					<router-link :to="'/d/' + oneMore.id">
						<h4 class="font-semibold text-xl">{{ oneMore.name }}</h4>
						<p>{{ oneMore.plain_text.substring(0, 85) + "..." }}</p>
					</router-link>
				</li>
			</ul>
		</div>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import api from "../config/axios";
import formatTargetDate from "../utils/formatTargetDate";
import ShareDevotional from "../components/ShareDevotional.vue";
import LabelText from "../components/LabelText.vue";


const getById = async (id:string) => {
	const response = await api.get(`/devotional/${id}`)
		.then(res => res.data)
		.catch(e => console.error(e));
	
	return response;
};

const getMore = async () => {
	const response = await api.get("/devotionals/0")
		.then(res => res.data)
		.catch(e => console.error(e));
	
	return response;
};


export default defineComponent({
	components: { ShareDevotional, LabelText },
	data() {
		return { 
			devotional: {} as Devotional, 
			more: [] as Devotional[],
			formatTargetDate
		};
	},

	async created() {
		this.devotional = await getById(this.$route.params.id.toString());
		this.more = await getMore();
	},

	async beforeRouteUpdate(to) {
		this.devotional = await getById(to.params.id.toString());
	},
});
</script>
