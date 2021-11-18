<template>
	<div class="md:sticky top-0 bg-white pb-4">
		<a :href="devotional.source" target="_blank" class="block w-full bg-green-500 text-white py-2 text-center rounded-md mb-4 text-sm">
			View Source
		</a>
	
		<LabelText text="Share this page" />
		<div class="mb-4">
			<ShareArray :href="currentUrl" :name="devotional.name" />
		</div>
	
		<LabelText text="Share from source" />
		<div class="mb-0">
			<ShareArray :href="devotional.source" :name="devotional.name" />
		</div>
	</div>
</template>


<script lang="ts">
import { defineComponent } from "vue";
import LabelText from "../components/LabelText.vue";
import ShareArray from "./ShareArray.vue";


export default defineComponent({
	props: ["devotional"],
	components: {
		LabelText, ShareArray
	},
	setup() {
		return {
			currentUrl: window.location.href
		};
	}
});
</script>
