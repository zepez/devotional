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
				<li class="py-3 d-list-link" v-for="oneMore in more" :key="oneMore.name">
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
import { defineComponent, ref, onMounted } from "vue";
import { useMeta } from "vue-meta";
import { useRoute } from "vue-router";
import api from "../config/axios";
import formatTargetDate from "../utils/formatTargetDate";
import ShareDevotional from "../components/ShareDevotional.vue";
import LabelText from "../components/LabelText.vue";


export default defineComponent({
	components: { ShareDevotional, LabelText },

	setup() {
		const devotional = ref({} as Devotional);
		const more = ref([] as Devotional[]);
		const route = useRoute();

		const { meta } = useMeta({
			title: "view", 
			htmlAttrs: { lang: "en" }
		});

		const getById = async () => {
			return await api.get(`/devotional/${route.params.id}`)
				.then(res => res.data)
				.catch(e => console.error(e));
		};

		const getMore = async () => {
			return await api.get("/devotionals/0")
				.then(res => res.data)
				.catch(e => console.error(e));
		};

		onMounted(async () => {
			more.value = await getMore();
			devotional.value = await getById();

			meta.title = devotional.value.name;
			meta.description = devotional.value.plain_text.substring(0, 100) + "...";
			meta.og = {
				site_name: window.location.host,
				title: devotional.value.name + " | zepez/devotional",
				description: devotional.value.plain_text.substring(0, 100) + "...",
				image: devotional.value.image,
				url: window.location.href
			};
			meta.twitter = {
				title: devotional.value.name + " | zepez/devotional",
				description: devotional.value.plain_text.substring(0, 100) + "...",
				image: devotional.value.image,
				card: "summary_large_image"
			};
		});

		return { devotional, more, formatTargetDate };
	},
});
</script>
