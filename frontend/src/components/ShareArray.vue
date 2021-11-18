<template>
	<div class="grid grid-cols-3 gap-4 text-white text-center text-xs">
		<a target="_blank" :href="'https://www.facebook.com/sharer/sharer.php?u=' + href">
			<div class="bg-blue-700 rounded-md py-2">
				Facebook
			</div>
		</a>
		<a target="_blank" :href="'https://twitter.com/intent/tweet?text=' + name + ' ' + href">
			<div class="bg-blue-400 rounded-md py-2">
				Twitter
			</div>
		</a>
		<a target="_blank" :href="'https://www.linkedin.com/sharing/share-offsite/?url=' + href">
			<div class="bg-blue-500 rounded-md py-2">
				LinkedIn
			</div>
		</a>
	</div>
</template>


<script lang="ts">
export default {
	props: ["href", "name"]
};
</script>
