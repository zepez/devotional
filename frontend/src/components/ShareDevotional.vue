<template>
	<a
		class="block w-full bg-blue-500 text-white py-2 text-center rounded-md mb-4"
		:href="devotional.source" target="_blank"
	>
		View Source
	</a>

	<p class="font-semibold uppercase text-xs mb-1">Share this page</p>
	<div class="mb-4">
		<a target="_blank" :href="'https://www.facebook.com/sharer/sharer.php?u=' + devotional.source">F</a>
		<a target="_blank" :href="'https://twitter.com/intent/tweet?text=' + devotional.name + ' ' + devotional.source">T</a>
		<a target="_blank" :href="'https://www.linkedin.com/sharing/share-offsite/?url=' + devotional.source">L</a>
	</div>

	<p class="font-semibold uppercase text-xs mb-1">Share from source</p>
	<div class="mb-4">
		<a target="_blank" :href="'https://www.facebook.com/sharer/sharer.php?u=' + devotional.source">F</a>
		<a target="_blank" :href="'https://twitter.com/intent/tweet?text=' + devotional.name + ' ' + devotional.source">T</a>
		<a target="_blank" :href="'https://www.linkedin.com/sharing/share-offsite/?url=' + devotional.source">L</a>
	</div>
</template>

<script lang="ts">
export default {
	props: ["devotional"]
};
</script>
