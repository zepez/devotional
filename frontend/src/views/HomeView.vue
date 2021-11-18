<template>
	<div class="container">
		<ul class="-my-3 divide-y divide-gray-200">
			<li v-for="devotional in devotionals" :key="devotional.name">
				<router-link :to="'/d/' + devotional.id" class="py-3 grid grid-cols-6 gap-4 home-list-link">
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
import { defineComponent, ref, onMounted } from "vue";
import api from "../config/axios";
import formatTargetDate from "../utils/formatTargetDate";
import { useMeta } from "vue-meta";


export default defineComponent({
	setup() {
		const devotionals = ref([] as Devotional[]);

		useMeta({
			title: "home",
			htmlAttrs: { lang: "en" },
			description: "",
			og: {
				site_name: window.location.host,
				title: "home | zepez/devotional",
				description: "This is a demo project, showing off gihub.com/zepez/devotional. Made by Alex Zepezauer.",
				image: window.location.host + "/social/share-card.jpg",
				url: window.location.href
			},
			twitter: {
				title: "home | zepez/devotional",
				description: "This is a demo project, showing off gihub.com/zepez/devotional. Made by Alex Zepezauer.",
				image: window.location.host + "/social/share-card.jpg",
				card: "summary_large_image"
			}
		});

		const getDevotionals = async () => {
			return await api.get("/devotionals/0")
				.then(res => res.data)
				.catch(e => console.log(e));
		};

		onMounted(async () => {
			devotionals.value = await getDevotionals();
			console.log(window.location);
		});

		return {
			devotionals,
			formatTargetDate,
		};
	}
});
</script>
