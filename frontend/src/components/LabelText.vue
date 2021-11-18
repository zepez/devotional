<template>
	<p class="font-semibold uppercase text-xs mb-1">{{ text }}</p>
</template>


<script lang="ts">
import { defineComponent } from "vue";


export default defineComponent({
	props: ["text"]
});
</script>
